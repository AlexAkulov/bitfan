# STATSD


## Synopsys


|     SETTING     |  TYPE   | REQUIRED |   DEFAULT VALUE    |
|-----------------|---------|----------|--------------------|
| host            | string  | false    | "localhost"        |
| port            | int     | false    |               8125 |
| protocol        | string  | false    | "udp"              |
| sender          | string  | false    | ""                 |
| count           | hash    | false    | {}                 |
| decrement       | array   | false    | []                 |
| gauge           | hash    | false    | {}                 |
| increment       | array   | false    | []                 |
| namespace       | string  | false    | "bitfan"           |
| sample_rate     | float32 | false    |                1.0 |
| set             | hash    | false    | {}                 |
| timing          | hash    | false    | {}                 |
| validate_regexp | string  | false    | "[^a-zA-Z0-9_:#-]" |


## Details

### host
* Value type is string
* Default value is `"localhost"`

The address of the statsd server.

### port
* Value type is int
* Default value is `8125`

The port to connect to on your statsd server.

### protocol
* Value type is string
* Default value is `"udp"`



### sender
* Value type is string
* Default value is `""`

The name of the sender. Dots will be replaced with underscores.

### count
* Value type is hash
* Default value is `{}`

A count metric. metric_name => count as hash

### decrement
* Value type is array
* Default value is `[]`

A decrement metric. Metric names as array.

### gauge
* Value type is hash
* Default value is `{}`

A gauge metric. metric_name => gauge as hash.

### increment
* Value type is array
* Default value is `[]`

An increment metric. Metric names as array.

### namespace
* Value type is string
* Default value is `"bitfan"`

The statsd namespace to use for this metric.

### sample_rate
* Value type is float32
* Default value is `1.0`

The sample rate for the metric.

### set
* Value type is hash
* Default value is `{}`

A set metric. metric_name => "string" to append as hash

### timing
* Value type is hash
* Default value is `{}`

A timing metric. metric_name => duration as hash

### validate_regexp
* Value type is string
* Default value is `"[^a-zA-Z0-9_:#-]"`

Defines the characters that allowed in metric names. Any character is not in this list, is replaced by with "_" (underscore)



## Configuration blueprint

```
statsd{
	host => "localhost"
	port => 8125
	protocol => "udp"
	sender => ""
	count => {}
	decrement => []
	gauge => {}
	increment => []
	namespace => "bitfan"
	sample_rate => 1.0
	set => {}
	timing => {}
	validate_regexp => "[^a-zA-Z0-9_:#-]"
}
```
