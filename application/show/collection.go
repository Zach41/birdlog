package show

import (
	"fmt"

	"github.com/spf13/cobra"
)

var collection = &cobra.Command{
	Use:     "collection",
	Aliases: []string{"coll"},
	Short:   "describe collection meta datas",
	Long:    "describe collection meta datas",
	Run: func(cmd *cobra.Command, args []string) {
		//TODO
		fmt.Printf("show collection %s\n", args[0])
	},
}
