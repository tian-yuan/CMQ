package svc

import (
	"context"
	proto "github.com/tian-yuan/CMQ/iotpb"
)

type rpchandler struct {

}

func (h *rpchandler) PublishMessage(ctx context.Context, in *proto.PublishMessageRequest, out *proto.PublishMessageResponse) error {
	return nil
}

func (h *rpchandler) Subscribe(ctx context.Context, in *proto.SubscribeMessageRequest, out *proto.SubscribeMessageResponse) error {
	return nil
}
