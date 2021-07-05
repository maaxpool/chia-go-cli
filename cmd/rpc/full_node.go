package rpc

import (
	"bytes"
	"chia-go-cli/logic/node"
	"chia-go-cli/logic/rpc/common"
	"chia-go-cli/logic/rpc/full_node"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io"
)

var fullNodePort uint16

func newFullNodeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "full-node",
		Short: "Full Node RPC API",
	}

	buildApi(cmd, full_node.RpcApis, buildFullNodeApiCall)

	flag := cmd.PersistentFlags()
	flag.Uint16Var(&fullNodePort, "port", 8555, "call port")

	return cmd
}

func buildFullNodeApiCall(api *common.TemplateRpcMethod) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		manage, err := node.NewManage(storePath, password)
		if err != nil {
			return err
		}

		getNode, err := manage.GetNode(nodeName)
		if err != nil {
			return err
		}

		client, err := common.NewClient(getNode, fullNodePort, node.CertTypePrivateFullNode)
		if err != nil {
			return err
		}

		resp, err := client.RawRequest(api)
		if err != nil {
			return err
		}

		if resp.StatusCode != 200 {
			return fmt.Errorf("call api %s got status %d", api.MethodName, resp.StatusCode)
		}

		respData, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		var out bytes.Buffer
		err = json.Indent(&out, respData, "", "  ")
		if err == nil {
			respData = out.Bytes()
		}

		fmt.Println(string(respData))
		return nil
	}
}
