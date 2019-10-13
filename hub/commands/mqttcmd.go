package commands

import (
	"github.com/spf13/cobra"
	"github.com/sirupsen/logrus"

	"github.com/tian-yuan/CMQ/hub/svc"
	"github.com/tian-yuan/iot-common/util"
	"strings"
	"time"
	"fmt"
)

var mqttCmd = &cobra.Command{
	Use:   "mqtt",
	Short: "start mqtt server",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("Start mqtt hub gateway v0.0.1 -- HEAD")

		var zkAddr string
		cmd.Flags().StringVarP(&zkAddr, "zkAddress", "z", "127.0.0.1:2181", "zk address array")
		logrus.Infof("start discovery client, zk address : %s", zkAddr)
		zkAddrArr := strings.Split(zkAddr, ";")
		util.Ctx.InitRegisterSvc(zkAddrArr)
		util.Ctx.InitPubEngineSvc(zkAddrArr)
		util.Ctx.InitMessageDispatcherSvc(zkAddrArr)

		var redisClusterAddr string
		cmd.Flags().StringVarP(&redisClusterAddr, "redisClusterAddr", "r", "127.0.0.1:7000", "redis cluster address")
		var reqTimeout time.Duration
		cmd.Flags().DurationVarP(&reqTimeout, "redisReqTimeout", "t", 5 * time.Second, "redis request timeout")
		var poolSize int32
		cmd.Flags().Int32VarP(&poolSize, "redisPoolSize", "s", 5, "redis client pool size")
		var redisSessionTimeout time.Duration
		cmd.Flags().DurationVarP(&redisSessionTimeout, "redisSessionTimeout", "o", 20 * time.Minute, "redis session timeout")
		var redisSessionRefresh time.Duration
		cmd.Flags().DurationVarP(&redisSessionRefresh, "redisSessionRefresh", "f", 20 * time.Minute, "redis session refresh")
		logrus.Infof("create redis cluster client : %s", redisClusterAddr)
		redisClient := util.GetClusterClient(redisClusterAddr, reqTimeout, int(poolSize))
		svc.Global.RedisClient = redisClient
		ss := util.NewRedisSessionStorage(redisClient)
		svc.Global.SessionStorage = ss
		svc.Global.ReqTimeOut = reqTimeout
		svc.Global.RedisSessionTimeOut = redisSessionTimeout
		svc.Global.RedisSessionRefresh = redisSessionRefresh

		conf := svc.NewMqttConf()
		cmd.Flags().StringVarP(&conf.MqttHost, "mqttHost", "m", "0.0.0.0", "mqtt hub bind host address.")
		cmd.Flags().Uint16VarP(&conf.MqttPort, "mqttPort", "p", 1883, "mqtt hub bind port.")
		mqttSvc := svc.NewMqttSvc(conf)
		mqttSvc.Start()

		httpconf := svc.NewH2cConf()
		cmd.Flags().StringVarP(&httpconf.Host, "Host", "a", "0.0.0.0", "http2 bind host address.")
		cmd.Flags().Uint16VarP(&httpconf.Port, "Port", "b", 9883, "http2 hub bind port.")
		svc.Global.SessionPrefix = httpconf.Host + ":" + fmt.Sprintf("%d", httpconf.Port)
		logrus.Infof("session prefix : %s", svc.Global.SessionPrefix)
		h2cSvc := svc.NewH2cSvc(httpconf)
		h2cSvc.Start()
	},
}
