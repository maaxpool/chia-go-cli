package cmd

import (
	"fmt"
	"github.com/LuttyYang/chia-go-cli/cmd/node"
	"github.com/LuttyYang/chia-go-cli/cmd/rpc"
	"github.com/spf13/cobra"
	"io"
	"os"
)

const Version = "0.0.1"

var Build = "local_build"

var RootCmd *cobra.Command

func init() {
	RootCmd = &cobra.Command{
		Use:              "chia-go-cli",
		Short:            "Chia Cli, Golang version",
		SilenceUsage:     true,
		SilenceErrors:    true,
		TraverseChildren: true,
		Args:             noArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return ShowHelp(os.Stderr)(cmd, args)
		},
		Version:               fmt.Sprintf("%s, build %s", Version, Build),
		DisableFlagsInUseLine: true,
	}

	RootCmd.AddCommand(rpc.NewRootCommand())
	RootCmd.AddCommand(node.NewRootCommand())
}

// ShowHelp shows the command help.
func ShowHelp(err io.Writer) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		cmd.SetErr(err)
		cmd.HelpFunc()(cmd, args)
		return nil
	}
}

func noArgs(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		return nil
	}
	return fmt.Errorf("chia-go-cli: '%s' is not a vaild command.\nSee 'chia-go-cli --help'", args[0])
}
