package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init()  {
	rootCmd.AddCommand(httpCmd)
	rootCmd.AddCommand(rpccmd)
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