package svc

import "github.com/sirupsen/logrus"

type MqttSvc struct {
	Conf *MqttConf
}

type MqttConf struct {
	MqttHost string
	MqttPort uint16
}

func NewMqttConf() *MqttConf {
	return &MqttConf{
		MqttHost: "0.0.0.0",
		MqttPort: 1883,
	}
}

func NewMqttSvc(conf *MqttConf) *MqttSvc {
	return &MqttSvc {
		Conf: conf,
	}
}

func (ms *MqttSvc) StartMqttSvc() {
	logrus.WithFields(logrus.Fields{
		"mqttHost": ms.Conf.MqttHost,
		"mqttPort": ms.Conf.MqttPort,
	}).Info("start mqtt server.")
	return
}
