package help

import (
	"fmt"
	"log"
	"strings"

	"github.com/Zach41/birdlog/application"
	"github.com/spf13/cobra"
)

func init() {
	if err := application.RegisterCommands(helpCmd, nil); err != nil {
		log.Fatal(err)
	}
}

var (
	allCmds = map[string]*cobra.Command{
		"exit": {
			Use:  "exit",
			Long: "exit the program",
		},
	}
)

func RegisterHelp(cmd *cobra.Command) error {
	if _, ok := allCmds[cmd.Use]; ok {
		return fmt.Errorf("command `%s` already registered", cmd.Use)
	}
	allCmds[cmd.Use] = cmd
	return nil
}

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "print help message",
	Long:  "print help message",
	Run: func(cmd *cobra.Command, args []string) {

		msgBuilder := strings.Builder{}
		if len(args) == 0 {
			msgBuilder.WriteString("COMMANDS:\n")
			cmdfmt := "    %-16s %s\n"
			for _, cmd := range allCmds {
				msgBuilder.WriteString(fmt.Sprintf(cmdfmt, cmd.Use, cmd.Long))
			}
			msgBuilder.WriteString("Run `help [command]` to get command usage")
			fmt.Println(msgBuilder.String())
		} else if len(args) == 1 {
			cmd, ok := allCmds[args[0]]
			if !ok {
				fmt.Printf("Unknown command `%s`, run `help` for usage\n", args[0])
				return
			}
			cmd.SetArgs([]string{"--help"})
			cmd.Execute()
			cmd.SetArgs(nil)
		} else {
			fmt.Println("Invalid args for `help` command, run `help` for usage")
		}
	},
}
