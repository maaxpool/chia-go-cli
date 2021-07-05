package node

import (
	"chia-go-cli/logic/node"
	"github.com/spf13/cobra"
)

func newCreateCommand() *cobra.Command {
	var (
		nodeName string
		nodeUrl  string
	)

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a node",

		RunE: func(cmd *cobra.Command, args []string) error {
			manage, err := node.NewManage(storePath, password)
			if err != nil {
				return err
			}

			_, err = manage.GetOrCreateNode(nodeName, nodeUrl)
			return nil
		},
	}

	flag := cmd.Flags()
	flag.StringVar(&nodeName, "name", "", "node name")
	flag.StringVar(&nodeUrl, "url", "", "node url")

	_ = cmd.MarkFlagRequired("name")

	return cmd
}
