package svc

import (
	"context"
	proto "github.com/tian-yuan/iot-common/iotpb"
	"github.com/sirupsen/logrus"
	"strings"
	"fmt"
)

type rpchandler struct {

}

func publishMessageToHub(info string, message string) error {
	// split session info
	infos := strings.Split(info, ":")
	addr := infos[0]
	port := infos[1]
	fd := infos[2]
	cli := GetH2client(addr + ":" + port)
	path := fmt.Sprintf("http://%s:%s?fd=%s&message=%s", addr, port, fd, message)
	_, err := cli.Get(path)
	if err != nil {
		logrus.Errorf("publish message to hub failed.")
		return nil
	}
	logrus.Infof("publish message to hub success.")
	return nil
}

func (h* rpchandler) PublishMessage(ctx context.Context, in *proto.PublishMessageRequest, out *proto.PublishMessageResponse) error {
	subs := Ctx.Match(in.Topic)
	for _, sub := range subs {
		// get guid from subs
		guid := sub.(uint32)
		logrus.Infof("device guid match topic : %s, guid : %d", in.Topic, guid)
		// get session info from redis by guid
		r := Global.RedisClient.Get("C:" + string(guid))
		if err := r.Err(); err != nil {
			continue
		}
		// send request to hub
		val := r.Val()
		publishMessageToHub(val, string(in.Payload))
	}
	out.Code = 200
	out.Message = "publish success."
	return nil
}

func (h* rpchandler) LoadSubTopic(ctx context.Context, in *proto.SubTopicLoadRequest, out *proto.SubTopicLoadResponse) error {
	return nil
}

func (h* rpchandler) Subscribe(ctx context.Context, in *proto.SubscribeMessageRequest, out *proto.SubscribeMessageResponse) error {
	err := Ctx.Subscribe(in.TopicFilter, int(in.Qos), in.Guid)
	if err != nil {
		logrus.Errorf("subscribe topic failed, topic : %s", in.TopicFilter)
	}
	// should write to database
	out.Code = 200
	out.Message = "subscribe success."
	return nil
}

func (h* rpchandler) UnSubscribe(ctx context.Context, in *proto.UnSubscribeMessageRequest, out *proto.UnSubscribeMessageResponse) error {
	return nil
}
