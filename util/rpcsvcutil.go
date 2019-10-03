package util

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/zookeeper"
)

const (
	REGISTER_MANAGER_SVC = "iot.register.manager"
	TOPIC_MANAGER_SVC = "iot.topic.manager"
	MESSAGE_DISPATCHER_SVC = "iot.message.dispatcher"
	TOPIC_ACL_SVC = "iot.topic.acl"
	PUBLISH_ENGINE_SVC = "iot.publish.engine"
	CONTROLLER_SVC = "iot.controller"
)

type RpcCtx struct {
	// The following fields is rpc service variable
	// {{
	RegisterSvc micro.Service
	TopicManagerSvc micro.Service
	MessageDispatcherSvc micro.Service
	PubEngineSvc micro.Service
	TopicAclSvc micro.Service
	ControllerSvc micro.Service
	// }}
}

var Ctx RpcCtx

func (ctx* RpcCtx) InitRegisterSvc(zkAddr []string) {
	optFunc := func(opt *registry.Options) {
		opt = &registry.Options {
			Addrs: zkAddr,
		}
	}
	registry := zookeeper.NewRegistry(optFunc)
	Ctx.RegisterSvc = micro.NewService(
		micro.Name(REGISTER_MANAGER_SVC + ".client"),
		micro.Registry(registry),
	)
	Ctx.RegisterSvc.Init()
}

func (ctx* RpcCtx) InitPubEngineSvc(zkAddr []string) {
	optFunc := func(opt *registry.Options) {
		opt = &registry.Options {
			Addrs: zkAddr,
		}
	}
	registry := zookeeper.NewRegistry(optFunc)
	Ctx.PubEngineSvc = micro.NewService(
		micro.Name(PUBLISH_ENGINE_SVC + ".client"),
		micro.Registry(registry),
	)
	Ctx.PubEngineSvc.Init()
}

func (ctx* RpcCtx) InitTopicManagerSvc(zkAddr []string) {
	optFunc := func(opt *registry.Options) {
		opt = &registry.Options {
			Addrs: zkAddr,
		}
	}
	registry := zookeeper.NewRegistry(optFunc)
	Ctx.TopicManagerSvc = micro.NewService(
		micro.Name(TOPIC_MANAGER_SVC + ".client"),
		micro.Registry(registry),
	)
	Ctx.TopicManagerSvc.Init()
}

func (ctx* RpcCtx) InitTopicAclSvc(zkAddr []string) {
	optFunc := func(opt *registry.Options) {
		opt = &registry.Options {
			Addrs: zkAddr,
		}
	}
	registry := zookeeper.NewRegistry(optFunc)
	Ctx.TopicAclSvc = micro.NewService(
		micro.Name(TOPIC_ACL_SVC + ".client"),
		micro.Registry(registry),
	)
	Ctx.TopicAclSvc.Init()
}

func (ctx* RpcCtx) InitMessageDispatcherSvc(zkAddr []string) {
	optFunc := func(opt *registry.Options) {
		opt = &registry.Options {
			Addrs: zkAddr,
		}
	}
	registry := zookeeper.NewRegistry(optFunc)
	Ctx.MessageDispatcherSvc = micro.NewService(
		micro.Name(MESSAGE_DISPATCHER_SVC + ".client"),
		micro.Registry(registry),
	)
	Ctx.MessageDispatcherSvc.Init()
}

func (ctx* RpcCtx) InitControllerSvc(zkAddr []string) {
	optFunc := func(opt *registry.Options) {
		opt = &registry.Options {
			Addrs: zkAddr,
		}
	}
	registry := zookeeper.NewRegistry(optFunc)
	Ctx.ControllerSvc = micro.NewService(
		micro.Name(CONTROLLER_SVC + ".client"),
		micro.Registry(registry),
	)
	Ctx.ControllerSvc.Init()
}
