package svc

import (
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/micro/go-micro/util/log"
)

type MqttClient interface {
	IsStarted() bool
	Start()
	Stop()
}

type mqttClient struct {
	opts *mqtt.ClientOptions

	stopCh chan struct{}
}

func NewMqttClient() MqttClient {
	mc := &mqttClient{}
	mc.opts = mqtt.NewClientOptions()
	mc.opts.SetAutoReconnect(true)
	mc.opts.SetClientID(GlobalMqttConfig.DeviceId)
	mc.opts.AddBroker(GlobalMqttConfig.BrokerAddress)
	mc.opts.SetUsername(GlobalMqttConfig.DeviceId)
	mc.opts.SetPassword(GlobalMqttConfig.Password)
	mc.opts.SetDefaultPublishHandler(mc.handleMessage)
	return mc
}

func (mc *mqttClient) Start() {
	log.Info("begin to start mqtt client.")
	if mc.IsStarted() {
		// already started, so return
		log.Infof("mqtt client is already started.")
		return
	}
	mc.stopCh = make(chan struct{})
	go mc.run()
}

func (mc *mqttClient) Stop() {
	log.Info("stop mqtt client.")
	if mc.stopCh != nil {
		close(mc.stopCh)
	}
}

func (mc *mqttClient) IsStarted() bool {
	if mc.stopCh == nil {
		return false
	}
	select {
	case <-mc.stopCh:
		return false
	default:
		return true
	}
}

func (mc *mqttClient) run() {
	for {
		select {
		case <-mc.stopCh:
			return
		default:
		}
		// do work
		c := mqtt.NewClient(mc.opts)
		if token := c.Connect(); token.Wait() && token.Error() != nil {
			log.Infof("connect to mqtt broker error : %s", token.Error())
			time.Sleep(10 * time.Second)
			go mc.run()
			return
		}
		// begin subscribe all topic
		// {product_key}/{device_id}/action
		// just subscribe the message from device
		topic := GlobalMqttConfig.ProductKey + "/device/#"
		log.Infof("subscribe mqtt broker, topic : %s", topic)
		if wsubToken := c.Subscribe(topic, 1, nil); wsubToken.Wait() && wsubToken.Error() != nil {
			log.Infof("subscribe to mqtt broker for topic : %s error, message : %s", topic, wsubToken.Error())
		}
		select {
		case <-mc.stopCh:
			c.Disconnect(1000)
		}
	}
}

func (mc *mqttClient) handleMessage(client mqtt.Client, msg mqtt.Message) {
	log.Infof("msg topic : %s, payload : %s", msg.Topic(), msg.Payload())
	kv := strings.Split(msg.Topic(), "/")
	if len(kv) < 4 {
		log.Infof("unhandle topic : %s", msg.Topic())
		return
	}
	productKey := kv[0]
	deviceId := kv[2]
	msgType := kv[3]
	log.Infof("handle message, product key : %s, device id : %s, msg type : %s", productKey, deviceId, msgType)
	if productKey != GlobalMqttConfig.ProductKey {
		log.Infof("unhandle product, key : %s", productKey)
		return
	}
	switch MsgType(msgType) {
	case Billing:
		OnBillingMessage(client, msg)
	case Monitor:
		OnMonitorMessage(client, msg)
	case Manager:
		OnManagerMessage(client, msg)
	default:
		log.Infof("unhandled message type : %s", msgType)
	}
}
