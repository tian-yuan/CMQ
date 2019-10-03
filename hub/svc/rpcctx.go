package svc

import (
	"github.com/micro/go-micro"
	"github.com/tian-yuan/CMQ/util"
)

type RpcCtx struct {
	// The following fields is rpc service variable
	// {{
	registerSvc micro.Service
	pubEngineSvc micro.Service
	// }}
}

var rpcctx RpcCtx

func init() {
	rpcctx.registerSvc = micro.NewService(micro.Name(util.REGISTER_MANAGER_SVC + "-client"))
	rpcctx.registerSvc.Init()

	rpcctx.pubEngineSvc = micro.NewService(micro.Name(util.PUBLISH_ENGINE_SVC + "-client"))
	rpcctx.pubEngineSvc.Init()
}
