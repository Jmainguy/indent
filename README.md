# indent
[![Go Report Card](https://goreportcard.com/badge/github.com/Jmainguy/indent)](https://goreportcard.com/badge/github.com/Jmainguy/indent)
[![Release](https://img.shields.io/github/release/Jmainguy/indent.svg?style=flat-square)](https://github.com/Jmainguy/indent/releases/latest)
[![Coverage Status](https://coveralls.io/repos/github/Jmainguy/indent/badge.svg?branch=main)](https://coveralls.io/github/Jmainguy/indent?branch=main)

# indent
A tool to indent a file to stdout using spaces.

## Usage
```/bin/bash
Usage of ./indent:
  -filename string
    	Filename to read
  -indent int
    	Amount of spaces to indent
# Example
indent --filename ~/tmp/logstash/syslog/values.yaml --indent 8 >> templates/syslog.yaml
```

## PreBuilt Binaries
Grab Binaries from [The Releases Page](https://github.com/Jmainguy/indent/releases)

## Build
```/bin/bash
export GO111MODULE=on
go build
```
