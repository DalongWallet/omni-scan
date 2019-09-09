package rpc

import (
	"encoding/json"
	"github.com/pkg/errors"
	"omni-scan/models"
)

type Block struct {
	Height                 int64 `json:"block"`
	Timestamp              int64 `json:"blocktime"`
	OmniTransactionsAmount int64 `json:"blocktransactions"`
}

func GetLatestBlockInfo() (block Block, err error) {
	var result []byte
	if result, err = DefaultRpcClient.SendJsonRpc("omni_getinfo"); err != nil {
		return block, err
	}
	err = json.Unmarshal(result, &block)
	err = errors.Wrap(err, "Unmarshal data to struct Block failed")
	return
}

// index can be block height or block index
func ListBlockTransactions(index int64)(txHashList []string, err error) {
	var result []byte
	if result, err = DefaultRpcClient.SendJsonRpc("omni_listblocktransactions", index); err != nil {
		return txHashList, err
	}
	err = json.Unmarshal(result, &txHashList)
	err = errors.Wrap(err, "Unmarshal data to txHashList failed")
	return
}

func GetTransaction(txHash string)(tx models.Transaction, err error) {
	var result []byte
	if result, err = DefaultRpcClient.SendJsonRpc("omni_gettransaction"); err != nil {
		return
	}
	err = json.Unmarshal(result, &tx)
	err = errors.Wrap(err, "Unmarshal data to struct Transaction failed")
	return
}

func GetBlockTransactions(index int64) (txList []models.Transaction,err error) {
	var result []byte
	if result, err = DefaultRpcClient.SendJsonRpc("omni_listblocktransactions", index); err != nil {
		return
	}
	err = json.Unmarshal(result, &txList)
	err = errors.Wrap(err, "Unmarshal data to []models.Transaction failed")
	return
}

func GetBalance(address string, propertyId int) (tokenBalance models.TokenBalance, err error) {
	var result []byte
	if result, err = DefaultRpcClient.SendJsonRpc("omni_getbalance", address); err != nil {
		return
	}
	err = json.Unmarshal(result, &tokenBalance)
	err = errors.Wrap(err, "Unmarshal data to struct tokenBalance failed")
	return
}

func GetAllBalancesForId(propertyId int) (addrTokenBalanceList []models.AddressTokenBalance, err error) {
	var result []byte
	if result, err = DefaultRpcClient.SendJsonRpc("omni_getallbalancesforid", propertyId); err != nil {
		return
	}
	err = json.Unmarshal(result, &addrTokenBalanceList)
	err = errors.Wrap(err, "Unmarshal data to []models.AddressTokenBalance failed")
	return
}

func GetAllBalancesForAddress(address string) (propertyTokenBalanceList []models.PropertyTokenBalance, err error) {
	var result []byte
	if result, err = DefaultRpcClient.SendJsonRpc("omni_getallbalancesforaddress", address); err != nil {
		return
	}
	err = json.Unmarshal(result, &propertyTokenBalanceList)
	err = errors.Wrap(err, "Unmarshal data to []models.PropertyTokenBalancen failed")
	return
}
