package svc

import (
	"github.com/micro/go-micro"
	proto "github.com/tian-yuan/iot-common/iotpb"
	"fmt"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/zookeeper"
	"github.com/tian-yuan/iot-common/util"
	"time"
)

type RpcSvc struct {

}

func NewRpcSvc() *RpcSvc {
	return &RpcSvc{}
}

func (svc *RpcSvc) Start(zkAddr []string) {
	optFunc := func(opt *registry.Options) {
		opt = &registry.Options {
			Addrs: zkAddr,
		}
	}
	registry := zookeeper.NewRegistry(optFunc)

	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name(util.MESSAGE_DISPATCHER_SVC),
		micro.Registry(registry),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	proto.RegisterMessageDispatcherHandler(service.Server(), new(rpchandler))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
