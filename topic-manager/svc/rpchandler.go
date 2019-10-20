package svc

import (
	"context"
	proto "github.com/tian-yuan/iot-common/iotpb"
	"github.com/sirupsen/logrus"
	"strings"
	"fmt"
	"net/url"
	"net/http"
)

type rpchandler struct {

}

const hubPublishPath = "/v1/imd/publish"

func publishMessageToHub(info string, message string, topic string, qos int8) error {
	logrus.Infof("publish message to hub, info : %s, message : %s, topic : %s, qos : %d",
		info, message, topic, qos)
	// split session info
	infos := strings.Split(info, ":")
	addr := infos[0]
	port := infos[1]
	fd := infos[2]
	cli := GetH2client(addr + ":" + port)
	u := &url.URL{
		Scheme: "http",
		Host:   addr + ":" + port,
		Path:   hubPublishPath,
	}

	query := url.Values{}
	query.Add("fd", fd)
	query.Add("message", message)
	query.Add("topic", topic)
	query.Add("qos", fmt.Sprintf("%d", qos))
	u.RawQuery = query.Encode()
	logrus.Infof("publish message to hub : %s", u.String())
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		logrus.Errorf("Construct http request error")
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	_, err = cli.Do(req)
	if err != nil {
		logrus.Errorf("publish message to hub failed.")
		return nil
	}
	logrus.Infof("publish message to hub success.")
	return nil
}

func (h* rpchandler) PublishMessage(ctx context.Context, in *proto.PublishMessageRequest, out *proto.PublishMessageResponse) error {
	logrus.Infof("topic manager receive publish message.")
	subs := Ctx.Match(in.Topic)
	for _, sub := range subs {
		// get guid from subs
		guid := sub.(uint32)
		logrus.Infof("device guid match topic : %s, guid : %d", in.Topic, guid)
		// get session info from redis by guid
		r := Global.RedisClient.Get("C:" + fmt.Sprintf("%d", guid))
		if err := r.Err(); err != nil {
			logrus.Error(err, "get session info from redis failed, guid : " + string(guid))
			continue
		}
		// send request to hub
		val := r.Val()
		publishMessageToHub(val, string(in.Payload), in.Topic, int8(in.Header.Qos))
	}
	out.Code = 200
	out.Message = "publish success."
	return nil
}

func (h* rpchandler) LoadSubTopic(ctx context.Context, in *proto.SubTopicLoadRequest, out *proto.SubTopicLoadResponse) error {
	logrus.Infof("handle load sub topic request, product key : %s.", in.ProductKey)
	return nil
}

func (h* rpchandler) Subscribe(ctx context.Context, in *proto.SubscribeMessageRequest, out *proto.SubscribeMessageResponse) error {
	logrus.Infof("handler subscribe message, topic : %s", in.TopicFilter)
	err := Ctx.Subscribe(in.TopicFilter, int(in.Qos), in.Guid)
	if err != nil {
		logrus.Errorf("subscribe topic failed, topic : %s", in.TopicFilter)
	}
	// should write to database
	out.Code = 200
	out.Message = "subscribe success."
	logrus.Infof("handler subscribe message success, topic : %s", in.TopicFilter)
	return nil
}

func (h* rpchandler) UnSubscribe(ctx context.Context, in *proto.UnSubscribeMessageRequest, out *proto.UnSubscribeMessageResponse) error {
	logrus.Info("handle unsubscribe message request.")
	return nil
}
