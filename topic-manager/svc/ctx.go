package svc

import (
	"github.com/sirupsen/logrus"
	"github.com/tian-yuan/CMQ/topic-manager/topic"
)

var Ctx Context

func init() {
	Ctx.m = topic.NewCSTrieMatcher()
}

type Context struct {
	m topic.Matcher
}

func (ctx *Context) Subscribe(topic string, qos int, guid uint32) error {
	_, err := ctx.m.Subscribe(topic, guid)
	if err != nil {
		logrus.Infof("subscribe topic %s failed.", topic)
	}
	return err
}

func (ctx *Context) Match(topic string) []topic.Subscriber {
	return ctx.m.Lookup(topic)
}
