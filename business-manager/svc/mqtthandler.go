package svc

import (
	"encoding/json"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/micro/go-micro/util/log"
)

type MsgType string

const (
	Billing   MsgType = "billing"
	Monitor   MsgType = "monitor"
	Parameter MsgType = "parameter"
	Action    MsgType = "action"
	Manager   MsgType = "manager"
)

type ActionType string

const (
	StartCharge    ActionType = "StartCharge"
	StopCharge     ActionType = "StopCharge"
	SetParameter   ActionType = "SetParameter"
	RealTimePacket ActionType = "RealTimePacket"
	ShutDown       ActionType = "ShutDown"
	Adjust         ActionType = "Adjust"
	Register       ActionType = "Register"
)

type RegisterMsg struct {
	SeqId         uint32 `json:"SeqId"`
	Action        string `json:"Action"`
	MachineId     string `json:"MachineId"`
	MachineType   string `json:"MachineType"`
	ChannelCounts uint32 `json:"ChannelCounts"`
	Version       string `json:"Version"`
}

type ParameterMsg struct {
	SeqId   uint32 `json:"SeqId"`
	Action  string `json:"Action"`
	Param1  string `json:"Param1"`
	Date    string `json:"Date"`
	Version string `json:"Version"`
}

type RealTimePacketMsg struct {
	SeqId   uint32 `json:"SeqId"`
	Action  string `json:"Action"`
	Param1  string `json:"Param1"`
	Version string `json:"Version"`
}

type BillingStartChargeMsg struct {
	SeqId           uint32 `json:"SeqId"`
	Action          string `json:"Action"`
	CardId          string `json:"CardId"`
	PhoneNumber     string `json:"PhoneNumber"`
	ConsumePassword string `json:"ConsumePassword"`
	Channel         uint32 `json:"Channel"`
	Version         string `json:"Version"`
}

type BillingStopChargeMsg struct {
	SeqId        uint32  `json:"SeqId"`
	Action       string  `json:"Action"`
	OrderType    int     `json:"OrderType"`
	ConsumeMoney float32 `json:"ConsumeMoney"`
	ConsumeType  string  `json:"ConsumeType"`
	Channel      uint32  `json:"Channel"`
	Idempotence  uint32  `json:"Idempotence"`
	Date         string  `json:"Date"`
	TradeId      string  `json:"TradeId"`
	Version      string  `json:"Version"`
}

func OnBillingMessage(client mqtt.Client, msg mqtt.Message) {
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(msg.Payload()), &dat); err == nil {
		log.Infof("unhandled body : %s", msg.Payload())
		return
	}
	seqId := dat["SeqId"]
	action := dat["Action"]
	log.Infof("on manager message, seq id : %s, action : %s", seqId, action)
	switch action {
	case StartCharge:
		onBillingStartCharge(client, msg)
	case StopCharge:
		onBillingStopCharge(client, msg)
	default:
		log.Infof("unhandled msg type action :%s .", action)
	}
}

func onBillingStartCharge(client mqtt.Client, msg mqtt.Message) {
}

func onBillingStopCharge(client mqtt.Client, msg mqtt.Message) {
}

func OnMonitorMessage(client mqtt.Client, msg mqtt.Message) {
}

func OnManagerMessage(client mqtt.Client, msg mqtt.Message) {
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(msg.Payload()), &dat); err != nil {
		log.Infof("unhandled body : %s, error : %s", msg.Payload(), err.Error())
		return
	}
	seqId := dat["SeqId"]
	action := dat["Action"]
	log.Infof("on manager message, seq id : %d, action : %s", seqId, action)
	switch ActionType(action.(string)) {
	case Register:
		onRegister(client, msg)
	default:
		log.Infof("unhandled msg type action : %s", action)
	}
}

func onRegister(client mqtt.Client, msg mqtt.Message) {
	// publish SetParameter message to device
	parameter := ParameterMsg{
		SeqId:   1,
		Action:  "SetParameter",
		Param1:  "test",
		Date:    time.Now().Format(time.RFC3339),
		Version: "0.1",
	}
	kv := strings.Split(msg.Topic(), "/")
	if len(kv) < 4 {
		log.Infof("unhandle topic : %s", msg.Topic)
		return
	}
	productKey := kv[0]
	deviceId := kv[2]
	topic := productKey + "/controller/" + deviceId + "/action"
	data, _ := json.Marshal(parameter)
	client.Publish(topic, 1, false, string(data))
}
