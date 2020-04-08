package app

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

func NewConvertCommand() *cobra.Command {
	cvtCmd := &cobra.Command{
		Use:   "c",
		Short: "convert time",
		Long:  "Convert timestamp to format string, etc",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if t, err := strconv.Atoi(args[0]); err != nil {
				fmt.Println("time no valid")
			} else {
				fmt.Println(time.Unix(int64(t)/1000, 0).Format("2006-01-02 15:04:05"))
			}
		},
	}
	return cvtCmd
}
