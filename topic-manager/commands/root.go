package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init()  {
	rootCmd.AddCommand(mqttCmd)
	rootCmd.AddCommand(coapCmd)
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

}