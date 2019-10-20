package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tian-yuan/CMQ/message-dispatcher/svc"
)

func init()  {
	rootCmd.AddCommand(httpCmd)
}

var rootCmd = &cobra.Command{
	Use: "cli",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("OK")
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func Stop() {
	svc.Global.TopicLoadSvc.Stop()
}