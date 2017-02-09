# Bitfan

Bitfan is an open source data processing pipeline.

![Bitfan logo](docs/static/noun_307496_cc.png "Bitfan")

## Install

### Download binary
linux, windows, osx available here : https://github.com/vjeantet/bitfan/releases

### Or compile from sources
```
$ go get -u github.com/vjeantet/bitfan
```

## Run 
Example with a remote configuration file which ingest data from stdin and output a tranformation to stdout.
```
$ bitfan run https://raw.githubusercontent.com/vjeantet/bitfan/master/examples.d/simple.conf
```
copy/paste this in your console

```
127.0.0.1 - - [11/Dec/2013:00:01:45 -0800] "GET /xampp/status.php HTTP/1.1" 200 3891 "http://cadenza/xampp/navi.php" "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.9; rv:25.0) Gecko/20100101 Firefox/25.0"
```

## Other commands
type `bitfan help` to display usage information

	
```
Usage:
  bitfan [flags]
  bitfan [command]

Available Commands:
  doc         Display documentation about plugins
  list        List running pipelines
  run         Run bitfan
  service     Install and manage bitfan service
  start       Start a new pipeline in a running bitfan
  stop        Stop a running pipeline
  test        Test configurations (files, url, directories)
  version     Display version informations

Flags:
  -f, --config string       Load the Logstash config from a file a directory or a url
  -t, --configtest          Test config file or directory
      --debug               Increase verbosity to the last level (trace), more verbose.
  -e, --eval string         Use the given string as the configuration data.
  -w, --filterworkers int   number of workers (default 4)
  -h, --help                help for bitfan
  -l, --log string          Log to a given path. Default is to log to stdout.
      --settings string     Set the directory containing the bitfan.toml settings (default "current dir, then ~/.bitfan/ then /etc/bitfan/")
      --verbose             Increase verbosity to the first level (info), less verbose.
  -V, --version             Display version info.

Use "bitfan [command] --help" for more information about a command.
```

  logstash flags works as well `-f`, `-e`, `--configtest`, ...


## TODO

- [x] parse logstash config file
- [x] support command line flags "à la logstash"
- [x] generic input support
- [x] generic filter support
- [x] generic output support
- [x] configuration condition (if else) support
- [x] dynamic %{field.key} support in config file
- [x] gracefully stop
- [x] gracefully start
- [x] name all contributors, imported packages, similar projects
- [x] use remote configuration file
- [x] include local and remote files from configuration files
- [ ] codec support
- [x] log to file
- [x] plugins autodocumentation
- [x] install bitfan as a system daemon / service
- [x] list currently runnnung pipelines
- [x] start new pipelines in a running instance
- [x] stop a pipeline without stopping other
- [x] import external configuration from configuration (use)
- [x] dispatch message to another configuration from configuration (fork)


# Supported inputs, filters and outputs in config file

type `bitfan doc` to list all available plugins

## INPUT

|  PLUGIN  |          DESCRIPTION           |
|----------|--------------------------------|
| twitter  |                                |
| exec     |                                |
| unix     |                                |
| stdin    | Reads events from standard  input  |
| file     |                                |
| beats    |                                |
| rabbitmq |                                |
| udp      |                                |
| syslog   |                                |
| readfile |                                |

type `bitfan doc pluginname` to get more information about plugin configuration and usage

## FILTER

| PLUGIN |          DESCRIPTION           |
|--------|--------------------------------|
| date   | Parses dates from fields to use as the Bitfan timestamp  for an event |
| grok   |                                |
| split  | Splits multi-line messages into distinct events |
| json   | Parses JSON events             |
| uuid   | Adds a UUID to events          |
| drop   | Drops all events               |
| geoip  | Adds geographical information about an IP address |
| kv     | Parses key-value pairs         |
| html   |                                |
| mutate |                                |

type `bitfan doc pluginname` to get more information about plugin configuration and usage

## OUTPUT

|     PLUGIN     |          DESCRIPTION           |
|----------------|--------------------------------|
| stdout         | Prints events to the standard output |
| null           | Drops everything received      |
| file           |                                |
| glusterfs      |                                |
| statsd         |                                |
| mongodb        |                                |
| elasticsearch  |                                |
| elasticsearch2 |                                |
| rabbitmq       |                                |
| email          | Sends email when output is received |

type `bitfan doc pluginname` to get more information about plugin configuration and usage

## SPECIAL for all sections
|     PLUGIN     |          DESCRIPTION           |
|----------------|--------------------------------|
| use         | reference another configuration file (URL or local path) to include (copy/paste) in your current configuration  |


# Used package
* kardianos/govendor Go vendor tool that works with the standard vendor file
* spf13/cobra - A Commander for modern Go CLI interactions
* bbuck/go-lexer (a forked version) - Lexer based on Rob Pike's talk on YouTube
* vjeantet/bitfan/lib - all plugins and runtime used by bitfan 


# Similar projets in go

* tsaikd/gogstash - Logstash like, written in golang
* packetzoom/logzoom - A lightweight replacement for logstash indexer in Go
* hailocab/logslam - A lightweight lumberjack protocol compliant logstash indexer


# Credits
logo "hand fan" by lastspark from the Noun Project

# Contributors
* Valere JEANTET
* Merlin Gaillard
* Alexander AKULOV
