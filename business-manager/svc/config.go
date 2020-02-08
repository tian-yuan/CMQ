package svc

import (
	"fmt"
	"os"
	"strings"

	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/util/log"
)

type MqttConfig struct {
	BrokerAddress string `json:"BrokerAddress"`
	ProductKey    string `json:"ProductKey"`
	DeviceId      string `json:"DeviceId"`
	Password      string `json:"Password"`
}

type HttpConfig struct {
	Address string `json:"Address"`
	Port    uint16 `json:"Port"`
}

var GlobalMqttConfig MqttConfig
var GlobalHttpConfig HttpConfig

func InitConfig() error {
	return loadConfig()
}

func loadConfig() error {
	pathName, err := os.Executable()
	if err != nil {
		panic(fmt.Errorf("get executable path failed : %s\n", err))
		return nil
	}
	index := strings.LastIndex(pathName, "/")
	path := string(pathName[0:index])
	configFile := path + "/conf/config.json"

	log.Infof("init config from config file : %s", configFile)
	err = config.LoadFile(configFile)
	if err != nil {
		log.Infof("load config file failed, error : %s", err.Error())
		return err
	}

	err = config.Get("mqtt").Scan(&GlobalMqttConfig)
	if err != nil {
		log.Infof("scan mqtt config failed, error : %s", err.Error())
		return err
	}

	err = config.Get("http").Scan(&GlobalHttpConfig)
	if err != nil {
		log.Infof("scan http config failed, error : %s", err.Error())
		return err
	}

	log.Infof("mqtt broker address : %s, product key : %s, device id : %s, password : %s",
		GlobalMqttConfig.BrokerAddress, GlobalMqttConfig.ProductKey, GlobalMqttConfig.DeviceId,
		GlobalMqttConfig.Password)
	log.Infof("http address : %s, http port : %d", GlobalHttpConfig.Address, GlobalHttpConfig.Port)
	return nil
}
