// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.
package full_node

import (
	"chia-go-cli/sdk/client"
	"encoding/json"
)

// GetBlockchainState Get the blockchain state
func GetBlockchainState() client.RpcMethod {
	req := &client.TemplateRpcMethod{}
	_ = json.Unmarshal([]byte("{\"MethodName\":\"GetBlockchainState\",\"Desc\":\"Get the blockchain state\",\"Method\":\"POST\",\"Url\":\"get_blockchain_state\",\"JsonTemplate\":\"{}\",\"ValInfo\":null}"), req)
	return req
}

// GetBlock Gets a full block by header hash
func GetBlock(headerHash string) client.RpcMethod {
	req := &client.TemplateRpcMethod{}
	_ = json.Unmarshal([]byte("{\"MethodName\":\"GetBlock\",\"Desc\":\"Gets a full block by header hash\",\"Method\":\"POST\",\"Url\":\"get_block\",\"JsonTemplate\":\"{\\\"header_hash\\\": \\\"\\\"}\",\"ValInfo\":[{\"Name\":\"header-hash\",\"Desc\":\"Header hash\",\"Type\":24,\"Default\":null,\"Path\":\"header_hash\",\"Data\":null}]}"), req)
	req.ValInfo[0].Data = &headerHash
	return req
}

// GetBlocks Gets a list of full blocks
func GetBlocks(start int, end int, excludeHeaderHash bool) client.RpcMethod {
	req := &client.TemplateRpcMethod{}
	_ = json.Unmarshal([]byte("{\"MethodName\":\"GetBlocks\",\"Desc\":\"Gets a list of full blocks\",\"Method\":\"POST\",\"Url\":\"get_blocks\",\"JsonTemplate\":\"{\\\"start\\\": 0, \\\"end\\\": 9999, \\\"exclude_header_hash\\\": false}\",\"ValInfo\":[{\"Name\":\"start\",\"Desc\":\"The start height\",\"Type\":2,\"Default\":null,\"Path\":\"start\",\"Data\":null},{\"Name\":\"end\",\"Desc\":\"The end height (non-inclusive)\",\"Type\":2,\"Default\":null,\"Path\":\"end\",\"Data\":null},{\"Name\":\"exclude-header-hash\",\"Desc\":\"Whether to exclude the header hash in the response (default false)\",\"Type\":1,\"Default\":false,\"Path\":\"exclude_header_hash\",\"Data\":null}]}"), req)
	req.ValInfo[0].Data = &start
	req.ValInfo[1].Data = &end
	req.ValInfo[2].Data = &excludeHeaderHash
	return req
}

// GetBlockRecordByHeight Retrieves a block record by height (assuming the height <= peak height)
func GetBlockRecordByHeight(height int) client.RpcMethod {
	req := &client.TemplateRpcMethod{}
	_ = json.Unmarshal([]byte("{\"MethodName\":\"GetBlockRecordByHeight\",\"Desc\":\"Retrieves a block record by height (assuming the height \\u003c= peak height)\",\"Method\":\"POST\",\"Url\":\"get_block_record_by_height\",\"JsonTemplate\":\"{\\\"height\\\": 0}\",\"ValInfo\":[{\"Name\":\"height\",\"Desc\":\"The height to get\",\"Type\":2,\"Default\":null,\"Path\":\"height\",\"Data\":null}]}"), req)
	req.ValInfo[0].Data = &height
	return req
}

// GetBlockRecord Retrieves a block record by header hash
func GetBlockRecord(headerHash string) client.RpcMethod {
	req := &client.TemplateRpcMethod{}
	_ = json.Unmarshal([]byte("{\"MethodName\":\"GetBlockRecord\",\"Desc\":\"Retrieves a block record by header hash\",\"Method\":\"POST\",\"Url\":\"get_block_record\",\"JsonTemplate\":\"{\\\"header_hash\\\": \\\"\\\"}\",\"ValInfo\":[{\"Name\":\"header-hash\",\"Desc\":\"The block's header_hash\",\"Type\":24,\"Default\":null,\"Path\":\"header_hash\",\"Data\":null}]}"), req)
	req.ValInfo[0].Data = &headerHash
	return req
}

// GetBlockRecords Retrieves block records in a range
func GetBlockRecords(start int, end int) client.RpcMethod {
	req := &client.TemplateRpcMethod{}
	_ = json.Unmarshal([]byte("{\"MethodName\":\"GetBlockRecords\",\"Desc\":\"Retrieves block records in a range\",\"Method\":\"POST\",\"Url\":\"get_block_records\",\"JsonTemplate\":\"{\\\"start\\\": 0, \\\"end\\\": 9999}\",\"ValInfo\":[{\"Name\":\"start\",\"Desc\":\"The start height\",\"Type\":2,\"Default\":null,\"Path\":\"start\",\"Data\":null},{\"Name\":\"end\",\"Desc\":\"The end height (non-inclusive)\",\"Type\":2,\"Default\":null,\"Path\":\"end\",\"Data\":null}]}"), req)
	req.ValInfo[0].Data = &start
	req.ValInfo[1].Data = &end
	return req
}

