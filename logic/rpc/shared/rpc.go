package shared

import (
	"github.com/LuttyYang/chia-go-cli/logic/rpc/common"
	"reflect"
)

var RpcApis = []*common.TemplateRpcMethod{
	{
		MethodName:   "GetConnections",
		Desc:         "Returns a list of peers that we are currently connected to",
		Method:       "POST",
		Url:          "get_connections",
		JsonTemplate: `{}`,
	},
	{
		MethodName:   "OpenConnection",
		Desc:         "Opens a connection to another peer",
		Method:       "POST",
		Url:          "open_connection",
		JsonTemplate: `{"host": "", "port": 0}`,
		ValInfo: []*common.TemplateValue{
			{
				Name:    "host",
				Desc:    "ip or dns name of the peer",
				Type:    reflect.String,
				Default: "",
				Path:    "host",
			},
			{
				Name:    "port",
				Desc:    "port of the peer",
				Type:    reflect.Int,
				Default: 0,
				Path:    "port",
			},
		},
	},
	{
		MethodName:   "CloseConnection",
		Desc:         "Closes a connection with a peer",
		Method:       "POST",
		Url:          "close_connection",
		JsonTemplate: `{"node_id": ""}`,
		ValInfo: []*common.TemplateValue{
			{
				Name:    "node-id",
				Desc:    "node id of the peer",
				Type:    reflect.String,
				Default: "",
				Path:    "node_id",
			},
		},
	},
	{
		MethodName:   "StopNode",
		Desc:         "Stops the node",
		Method:       "POST",
		Url:          "stop_node",
		JsonTemplate: `{}`,
	},
}
