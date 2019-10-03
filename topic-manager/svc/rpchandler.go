package svc

import (
	"context"
	proto "github.com/tian-yuan/CMQ/iotpb"
)

type rpchandler struct {

}

func (h* rpchandler) PublishMessage(ctx context.Context, in *proto.PublishMessageRequest, out *proto.PublishMessageResponse) error {
	return nil
}

func (h* rpchandler) LoadSubTopic(ctx context.Context, in *proto.SubTopicLoadRequest, out *proto.SubTopicLoadResponse) error {
	return nil
}

