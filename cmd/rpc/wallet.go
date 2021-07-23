package rpc

import (
	"bytes"
	"chia-go-cli/logic/node"
	"chia-go-cli/logic/rpc/common"
	"chia-go-cli/sdk/client"
	"chia-go-cli/sdk/wallet"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"io"
)

var walletPort uint16

func newWalletCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "wallet",
		Short: "Wallet RPC API",
	}

	buildApi(cmd, wallet.RpcApis, buildWalletApiCall)

	flag := cmd.PersistentFlags()
	flag.Uint16Var(&walletPort, "port", 9256, "call port")

	return cmd
}

func buildWalletApiCall(api *client.TemplateRpcMethod) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		manage, err := node.NewManage(storePath, password)
		if err != nil {
			return err
		}

		getNode, err := manage.GetNode(nodeName)
		if err != nil {
			return err
		}

		client, err := common.NewClient(getNode, walletPort, node.CertTypePrivateWallet)
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
