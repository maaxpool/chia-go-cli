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

// GetInitialFreezePeriod Retrieves the initial freeze period for the blockchain (no transactions allowed).
func GetInitialFreezePeriod() client.RpcMethod {
	req := &client.TemplateRpcMethod{}
	_ = json.Unmarshal([]byte("{\"MethodName\":\"GetInitialFreezePeriod\",\"Desc\":\"Retrieves the initial freeze period for the blockchain (no transactions allowed).\",\"Method\":\"POST\",\"Url\":\"get_initial_freeze_period\",\"JsonTemplate\":\"{}\",\"ValInfo\":null}"), req)
	return req
}

// GetTransaction Gets a transaction record by transaction id
func GetTransaction(transactionId string) client.RpcMethod {
	req := &client.TemplateRpcMethod{}
	_ = json.Unmarshal([]byte("{\"MethodName\":\"GetTransaction\",\"Desc\":\"Gets a transaction record by transaction id\",\"Method\":\"POST\",\"Url\":\"get_transaction\",\"JsonTemplate\":\"{\\\"transaction_id\\\": \\\"\\\"}\",\"ValInfo\":[{\"Name\":\"transaction-id\",\"Desc\":\"Transaction id\",\"Type\":24,\"Default\":null,\"Path\":\"transaction_id\",\"Data\":null}]}"), req)
	req.ValInfo[0].Data = &transactionId
	return req
}

// GetTransactions Gets transaction records
func GetTransactions(walletId int, start int, end int) client.RpcMethod {
	req := &client.TemplateRpcMethod{}
	_ = json.Unmarshal([]byte("{\"MethodName\":\"GetTransactions\",\"Desc\":\"Gets transaction records\",\"Method\":\"POST\",\"Url\":\"get_transactions\",\"JsonTemplate\":\"{\\\"wallet_id\\\": 1, \\\"start\\\": 0, \\\"end\\\": 50}\",\"ValInfo\":[{\"Name\":\"wallet-id\",\"Desc\":\"Wallet Id\",\"Type\":2,\"Default\":1,\"Path\":\"wallet_id\",\"Data\":null},{\"Name\":\"start\",\"Desc\":\"start\",\"Type\":2,\"Default\":0,\"Path\":\"start\",\"Data\":null},{\"Name\":\"end\",\"Desc\":\"end\",\"Type\":2,\"Default\":50,\"Path\":\"end\",\"Data\":null}]}"), req)
	req.ValInfo[0].Data = &walletId
	req.ValInfo[1].Data = &start
	req.ValInfo[2].Data = &end
	return req
}

// GetNextAddress Gets a new (or not new) address
func GetNextAddress(walletId int, newAddress bool) client.RpcMethod {
	req := &client.TemplateRpcMethod{}
	_ = json.Unmarshal([]byte("{\"MethodName\":\"GetNextAddress\",\"Desc\":\"Gets a new (or not new) address\",\"Method\":\"POST\",\"Url\":\"get_next_address\",\"JsonTemplate\":\"{\\\"new_address\\\": true, \\\"wallet_id\\\": 1}\",\"ValInfo\":[{\"Name\":\"wallet-id\",\"Desc\":\"Wallet Id\",\"Type\":2,\"Default\":1,\"Path\":\"wallet_id\",\"Data\":null},{\"Name\":\"new-address\",\"Desc\":\"New address\",\"Type\":1,\"Default\":1,\"Path\":\"new_address\",\"Data\":null}]}"), req)
	req.ValInfo[0].Data = &walletId
	req.ValInfo[1].Data = &newAddress
	return req
}

// SendTransaction Sends a standard transaction to a target puzzle_hash.
func SendTransaction(walletId int, address string) client.RpcMethod {
	req := &client.TemplateRpcMethod{}
	_ = json.Unmarshal([]byte("{\"MethodName\":\"SendTransaction\",\"Desc\":\"Sends a standard transaction to a target puzzle_hash.\",\"Method\":\"POST\",\"Url\":\"send_transaction\",\"JsonTemplate\":\"{\\\"wallet_id\\\": 1,\\\"amount\\\": 0,\\\"fee\\\": 0,\\\"address\\\":\\\"\\\"}\",\"ValInfo\":[{\"Name\":\"wallet-id\",\"Desc\":\"Wallet Id\",\"Type\":2,\"Default\":1,\"Path\":\"wallet_id\",\"Data\":null},{\"Name\":\"address\",\"Desc\":\"Address\",\"Type\":24,\"Default\":1,\"Path\":\"address\",\"Data\":null}]}"), req)
	req.ValInfo[0].Data = &walletId
	req.ValInfo[1].Data = &address
	return req
}

// GetTransactionCount Gets the number of transactions in this wallet.
func GetTransactionCount(walletId int) client.RpcMethod {
	req := &client.TemplateRpcMethod{}
	_ = json.Unmarshal([]byte("{\"MethodName\":\"GetTransactionCount\",\"Desc\":\"Gets the number of transactions in this wallet.\",\"Method\":\"POST\",\"Url\":\"get_transaction_count\",\"JsonTemplate\":\"{\\\"wallet_id\\\": 1}\",\"ValInfo\":[{\"Name\":\"wallet-id\",\"Desc\":\"Wallet Id\",\"Type\":2,\"Default\":1,\"Path\":\"wallet_id\",\"Data\":null}]}"), req)
	req.ValInfo[0].Data = &walletId
	return req
}
