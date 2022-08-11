package show

import (
	"fmt"

	"github.com/spf13/cobra"
)

var segment = &cobra.Command{
	Use:     "segment",
	Aliases: []string{"seg"},
	Short:   "describe segment meta datas",
	Long:    "describe segment meta datas",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO
		fmt.Printf("show segmengt %s\n", args[0])
	},
}
