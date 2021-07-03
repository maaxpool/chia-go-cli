package full_node

import (
	"github.com/LuttyYang/chia-go-cli/logic/rpc/common"
)

var RpcApis = []*common.TemplateRpcMethod{
	{
		MethodName:   "GetBlockchainState",
		Desc:         "Get the blockchain state",
		Method:       "POST",
		Url:          "get_blockchain_state",
		JsonTemplate: `{}`,
	},
}
