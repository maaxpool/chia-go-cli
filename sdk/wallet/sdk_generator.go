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
	req.ValInfo[0].Data = walletId
	return req
}

// GetTransactions Gets transaction records
func GetTransactions(walletId int, offset int) client.RpcMethod {
	req := &client.TemplateRpcMethod{}
	_ = json.Unmarshal([]byte("{\"MethodName\":\"GetTransactions\",\"Desc\":\"Gets transaction records\",\"Method\":\"POST\",\"Url\":\"get_transactions\",\"JsonTemplate\":\"{\\\"wallet_id\\\": 1, \\\"offset\\\": 0}\",\"ValInfo\":[{\"Name\":\"wallet-id\",\"Desc\":\"Wallet Id\",\"Type\":2,\"Default\":1,\"Path\":\"wallet_id\",\"Data\":null},{\"Name\":\"offset\",\"Desc\":\"Offset\",\"Type\":2,\"Default\":1,\"Path\":\"offset\",\"Data\":null}]}"), req)
	req.ValInfo[0].Data = walletId
	req.ValInfo[1].Data = offset
	return req
}
