package svc

import (
	"context"
	proto "github.com/tian-yuan/CMQ/iotpb"
)

type rpchandler struct {

}

func (h* rpchandler) Registry(context.Context, *proto.ConnectMessageRequest, *proto.ConnectMessageResponse) error {
	return nil
}



