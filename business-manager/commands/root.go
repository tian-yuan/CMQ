package commands

import (
	"github.com/spf13/cobra"
	"github.com/tian-yuan/CMQ/business-manager/svc"
)

func init() {
}

var rootCmd = &cobra.Command{
	Use: "cli",
	Run: func(cmd *cobra.Command, args []string) {
		var tracerAddr string
		cmd.Flags().StringVarP(&tracerAddr, "tracerAddress", "j", "127.0.0.1:6831", "tracer address array")

		httpconf := svc.NewH2cConf()
		cmd.Flags().StringVarP(&httpconf.Host, "HttpHost", "t", "0.0.0.0", "http2 bind host address.")
		cmd.Flags().Uint16VarP(&httpconf.Port, "HttpPort", "r", 18080, "http2 hub bind port.")
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
