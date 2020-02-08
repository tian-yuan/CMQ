
import (
	"encoding/json"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/micro/go-micro/util/log"
)

type MsgType string

const (
	Billing   MsgType = "billing"
	Monitor   MsgType = "monitor"
	Parameter MsgType = "parameter"
)

type ActionType string

const (
	StartCharge    ActionType = "StartCharge"
	StopCharge     ActionType = "StopCharge"
	SetParameter   ActionType = "SetParameter"
	RealTimePacket ActionType = "RealTimePacket"
	ShutDown       ActionType = "ShutDown"
	Adjust         ActionType = "Adjust"
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
	if err := json.Unmarshal([]byte(mqtt.Payload), &dat); err == nil {
		log.Infof("unhandled body : %s", mqtt.Payload)
		return
	}
	seqId := dat["SeqId"]
	action := dat["Action"]
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
