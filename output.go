// Package xlognsq is an xlog to NSQd output for github.com/rs/xlog
package xlognsq

import "encoding/json"

// Publisher is a specialized interface to nsq.Producer
type Publisher interface {
	Publish(topic string, body []byte) error
}

// Output is a xlog to nsqd output
type Output struct {
	topic string
	pub   Publisher
}

// New returns an xlognsq.Output ready to publish to the given
// topic with the provided NSQ publisher
func New(topic string, pub Publisher) *Output {
	return &Output{
		topic: topic,
		pub:   pub,
	}
}

// Write implements xlog.Output interface
func (o Output) Write(fields map[string]interface{}) error {
	b, err := json.Marshal(fields)
	if err != nil {
		return err
	}
	return o.pub.Publish(o.topic, b)
}
