package full_node

import (
	"chia-go-cli/sdk/client"
	"reflect"
)

var RpcApis = []*client.TemplateRpcMethod{
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
		ValInfo: []*client.TemplateValue{
			{
				Name: "header-hash",
				Desc: "Header hash",
				Type: reflect.String,
				Path: "header_hash",
			},
		},
	},
	{
		MethodName:   "GetBlocks",
		Desc:         "Gets a list of full blocks",
		Method:       "POST",
		Url:          "get_blocks",
		JsonTemplate: `{"start": 0, "end": 9999, "exclude_header_hash": false}`,
		ValInfo: []*client.TemplateValue{
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
	{
		MethodName:   "GetBlockRecordByHeight",
		Desc:         "Retrieves a block record by height (assuming the height <= peak height)",
		Method:       "POST",
		Url:          "get_block_record_by_height",
		JsonTemplate: `{"height": 0}`,
		ValInfo: []*client.TemplateValue{
			{
				Name: "height",
				Desc: "The height to get",
				Type: reflect.Int,
				Path: "height",
			},
		},
	},
	{
		MethodName:   "GetBlockRecord",
		Desc:         "Retrieves a block record by header hash",
		Method:       "POST",
		Url:          "get_block_record",
		JsonTemplate: `{"header_hash": ""}`,
		ValInfo: []*client.TemplateValue{
			{
				Name: "header-hash",
				Desc: "The block's header_hash",
				Type: reflect.String,
				Path: "header_hash",
			},
		},
	},
	{
		MethodName:   "GetBlockRecords",
		Desc:         "Retrieves block records in a range",
		Method:       "POST",
		Url:          "get_block_records",
		JsonTemplate: `{"start": 0, "end": 9999}`,
		ValInfo: []*client.TemplateValue{
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
		},
	},
	{
		MethodName:   "GetUnfinishedBlockHeaders",
		Desc:         "Retrieves recent unfinished header blocks",
		Method:       "POST",
		Url:          "get_unfinished_block_headers",
		JsonTemplate: `{}`,
	},
	{
		MethodName:   "GetNetworkSpace",
		Desc:         "Retrieves an estimate of the total plotted space of all farmers, in bytes",
		Method:       "POST",
		Url:          "get_network_space",
		JsonTemplate: `{"older_block_header_hash": "", "newer_block_header_hash": ""}`,
		ValInfo: []*client.TemplateValue{
			{
				Name: "older-block-header-hash",
				Desc: "The start header hash",
				Type: reflect.String,
				Path: "older_block_header_hash",
			},
			{
				Name: "newer-block-header-hash",
				Desc: "The end header hash",
				Type: reflect.String,
				Path: "newer_block_header_hash",
			},
		},
	},
	{
		MethodName:   "GetAdditionsAndRemovals",
		Desc:         "Retrieves the additions and removals(state transitions) for a certain block. Returns coin records for each addition and removal",
		Method:       "POST",
		Url:          "get_additions_and_removals",
		JsonTemplate: `{"header_hash": ""}`,
		ValInfo: []*client.TemplateValue{
			{
				Name: "header-hash",
				Desc: "Header hash of the block",
				Type: reflect.String,
				Path: "header_hash",
			},
		},
	},
	{
		MethodName:   "GetInitialFreezePeriod",
		Desc:         "Retrieves the initial freeze period for the blockchain (no transactions allowed)",
		Method:       "POST",
		Url:          "get_initial_freeze_period",
		JsonTemplate: `{}`,
	},
	{
		MethodName:   "GetNetworkInfo",
		Desc:         "Retrieves some information about the current network",
		Method:       "POST",
		Url:          "get_network_info",
		JsonTemplate: `{}`,
	},
	{
		MethodName:   "GetCoinRecordsByPuzzleHash",
		Desc:         "Retrieves a list of coin records with a certain puzzle hash",
		Method:       "POST",
		Url:          "get_coin_records_by_puzzle_hash",
		JsonTemplate: `{"puzzle_hash": "", "start_height": 0, "end_height": 0, "include_spend_coins": false}`,
		ValInfo: []*client.TemplateValue{
			{
				Name: "puzzle-hash",
				Desc: "Puzzle hash to search for",
				Type: reflect.String,
				Path: "puzzle_hash",
			},
			{
				Name: "start-height",
				Desc: "Confirmation start height for search",
				Type: reflect.Int,
				Path: "start_height",
			},
			{
				Name: "end-height",
				Desc: "Confirmation end height for search",
				Type: reflect.Int,
				Path: "end_height",
			},
			{
				Name: "include-spend-coins",
				Desc: "Whether to include spent coins too, instead of just unspent",
				Type: reflect.Bool,
				Path: "include_spend_coins",
			},
		},
	},
	{
		MethodName:   "GetCoinRecordByName",
		Desc:         "Retrieves a coin record by its name/id",
		Method:       "POST",
		Url:          "get_coin_record_by_name",
		JsonTemplate: `{"name": ""}`,
		ValInfo: []*client.TemplateValue{
			{
				Name: "name",
				Desc: "Coin name",
				Type: reflect.String,
				Path: "name",
			},
		},
	},
	{
		MethodName:   "PushTx",
		Desc:         "Pushes a transaction / spend bundle to the mempool and blockchain. Returns whether the spend bundle was successfully included into the mempool",
		Method:       "POST",
		Url:          "push_tx",
		JsonTemplate: `{"spend_bundle": ""}`,
		ValInfo: []*client.TemplateValue{
			{
				Name:       "spend-bundle",
				Desc:       "Spend bundle to submit, in JSON",
				Type:       reflect.String,
				Path:       "spend_bundle",
				FormatFunc: client.String2JsonFormatFunc,
			},
		},
	},
	{
		MethodName:   "GetAllMempoolTxIds",
		Desc:         "Returns a list of all transaction IDs(spend bundle hashes) in the mempool",
		Method:       "POST",
		Url:          "get_all_mempool_tx_ids",
		JsonTemplate: `{}`,
	},
	{
		MethodName:   "GetAllMempoolItems",
		Desc:         "Returns all items in the mempool",
		Method:       "POST",
		Url:          "get_all_mempool_items",
		JsonTemplate: `{}`,
	},
	{
		MethodName:   "GetMempoolItemByTxId",
		Desc:         "Gets a mempool item by tx id",
		Method:       "POST",
		Url:          "get_mempool_item_by_tx_id",
		JsonTemplate: `{"tx_id": ""}`,
		ValInfo: []*client.TemplateValue{
			{
				Name: "tx-id",
				Desc: "Spend bundle hash",
				Type: reflect.String,
				Path: "tx_id",
			},
		},
	},
}