// GetUnfinishedBlockHeaders Retrieves recent unfinished header blocks
func GetUnfinishedBlockHeaders() client.RpcMethod {
	req := &client.TemplateRpcMethod{}
	_ = json.Unmarshal([]byte("{\"MethodName\":\"GetUnfinishedBlockHeaders\",\"Desc\":\"Retrieves recent unfinished header blocks\",\"Method\":\"POST\",\"Url\":\"get_unfinished_block_headers\",\"JsonTemplate\":\"{}\",\"ValInfo\":null}"), req)
	return req
}

// GetNetworkSpace Retrieves an estimate of the total plotted space of all farmers, in bytes
func GetNetworkSpace(olderBlockHeaderHash string, newerBlockHeaderHash string) client.RpcMethod {
	req := &client.TemplateRpcMethod{}
	_ = json.Unmarshal([]byte("{\"MethodName\":\"GetNetworkSpace\",\"Desc\":\"Retrieves an estimate of the total plotted space of all farmers, in bytes\",\"Method\":\"POST\",\"Url\":\"get_network_space\",\"JsonTemplate\":\"{\\\"older_block_header_hash\\\": \\\"\\\", \\\"newer_block_header_hash\\\": \\\"\\\"}\",\"ValInfo\":[{\"Name\":\"older-block-header-hash\",\"Desc\":\"The start header hash\",\"Type\":24,\"Default\":null,\"Path\":\"older_block_header_hash\",\"Data\":null},{\"Name\":\"newer-block-header-hash\",\"Desc\":\"The end header hash\",\"Type\":24,\"Default\":null,\"Path\":\"newer_block_header_hash\",\"Data\":null}]}"), req)
	req.ValInfo[0].Data = &olderBlockHeaderHash
	req.ValInfo[1].Data = &newerBlockHeaderHash
	return req
}

// GetAdditionsAndRemovals Retrieves the additions and removals(state transitions) for a certain block. Returns coin records for each addition and removal
func GetAdditionsAndRemovals(headerHash string) client.RpcMethod {
	req := &client.TemplateRpcMethod{}
	_ = json.Unmarshal([]byte("{\"MethodName\":\"GetAdditionsAndRemovals\",\"Desc\":\"Retrieves the additions and removals(state transitions) for a certain block. Returns coin records for each addition and removal\",\"Method\":\"POST\",\"Url\":\"get_additions_and_removals\",\"JsonTemplate\":\"{\\\"header_hash\\\": \\\"\\\"}\",\"ValInfo\":[{\"Name\":\"header-hash\",\"Desc\":\"Header hash of the block\",\"Type\":24,\"Default\":null,\"Path\":\"header_hash\",\"Data\":null}]}"), req)
	req.ValInfo[0].Data = &headerHash
	return req
}

// GetInitialFreezePeriod Retrieves the initial freeze period for the blockchain (no transactions allowed)
func GetInitialFreezePeriod() client.RpcMethod {
	req := &client.TemplateRpcMethod{}
	_ = json.Unmarshal([]byte("{\"MethodName\":\"GetInitialFreezePeriod\",\"Desc\":\"Retrieves the initial freeze period for the blockchain (no transactions allowed)\",\"Method\":\"POST\",\"Url\":\"get_initial_freeze_period\",\"JsonTemplate\":\"{}\",\"ValInfo\":null}"), req)
	return req
}

// GetNetworkInfo Retrieves some information about the current network
func GetNetworkInfo() client.RpcMethod {
	req := &client.TemplateRpcMethod{}
	_ = json.Unmarshal([]byte("{\"MethodName\":\"GetNetworkInfo\",\"Desc\":\"Retrieves some information about the current network\",\"Method\":\"POST\",\"Url\":\"get_network_info\",\"JsonTemplate\":\"{}\",\"ValInfo\":null}"), req)
	return req
}

