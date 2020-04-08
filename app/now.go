package app

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var Format bool
var Timestamp bool

func NewNowCommand() *cobra.Command {
	nowCmd := &cobra.Command{
		Use:   "now",
		Short: "print current time",
		Long:  "print current time",
		Run: func(cmd *cobra.Command, args []string) {
			if Timestamp {
				fmt.Println(time.Now().UnixNano() / 1000000)
				return
			}
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		},
	}
	nowCmd.Flags().BoolVarP(&Format, "format", "f", true, "print current time with YYYY-MM-DD HH:mm:ss")
	nowCmd.Flags().BoolVarP(&Timestamp, "timestamp", "t", false, "print current time with timestamp")
	return nowCmd
}
