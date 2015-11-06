package xlognsq

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type pub struct {
	lastTopic string
	lastBody  []byte
}

func (p *pub) Publish(topic string, body []byte) error {
	p.lastTopic = topic
	p.lastBody = body
	return nil
}

func TestNew(t *testing.T) {
	p := &pub{}
	o := New("test", p)
	assert.Equal(t, p, o.pub)
	assert.Equal(t, "test", o.topic)
}

func TestWrite(t *testing.T) {
	p := &pub{}
	o := New("test", p)
	err := o.Write(map[string]interface{}{"foo": "bar"})
	assert.NoError(t, err)
	assert.Equal(t, "test", p.lastTopic)
	assert.Equal(t, "{\"foo\":\"bar\"}", string(p.lastBody))
}