// GetCoinRecordsByPuzzleHash Retrieves a list of coin records with a certain puzzle hash
func GetCoinRecordsByPuzzleHash(puzzleHash string, startHeight int, endHeight int, includeSpendCoins bool) client.RpcMethod {
	req := &client.TemplateRpcMethod{}
	_ = json.Unmarshal([]byte("{\"MethodName\":\"GetCoinRecordsByPuzzleHash\",\"Desc\":\"Retrieves a list of coin records with a certain puzzle hash\",\"Method\":\"POST\",\"Url\":\"get_coin_records_by_puzzle_hash\",\"JsonTemplate\":\"{\\\"puzzle_hash\\\": \\\"\\\", \\\"start_height\\\": 0, \\\"end_height\\\": 0, \\\"include_spend_coins\\\": false}\",\"ValInfo\":[{\"Name\":\"puzzle-hash\",\"Desc\":\"Puzzle hash to search for\",\"Type\":24,\"Default\":null,\"Path\":\"puzzle_hash\",\"Data\":null},{\"Name\":\"start-height\",\"Desc\":\"Confirmation start height for search\",\"Type\":2,\"Default\":null,\"Path\":\"start_height\",\"Data\":null},{\"Name\":\"end-height\",\"Desc\":\"Confirmation end height for search\",\"Type\":2,\"Default\":null,\"Path\":\"end_height\",\"Data\":null},{\"Name\":\"include-spend-coins\",\"Desc\":\"Whether to include spent coins too, instead of just unspent\",\"Type\":1,\"Default\":null,\"Path\":\"include_spend_coins\",\"Data\":null}]}"), req)
	req.ValInfo[0].Data = &puzzleHash
	req.ValInfo[1].Data = &startHeight
	req.ValInfo[2].Data = &endHeight
	req.ValInfo[3].Data = &includeSpendCoins
	return req
}

// GetCoinRecordByName Retrieves a coin record by its name/id
func GetCoinRecordByName(name string) client.RpcMethod {
	req := &client.TemplateRpcMethod{}
	_ = json.Unmarshal([]byte("{\"MethodName\":\"GetCoinRecordByName\",\"Desc\":\"Retrieves a coin record by its name/id\",\"Method\":\"POST\",\"Url\":\"get_coin_record_by_name\",\"JsonTemplate\":\"{\\\"name\\\": \\\"\\\"}\",\"ValInfo\":[{\"Name\":\"name\",\"Desc\":\"Coin name\",\"Type\":24,\"Default\":null,\"Path\":\"name\",\"Data\":null}]}"), req)
	req.ValInfo[0].Data = &name
	return req
}

// PushTx Pushes a transaction / spend bundle to the mempool and blockchain. Returns whether the spend bundle was successfully included into the mempool
func PushTx(spendBundle string) client.RpcMethod {
	req := &client.TemplateRpcMethod{}
	_ = json.Unmarshal([]byte("{\"MethodName\":\"PushTx\",\"Desc\":\"Pushes a transaction / spend bundle to the mempool and blockchain. Returns whether the spend bundle was successfully included into the mempool\",\"Method\":\"POST\",\"Url\":\"push_tx\",\"JsonTemplate\":\"{\\\"spend_bundle\\\": \\\"\\\"}\",\"ValInfo\":[{\"Name\":\"spend-bundle\",\"Desc\":\"Spend bundle to submit, in JSON\",\"Type\":24,\"Default\":null,\"Path\":\"spend_bundle\",\"Data\":null}]}"), req)
	req.ValInfo[0].Data = &spendBundle
	req.ValInfo[0].FormatFunc = RpcApis[13].ValInfo[0].FormatFunc
	return req
}

// GetAllMempoolTxIds Returns a list of all transaction IDs(spend bundle hashes) in the mempool
func GetAllMempoolTxIds() client.RpcMethod {
	req := &client.TemplateRpcMethod{}
	_ = json.Unmarshal([]byte("{\"MethodName\":\"GetAllMempoolTxIds\",\"Desc\":\"Returns a list of all transaction IDs(spend bundle hashes) in the mempool\",\"Method\":\"POST\",\"Url\":\"get_all_mempool_tx_ids\",\"JsonTemplate\":\"{}\",\"ValInfo\":null}"), req)
	return req
}

// GetAllMempoolItems Returns all items in the mempool
func GetAllMempoolItems() client.RpcMethod {
	req := &client.TemplateRpcMethod{}
	_ = json.Unmarshal([]byte("{\"MethodName\":\"GetAllMempoolItems\",\"Desc\":\"Returns all items in the mempool\",\"Method\":\"POST\",\"Url\":\"get_all_mempool_items\",\"JsonTemplate\":\"{}\",\"ValInfo\":null}"), req)
	return req
}

// GetMempoolItemByTxId Gets a mempool item by tx id
func GetMempoolItemByTxId(txId string) client.RpcMethod {
	req := &client.TemplateRpcMethod{}
	_ = json.Unmarshal([]byte("{\"MethodName\":\"GetMempoolItemByTxId\",\"Desc\":\"Gets a mempool item by tx id\",\"Method\":\"POST\",\"Url\":\"get_mempool_item_by_tx_id\",\"JsonTemplate\":\"{\\\"tx_id\\\": \\\"\\\"}\",\"ValInfo\":[{\"Name\":\"tx-id\",\"Desc\":\"Spend bundle hash\",\"Type\":24,\"Default\":null,\"Path\":\"tx_id\",\"Data\":null}]}"), req)
	req.ValInfo[0].Data = &txId
	return req
}
