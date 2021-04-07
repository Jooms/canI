package cmd

import (
	"fmt"
	"os"

	"github.com/Jooms/canI/pkg/cmd/action"

	"github.com/spf13/cobra"
)

var cfgFile string

var RootCmd = &cobra.Command{
	Use:   "canI",
	Short: "The canI CLI.",
	Long:  `For simulating interactions with the canI service`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	action.InitCmd(RootCmd)
}
