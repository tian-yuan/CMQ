package svc

import (
	"context"
	proto "github.com/tian-yuan/iot-common/iotpb"
	"github.com/sirupsen/logrus"
)

type rpchandler struct {

}

func (h *rpchandler) PublishMessage(ctx context.Context, in *proto.PublishMessageRequest, out *proto.PublishMessageResponse) error {
	logrus.Infof("message dispatcher service receive publish message, topic: %s, payload : %s",
		in.Topic, string(in.Payload))

	err := Global.TopicLoadSvc.PublishMessage(in, out)
	if err != nil {
		logrus.Errorf("publish to topic manager failed, error : %v.", err)
	}
	return err
}

func (h *rpchandler) Subscribe(ctx context.Context, in *proto.SubscribeMessageRequest, out *proto.SubscribeMessageResponse) error {
	// send subscribe message to iot topic manager
	err := Global.TopicLoadSvc.Subscribe(in, out)
	if err != nil {
		logrus.Errorf("subscribe to topic manager failed, err : %v.", err)
	}
	return err
}

func (h *rpchandler) UnSubscribe(ctx context.Context, in *proto.UnSubscribeMessageRequest, out *proto.UnSubscribeMessageResponse) error {
	err := Global.TopicLoadSvc.UnSubscribe(in, out)
	if err != nil {
		logrus.Errorf("unsubscribe to topic manager failed, err : %v.", err)
	}
	return err
}
