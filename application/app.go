package application

import (
	"fmt"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	cmdset = map[string]*CommandInfo{}
)

type Analyze struct {
}

func (app *Analyze) execute(input string) {
	args, err := parseArgs(input)
	if err != nil {
		fmt.Printf("error input: %v\n", err)
		return
	}
	if isExit(args[0]) {
		fmt.Println("byte")
		os.Exit(0)
	}
	cmd, ok := cmdset[args[0]]
	if !ok {
		fmt.Printf("unknown command \"%s\"\nRun \"help\" for usage\n", args[0])
		return
	}
	var execargs []string
	if len(args) > 1 {
		execargs = args[1:]
	}
	cmd.SetArgs(execargs)
	cmd.Execute()
	cmd.SetArgs(execargs)
}

func (app *Analyze) getSuggestions(prompt.Document) []prompt.Suggest {
	// TODO
	return nil
}

func NewAnalyzier() *prompt.Prompt {
	app := &Analyze{}
	return prompt.New(app.execute, app.getSuggestions)
}

func RegisterCommands(cmd *cobra.Command, flags *pflag.FlagSet) error {
	keys := []string{cmd.Use}
	keys = append(keys, cmd.Aliases...)
	info := &CommandInfo{
		CommandLine: flags,
		Command:     cmd,
	}
	for _, key := range keys {
		if _, ok := cmdset[key]; ok {
			return fmt.Errorf("command %s already registered", key)
		}
		cmdset[key] = info
	}
	return nil
}

func isExit(input string) bool {
	return strings.ToLower(strings.TrimSpace(input)) == "exit"
}

// func parseInput(input string) []string {
// 	return nil
// }

// func parseToken
