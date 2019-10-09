package svc

import (
	"context"
	proto "github.com/tian-yuan/iot-common/iotpb"
	"github.com/sirupsen/logrus"
	"github.com/tian-yuan/iot-common/util"
)

type rpchandler struct {

}

func (h *rpchandler) PublishMessage(ctx context.Context, in *proto.PublishMessageRequest, out *proto.PublishMessageResponse) error {
	logrus.Infof("message dispatcher service receive publish message, topic: %s, payload : %s",
		in.Topic, string(in.Payload))
	topicMangerCli := proto.NewTopicManagerService(util.TOPIC_MANAGER_SVC,
		util.Ctx.TopicManagerSvc.Client())
	out, err := topicMangerCli.PublishMessage(context.TODO(), in)
	if err != nil {
		logrus.Error("publish to topic manager failed.")
	}
	return err
}

func (h *rpchandler) Subscribe(ctx context.Context, in *proto.SubscribeMessageRequest, out *proto.SubscribeMessageResponse) error {
	// send subscribe message to iot topic manager
	topicMangerCli := proto.NewTopicManagerService(util.TOPIC_MANAGER_SVC,
		util.Ctx.TopicManagerSvc.Client())
	out, err := topicMangerCli.Subscribe(context.TODO(), in)
	if err != nil {
		logrus.Error("subscribe to topic manager failed.")
	}
	return err
}

func (h *rpchandler) UnSubscribe(ctx context.Context, in *proto.UnSubscribeMessageRequest, out *proto.UnSubscribeMessageResponse) error {
	topicMangerCli := proto.NewTopicManagerService(util.TOPIC_MANAGER_SVC,
		util.Ctx.TopicManagerSvc.Client())
	out, err := topicMangerCli.UnSubscribe(context.TODO(), in)
	if err != nil {
		logrus.Error("unsubscribe to topic manager failed.")
	}
	return err
}