package commands

import (
	"github.com/spf13/cobra"
	"github.com/sirupsen/logrus"

	"github.com/tian-yuan/CMQ/hub/svc"
)

var mqttCmd = &cobra.Command{
	Use:   "mqtt",
	Short: "start mqtt server",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("Start mqtt hub gateway v0.0.1 -- HEAD")
		conf := svc.NewMqttConf()
		cmd.Flags().StringVarP(&conf.MqttHost, "mqttHost", "m", "0.0.0.0", "mqtt hub bind host address.")
		cmd.Flags().Uint16VarP(&conf.MqttPort, "mqttPort", "p", 1883, "mqtt hub bind port.")
		mqttSvc := svc.NewMqttSvc(conf)
		mqttSvc.Start()
	},
}
