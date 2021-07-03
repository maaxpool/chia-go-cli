package rpc

import (
	"github.com/LuttyYang/chia-go-cli/logic/rpc/common"
	"github.com/spf13/cobra"
	"log"
	"reflect"
)

type buildApiCallback func(api *common.TemplateRpcMethod) func(cmd *cobra.Command, args []string) error

var (
	storePath string
	password  string
	nodeName  string
)

func NewRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rpc",
		Short: "Call Rpc Api",
	}

	flag := cmd.PersistentFlags()
	flag.StringVar(&password, "password", "", "store password")
	flag.StringVar(&storePath, "store-path", ".chia-go-cli-store", "store path")
	flag.StringVar(&nodeName, "node", "", "node name")

	_ = cmd.MarkPersistentFlagRequired("node")

	cmd.AddCommand(newWalletCommand())
	cmd.AddCommand(newFullNodeCommand())
	cmd.AddCommand(newSharedCommand())

	return cmd
}

func buildApi(cmd *cobra.Command, apis []*common.TemplateRpcMethod, call buildApiCallback) {
	for _, api := range apis {
		apiCommand := &cobra.Command{
			Use:   api.MethodName,
			Short: api.Desc,
			RunE:  call(api),
		}

		flag := apiCommand.Flags()
		for _, value := range api.ValInfo {
			switch value.Type {
			case reflect.Int:
				if value.Default == nil {
					value.Data = flag.Int(value.Name, 0, value.Desc)
					_ = apiCommand.MarkFlagRequired(value.Name)
				} else {
					value.Data = flag.Int(value.Name, value.Default.(int), value.Desc)
				}
			case reflect.String:
				if value.Default == nil {
					value.Data = flag.String(value.Name, "", value.Desc)
					_ = apiCommand.MarkFlagRequired(value.Name)
				} else {
					value.Data = flag.String(value.Name, value.Default.(string), value.Desc)
				}
			case reflect.Bool:
				if value.Default == nil {
					value.Data = flag.Bool(value.Name, false, value.Desc)
					_ = apiCommand.MarkFlagRequired(value.Name)
				} else {
					value.Data = flag.Bool(value.Name, value.Default.(bool), value.Desc)
				}
			default:
				log.Panicf("build api command fail, not support type: %v", value.Type)
			}
		}

		cmd.AddCommand(apiCommand)
	}
}
