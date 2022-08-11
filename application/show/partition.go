package show

import (
	"fmt"

	"github.com/spf13/cobra"
)

var partition = &cobra.Command{
	Use:     "partition",
	Aliases: []string{"part"},
	Short:   "describe partition meta datas",
	Long:    "describe partition meta datas",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO
		fmt.Printf("show partition %s\n", args[0])
	},
}
