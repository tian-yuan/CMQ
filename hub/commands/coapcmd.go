package commands

import (
	"github.com/spf13/cobra"

	"github.com/micro/go-micro/util/log"
	"github.com/tian-yuan/CMQ/hub/svc"
)

var coapCmd = &cobra.Command{
	Use:   "coap",
	Short: "start coap server",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Start coap hub gateway v0.0.1 -- HEAD")
		conf := svc.NewCoapConf()
		cmd.Flags().StringVarP(&conf.Type, "coapType", "t", "udp", "coap hub bind address type.")
		cmd.Flags().StringVarP(&conf.CoapHost, "coapHost", "m", "0.0.0.0", "coap hub bind host address.")
		cmd.Flags().Uint16VarP(&conf.CoapPort, "coapPort", "p", 5683, "coap hub bind port.")
		coapSvc := svc.NewCoapSvc(conf)
		coapSvc.Start()
	},
}
