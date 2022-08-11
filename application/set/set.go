package set

import (
	"log"

	"github.com/Zach41/birdlog/application"
	"github.com/Zach41/birdlog/application/help"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func init() {
	if err := help.RegisterHelp(setCmd); err != nil {
		log.Fatal(err)
	}
	if err := application.RegisterCommands(setCmd, cmdLine); err != nil {
		log.Fatal(err)
	}
	setCmd.Flags().AddFlagSet(cmdLine)
	setCmd.AddCommand(dlCmd)
}

// flagset for set command
var cmdLine = pflag.NewFlagSet("set", pflag.ContinueOnError)

var setCmd = &cobra.Command{
	Use:               "set",
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	Short:             "set global variables for log profile",
	Long:              "set global variables for log profile",
}
