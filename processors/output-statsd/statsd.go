//go:generate bitfanDoc
package statsd

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/ShowMax/go-fqdn"
	"github.com/vjeantet/bitfan/processors"
	"gopkg.in/alexcesaro/statsd.v2"
)

func New() processors.Processor {
	return &processor{opt: &options{}}
}

type processor struct {
	processors.Base
	metricRe *regexp.Regexp
	opt      *options
	conn     *statsd.Client
}

type options struct {
	// The address of the statsd server.
	// @Default "localhost"
	Host string `mapstructure:"host"`

	// The port to connect to on your statsd server.
	// @Default 8125
	Port int `mapstructure:"port"`

	// @Default "udp"
	Protocol string `mapstructure:"protocol"`

	// The name of the sender. Dots will be replaced with underscores.
	Sender string `mapstructure:"sender"`

	// A count metric. metric_name => count as hash
	Count map[string]interface{} `mapstructure:"count"`

	// A decrement metric. Metric names as array.
	Decrement []string `mapstructure:"decrement"`

	// A gauge metric. metric_name => gauge as hash.
	Gauge map[string]interface{} `mapstructure:"gauge"`

	// An increment metric. Metric names as array.
	Increment []string `mapstructure:"increment"`

	// The statsd namespace to use for this metric.
	// @Default "bitfan"
	NameSpace string `mapstructure:"namespace"`

	// The sample rate for the metric.
	// @Default 1.0
	SampleRate float32 `mapstructure:"sample_rate"`

	// A set metric. metric_name => "string" to append as hash
	Set map[string]interface{} `mapstructure:"set"`

	// A timing metric. metric_name => duration as hash
	Timing map[string]interface{} `mapstructure:"timing"`

	// Defines the characters that allowed in metric names. Any character is not in this list, is replaced by with "_" (underscore)
	// @Default "[^a-zA-Z0-9_:#-]"
	ValidateRegexp string `mapstructure:"validate_regexp"`
}

func (p *processor) Configure(ctx processors.ProcessorContext, conf map[string]interface{}) error {
	defaults := options{
		Host:       "localhost",
		Port:       8125,
		Protocol:   "udp",
		Sender:     fqdn.Get(),
		NameSpace:  "bitfan",
		SampleRate: 1.0,
		ValidateRegexp: "[^a-zA-Z0-9_:#-]",
	}
	p.opt = &defaults
	if err := p.ConfigureAndValidate(ctx, conf, p.opt); err != nil {
		return err
	}
	var err error
	p.metricRe, err = regexp.Compile(p.opt.ValidateRegexp)
	return err
}

func (p *processor) Receive(e processors.IPacket) error {
	for key, value := range p.opt.Count {
		v, err := p.dynamicValue(value, e)
		if err != nil {
			p.Logger.Warnf("string [%v] can't used as counter value: %v", value, err)
			continue
		}
		p.conn.Count(p.dynamicKey(key, e), v)
	}
	for _, key := range p.opt.Increment {
		p.conn.Count(p.dynamicKey(key, e), 1)
	}
	for _, key := range p.opt.Decrement {
		p.conn.Count(p.dynamicKey(key, e), -1)
	}
	for key, value := range p.opt.Gauge {
		v, err := p.dynamicValue(value, e)
		if err != nil {
			p.Logger.Warnf("string [%v] can't used as gauge value: %v", value, err)
			continue
		}
		p.conn.Gauge(p.dynamicKey(key, e), v)
	}
	for key, value := range p.opt.Timing {
		v, err := p.dynamicValue(value, e)
		if err != nil {
			p.Logger.Warnf("string [%v] can't used as timing value: %v", value, err)
			continue
		}
		p.conn.Timing(p.dynamicKey(key, e), v)
	}
	for key, value := range p.opt.Set {
		v, err := p.dynamicValue(value, e)
		if err != nil {
			p.Logger.Warnf("string [%v] can't used as set value: %v", value, err)
			continue
		}
		p.conn.Unique(p.dynamicKey(key, e), fmt.Sprintf("%v", v))
	}
	return nil
}

func (p *processor) dynamicValue(value interface{}, e processors.IPacket) (float64, error) {
	v := fmt.Sprintf("%v", value)
	processors.Dynamic(&v, e.Fields())
	return strconv.ParseFloat(v, 10)
}

func (p *processor) dynamicKey(key string, e processors.IPacket) string {
	k, s := key, p.opt.Sender
	processors.Dynamic(&k, e.Fields())
	processors.Dynamic(&s, e.Fields())
	return fmt.Sprintf("%s.%s", p.metricRe.ReplaceAllString(s, "_"), k)
}

func (p *processor) Start(e processors.IPacket) error {
	url := fmt.Sprintf("%s:%d", p.opt.Host, p.opt.Port)
	var err error
	p.conn, err = statsd.New(
		statsd.Address(url),
		statsd.Prefix(p.opt.NameSpace),
		statsd.Network(p.opt.Protocol),
		statsd.SampleRate(p.opt.SampleRate),
		statsd.ErrorHandler(func(err error) {
			p.Logger.Error(err)
		}),
	)
	return err
}

func (p *processor) Stop(e processors.IPacket) error {
	p.conn.Close()
	return nil
}
