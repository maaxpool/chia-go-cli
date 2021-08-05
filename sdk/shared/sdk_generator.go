// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.
package shared

import (
	"chia-go-cli/sdk/client"
	"encoding/json"
)

// GetConnections Returns a list of peers that we are currently connected to
func GetConnections() client.RpcMethod {
	req := &client.TemplateRpcMethod{}
	_ = json.Unmarshal([]byte("{\"MethodName\":\"GetConnections\",\"Desc\":\"Returns a list of peers that we are currently connected to\",\"Method\":\"POST\",\"Url\":\"get_connections\",\"JsonTemplate\":\"{}\",\"ValInfo\":null}"), req)
	return req
}

// OpenConnection Opens a connection to another peer
func OpenConnection(host string, port int) client.RpcMethod {
	req := &client.TemplateRpcMethod{}
	_ = json.Unmarshal([]byte("{\"MethodName\":\"OpenConnection\",\"Desc\":\"Opens a connection to another peer\",\"Method\":\"POST\",\"Url\":\"open_connection\",\"JsonTemplate\":\"{\\\"host\\\": \\\"\\\", \\\"port\\\": 0}\",\"ValInfo\":[{\"Name\":\"host\",\"Desc\":\"ip or dns name of the peer\",\"Type\":24,\"Default\":\"\",\"Path\":\"host\",\"Data\":null},{\"Name\":\"port\",\"Desc\":\"port of the peer\",\"Type\":2,\"Default\":0,\"Path\":\"port\",\"Data\":null}]}"), req)
	req.ValInfo[0].Data = &host
	req.ValInfo[1].Data = &port
	return req
}

// CloseConnection Closes a connection with a peer
func CloseConnection(nodeId string) client.RpcMethod {
	req := &client.TemplateRpcMethod{}
	_ = json.Unmarshal([]byte("{\"MethodName\":\"CloseConnection\",\"Desc\":\"Closes a connection with a peer\",\"Method\":\"POST\",\"Url\":\"close_connection\",\"JsonTemplate\":\"{\\\"node_id\\\": \\\"\\\"}\",\"ValInfo\":[{\"Name\":\"node-id\",\"Desc\":\"node id of the peer\",\"Type\":24,\"Default\":\"\",\"Path\":\"node_id\",\"Data\":null}]}"), req)
	req.ValInfo[0].Data = &nodeId
	return req
}

// StopNode Stops the node
func StopNode() client.RpcMethod {
	req := &client.TemplateRpcMethod{}
	_ = json.Unmarshal([]byte("{\"MethodName\":\"StopNode\",\"Desc\":\"Stops the node\",\"Method\":\"POST\",\"Url\":\"stop_node\",\"JsonTemplate\":\"{}\",\"ValInfo\":null}"), req)
	return req
}
