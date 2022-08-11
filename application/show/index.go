package show

import (
	"fmt"

	"github.com/spf13/cobra"
)

var index = &cobra.Command{
	Use:     "index",
	Aliases: []string{"idx"},
	Short:   "describe index meta datas",
	Long:    "describe index meta datas",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO
		fmt.Printf("show index %s\n", args[0])
	},
}
