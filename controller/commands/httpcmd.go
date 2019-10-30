package commands

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/tian-yuan/CMQ/controller/svc"
)

var httpCmd = &cobra.Command{
	Use:   "controller",
	Short: "start controller server",
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("Start controller server v0.0.1 -- HEAD")

		dbconf := svc.NewDatabaseConfig()
		cmd.Flags().StringVarP(&dbconf.Host, "SqlHost", "m", "127.0.0.1", "mysql host address")
		cmd.Flags().Uint16VarP(&dbconf.Port, "SqlPort", "p", 3306, "mysql port")
		cmd.Flags().StringVarP(&dbconf.Username, "Username", "u", "root", "mysql username")
		cmd.Flags().StringVarP(&dbconf.Password, "password", "w", "123456", "mysql password")
		cmd.Flags().StringVarP(&dbconf.Database, "database", "d", "push_device", "mysql database")
		dbsvc := svc.NewDatabaseSvc(dbconf)
		svc.Ctx.Dbsvc = dbsvc
		dbsvc.Start()

		httpconf := svc.NewH2cConf()
		cmd.Flags().StringVarP(&httpconf.Host, "HttpHost", "t", "0.0.0.0", "http2 bind host address.")
		cmd.Flags().Uint16VarP(&httpconf.Port, "HttpPort", "r", 18080, "http2 hub bind port.")
		h2cSvc := svc.NewH2cSvc(httpconf)
		h2cSvc.Start()
	},
}
