package node

import (
	"chia-go-cli/logic/node"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func newListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List Node",

		RunE: func(cmd *cobra.Command, args []string) error {
			manage, err := node.NewManage(storePath, password)
			if err != nil {
				return err
			}

			list, err := manage.GetNodeList()
			if err != nil {
				return err
			}

			table := tablewriter.NewWriter(os.Stdout)
			table.SetBorder(false)
			table.SetHeader([]string{"Name", "NodeUrl", "Stored Certs"})

			for _, n := range list {
				table.Append([]string{n.Name(), n.NodeUrl(), strings.Join(n.StoredCerts(), ", ")})
			}

			table.Render()

			return nil
		},
	}

	return cmd
}
