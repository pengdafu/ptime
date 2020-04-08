package app

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var timestamp bool

func NewPTimeCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "ptime",
		Short: "time convert tool",
		Long: `The time convert tool.
You can use it to print the current time, or you can use it to convert the time stamp into a readable time string.
It has many functions, you can use -- help to view.`,
		Run: func(cmd *cobra.Command, args []string) {
			if timestamp {
				fmt.Println(time.Now().UnixNano() / 1000000)
				return
			}
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		},
	}
	rootCmd.Flags().BoolVarP(&timestamp, "timestamp", "t", false, "print current time with timestamp")
	rootCmd.AddCommand(NewNowCommand())
	rootCmd.AddCommand(NewConvertCommand())
	return rootCmd
}
