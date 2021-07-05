package node

import (
	"chia-go-cli/logic/node"
	"github.com/spf13/cobra"
)

func newRemoveCommand() *cobra.Command {
	var (
		nodeName string
	)

	cmd := &cobra.Command{
		Use:   "remove",
		Short: "remove a node",

		RunE: func(cmd *cobra.Command, args []string) error {
			manage, err := node.NewManage(storePath, password)
			if err != nil {
				return err
			}

			return manage.RemoveNode(nodeName)
		},
	}

	flag := cmd.Flags()
	flag.StringVar(&nodeName, "name", "", "node name")

	_ = cmd.MarkFlagRequired("name")

	return cmd
}
