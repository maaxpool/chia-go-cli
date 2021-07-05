package wallet

import (
	"chia-go-cli/logic/rpc/common"
	"reflect"
)

var RpcApis = []*common.TemplateRpcMethod{
	{
		MethodName:   "GetWalletBalance",
		Desc:         "Get the wallet balance",
		Method:       "POST",
		Url:          "get_wallet_balance",
		JsonTemplate: `{"wallet_id": 1}`,
		ValInfo: []*common.TemplateValue{
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
