package set

import (
	"fmt"

	"github.com/Zach41/birdlog/application/help"
	"github.com/araddon/dateparse"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

const (
	timeLayout = "2000-01-01T12:00:00"
)

var (
	logDeadline *string
	dlCmdLine   = pflag.NewFlagSet("deadline", pflag.ContinueOnError)
)

func init() {
	help.RegisterHelp(dlCmd)
	dlCmd.Flags().AddFlagSet(dlCmdLine)
}

var dlCmd = &cobra.Command{
	Use:     "deadline [time to set]",
	Aliases: []string{"dl"},
	Long:    "set deadline for lastest log",
	Short:   "set deadline for lastest log",
	Example: "set deadline '2022.07.08 12:00:00'",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		t, err := dateparse.ParseAny(args[0])
		if err != nil {
			fmt.Printf("invalid time format: %v\n", err)
			return
		}
		// TODO
		fmt.Println(t)
	},
}
