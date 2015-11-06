# XLog to NSQ Output

[![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/rs/xlog-nsq) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/rs/xlog-nsq/master/LICENSE) [![Build Status](https://travis-ci.org/rs/xlog-nsq.svg?branch=master)](https://travis-ci.org/rs/xlog-nsq) [![Coverage](http://gocover.io/_badge/github.com/rs/xlog-nsq)](http://gocover.io/github.com/rs/xlog-nsq)

xlog-nsq is an xlog to [NSQd](http://nsq.io) output for [github.com/rs/xlog](https://github.com/rs/xlog).

## Install

    go get github.com/rs/xlog-nsq

## Usage

```go
c := xhandler.Chain{}

nsqCfg := nsq.NewConfig()
nsqTopic := "mytopic"
nsqAddr := "127.0.0.1:4150"
p := nsq.NewProducer(nsqAddr, nsqCfg)
o := xlognsq.New(nsqTopic, p)

logCfg := xlog.Config{
    Output: xlog.NewOutputChannel(o),
}

c.UseC(xlog.NewHandler(logCfg))
```

## Licenses

All source code is licensed under the [MIT License](https://raw.github.com/rs/xlog-nsq/master/LICENSE).
