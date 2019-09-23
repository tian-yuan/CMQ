package svc

import (
	"github.com/CMQ/topic-manager/topic"
	"github.com/sirupsen/logrus"
)

var ctx Context

func init() {
}

type Context struct {
	m topic.Matcher
}

func (ctx *Context) subscribe(topic string, qos int, guid uint32) error {
	_, err := m.Subscribe(topic, guid)
	if err != nil {
		logrus.Infof("subscribe topic %s failed.", topic)
	}
	return err
}

func (ctx *Context) match(topic string) []Subscriber {
	return m.Lookup(topic)
}
