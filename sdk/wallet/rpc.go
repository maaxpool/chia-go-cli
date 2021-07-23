package wallet

import (
	"chia-go-cli/sdk/client"
	"reflect"
)

var RpcApis = []*client.TemplateRpcMethod{
	{
		MethodName:   "GetWalletBalance",
		Desc:         "Get the wallet balance",
		Method:       "POST",
		Url:          "get_wallet_balance",
		JsonTemplate: `{"wallet_id": 1}`,
		ValInfo: []*client.TemplateValue{
			{
				Name:    "wallet-id",
				Desc:    "Wallet Id",
				Type:    reflect.Int,
				Default: 1,
				Path:    "wallet_id",
			},
		},
	},
}
