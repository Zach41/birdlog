package application

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type CommandInfo struct {
	*cobra.Command
	CommandLine *pflag.FlagSet
}
