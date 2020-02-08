package commands

import (
	"github.com/micro/go-micro/util/log"
	"github.com/spf13/cobra"
	"github.com/tian-yuan/CMQ/business-manager/svc"
)

func init() {
}

var rootCmd = &cobra.Command{
	Use: "cli",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("start business manager.")
		var tracerAddr string
		cmd.Flags().StringVarP(&tracerAddr, "tracerAddress", "j", "127.0.0.1:6831", "tracer address array")

		mc := svc.NewMqttClient()
		mc.Start()

		httpconf := svc.NewH2cConf()
		httpconf.Host = svc.GlobalHttpConfig.Address
		httpconf.Port = svc.GlobalHttpConfig.Port
		svc.Global.HttpSvc = svc.NewH2cSvc(httpconf)
		svc.Global.HttpSvc.Start(tracerAddr)
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func Stop() {
	svc.Global.HttpSvc.Stop()
}
