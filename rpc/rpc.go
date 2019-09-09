package rpc

import (
	"encoding/json"
	"omni-scan/models"
)

type Block struct {
	Height                 int64 `json:"block"`
	Timestamp              int64 `json:"blocktime"`
	OmniTransactionsAmount int64 `json:"blocktransactions"`
}

func GetLatestBlockInfo() (Block, error) {
	var block Block
	result, err := DefaultClient.SendJsonRpc("omni_getinfo")
	if err != nil {
		return block, err
	}
	if err = json.Unmarshal(result, &block); err != nil {
		return block, err
	}
	return block, err
}

// index can be block height or block index
func GetBlockTransactions(index int64) ([]models.Transaction, error) {
	var transactionList []models.Transaction
	result, err := DefaultClient.SendJsonRpc("omni_listblocktransactions", index)
	if err != nil {
		return transactionList, err
	}
	if err = json.Unmarshal(result, &transactionList); err != nil {
		return transactionList, err
	}
	return transactionList, nil
}
