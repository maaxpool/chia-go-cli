package full_node

import (
	"github.com/LuttyYang/chia-go-cli/logic/rpc/common"
	"reflect"
)

var RpcApis = []*common.TemplateRpcMethod{
	{
		MethodName:   "GetBlockchainState",
		Desc:         "Get the blockchain state",
		Method:       "POST",
		Url:          "get_blockchain_state",
		JsonTemplate: `{}`,
	},
	{
		MethodName:   "GetBlock",
		Desc:         "Gets a full block by header hash",
		Method:       "POST",
		Url:          "get_block",
		JsonTemplate: `{"header_hash": ""}`,
		ValInfo: []*common.TemplateValue{
			{
				Name:    "header-hash",
				Desc:    "Header hash",
				Type:    reflect.String,
				Default: "",
				Path:    "header_hash",
			},
		},
	},
	{
		MethodName:   "GetBlocks",
		Desc:         "Gets a list of full blocks",
		Method:       "POST",
		Url:          "get_blocks",
		JsonTemplate: `{"start": 0, "end": 9999, "exclude_header_hash": false}`,
		ValInfo: []*common.TemplateValue{
			{
				Name: "start",
				Desc: "The start height",
				Type: reflect.Int,
				Path: "start",
			},
			{
				Name: "end",
				Desc: "The end height (non-inclusive)",
				Type: reflect.Int,
				Path: "end",
			},
			{
				Name:    "exclude-header-hash",
				Desc:    "Whether to exclude the header hash in the response (default false)",
				Type:    reflect.Bool,
				Default: false,
				Path:    "exclude_header_hash",
			},
		},
	},
}
