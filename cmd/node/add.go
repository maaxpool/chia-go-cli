package node

import (
	"chia-go-cli/logic/node"
	"github.com/spf13/cobra"
)

func newAddCommand() *cobra.Command {
	var (
		nodeName    string
		certType    string
		certCrtPath string
		certKeyPath string
	)

	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add cert to node",

		RunE: func(cmd *cobra.Command, args []string) error {
			manage, err := node.NewManage(storePath, password)
			if err != nil {
				return err
			}

			cType, err := node.GetCertTypeFromString(certType)
			if err != nil {
				return err
			}

			n, err := manage.GetNode(nodeName)
			if err != nil {
				return err
			}

			return n.AddCert(cType, certCrtPath, certKeyPath)
		},
	}

	flag := cmd.Flags()
	flag.StringVar(&nodeName, "name", "", "node name")
	flag.StringVar(&certType, "cert-type", "", "cert type")
	flag.StringVar(&certCrtPath, "crt-path", "", "cert crt path")
	flag.StringVar(&certKeyPath, "key-path", "", "cert key path")

	_ = cmd.MarkFlagRequired("name")
	_ = cmd.MarkFlagRequired("cert-type")
	_ = cmd.MarkFlagRequired("crt-path")
	_ = cmd.MarkFlagRequired("key-path")

	return cmd
}
