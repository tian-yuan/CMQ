package svc

import (
	"github.com/micro/go-micro"
	proto "github.com/tian-yuan/CMQ/iotpb"
	"fmt"
	"github.com/tian-yuan/CMQ/util"
)

type RpcSvc struct {

}

func NewRpcSvc() *RpcSvc {
	return &RpcSvc{}
}

func (svc *RpcSvc) Start() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name(util.PUBLISH_ENGINE_SVC),
	)

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	proto.RegisterPublishEngineHandler(service.Server(), new(rpchandler))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
