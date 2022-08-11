package show

import (
	"log"

	"github.com/Zach41/birdlog/application"
	"github.com/Zach41/birdlog/application/help"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func init() {
	if err := help.RegisterHelp(show); err != nil {
		log.Fatal(err)
	}
	if err := application.RegisterCommands(show, cmdLine); err != nil {
		log.Fatal(err)
	}
	show.Flags().AddFlagSet(cmdLine)
	show.AddCommand(collection, index, segment, partition)
}

// flagsets for show command
var cmdLine = pflag.NewFlagSet("show", pflag.ContinueOnError)

var show = &cobra.Command{
	Use:               "show",
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
	Short:             "describes meta datas of collections, indexes, partitions or segments",
	Long:              "describes meta datas of collections, indexes, partitions or segments",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// func Command() *cobra.Command {
// 	show.AddCommand(collection, index, segment, partition)
// 	return show
// }
