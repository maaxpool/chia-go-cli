// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.
package wallet

import (
	"chia-go-cli/sdk/client"
	"encoding/json"
)

// GetWalletBalance Get the wallet balance
func GetWalletBalance(walletId int) client.RpcMethod {
	req := &client.TemplateRpcMethod{}
	_ = json.Unmarshal([]byte("{\"MethodName\":\"GetWalletBalance\",\"Desc\":\"Get the wallet balance\",\"Method\":\"POST\",\"Url\":\"get_wallet_balance\",\"JsonTemplate\":\"{\\\"wallet_id\\\": 1}\",\"ValInfo\":[{\"Name\":\"wallet-id\",\"Desc\":\"Wallet Id\",\"Type\":2,\"Default\":1,\"Path\":\"wallet_id\",\"Data\":null}]}"), req)
	req.ValInfo[0].Data = &walletId
	return req
}

// GetTransactions Gets transaction records
func GetTransactions(walletId int, start int, end int) client.RpcMethod {
	req := &client.TemplateRpcMethod{}
	_ = json.Unmarshal([]byte("{\"MethodName\":\"GetTransactions\",\"Desc\":\"Gets transaction records\",\"Method\":\"POST\",\"Url\":\"get_transactions\",\"JsonTemplate\":\"{\\\"wallet_id\\\": 1, start: 0, end: 50}\",\"ValInfo\":[{\"Name\":\"wallet-id\",\"Desc\":\"Wallet Id\",\"Type\":2,\"Default\":1,\"Path\":\"wallet_id\",\"Data\":null},{\"Name\":\"start\",\"Desc\":\"start\",\"Type\":2,\"Default\":0,\"Path\":\"start\",\"Data\":null},{\"Name\":\"end\",\"Desc\":\"end\",\"Type\":2,\"Default\":50,\"Path\":\"end\",\"Data\":null}]}"), req)
	req.ValInfo[0].Data = &walletId
	req.ValInfo[1].Data = &start
	req.ValInfo[2].Data = &end
	return req
}
