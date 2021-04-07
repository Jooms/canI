package action

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

// helloCmd represents the hello command
var actionCmd = &cobra.Command{
	Use:   "action",
	Short: "something to do with actions",
	Run: func(cmd *cobra.Command, args []string) {
		glog.V(2).Infof("canI action called")
		fmt.Println("canI action called")
	},
}

func InitCmd(RootCmd *cobra.Command) {
	RootCmd.AddCommand(actionCmd)
}
