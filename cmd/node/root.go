package node

import "github.com/spf13/cobra"

var (
	storePath string
	password  string
)

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "node",
		Short: "Node Manage",
	}

	cmd.AddCommand(newAddCommand())
	cmd.AddCommand(newListCommand())
	cmd.AddCommand(newRemoveCommand())
	cmd.AddCommand(newCreateCommand())

	flag := cmd.PersistentFlags()
	flag.StringVar(&password, "password", "", "store password")
	flag.StringVar(&storePath, "store-path", ".chia-go-cli-store", "store path")

	return cmd
}
