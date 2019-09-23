package commands

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/CMQ/controller/svc"
)

var httpCmd = &cobra.Command{
	Use:   "controller",
	Short: "start controller server",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("Start controller server v0.0.1 -- HEAD")
		conf := svc.NewH2cConf()
		cmd.Flags().StringVarP(&conf.Host, "Host", "m", "0.0.0.0", "http2 bind host address.")
		cmd.Flags().Uint16VarP(&conf.Port, "Port", "p", 9883, "http2 hub bind port.")
		h2cSvc := svc.NewH2cSvc(conf)
		h2cSvc.Start()
	},
}
