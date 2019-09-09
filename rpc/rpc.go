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

func (client *OmniClient) GetLatestBlockInfo() (block models.OmniInfoResult, err error) {
	cmd := GetInfoCommand{}

	var result []byte
	if result, err = client.Exec(cmd); err != nil {
		return block, err
	}

	err = json.Unmarshal(result, &block)
	return
}

// index can be block height or block index
func (client *OmniClient) ListBlockTransactions(index int)(txIdList []string, err error) {
	cmd := ListBlockTransactionsCommand{
		Index: index,
	}

	var result []byte
	if result, err = client.Exec(cmd); err != nil {
		return txIdList, err
	}

	err = json.Unmarshal(result, &txIdList)
	return
}

func (client *OmniClient) ListBlocksTransactions(firstBlock, lastBlock int64)(txIdList []string, err error) {
	cmd := ListBlocksTransactionsCommand{
		FirstBlock: firstBlock,
		LastBlock: lastBlock,
	}

	var result []byte
	if result, err = client.Exec(cmd); err != nil {
		return txIdList, err
	}

	err = json.Unmarshal(result, &txIdList)
	return
}


func (client *OmniClient) GetTransaction(txId string)(tx models.Transaction, err error) {
	cmd := GetTransactionCommand{
		TxId: txId,
	}

	var result []byte
	if result, err = client.Exec(cmd); err != nil {
		return
	}

	err = json.Unmarshal(result, &tx)
	return
}

func (client *OmniClient) GetBalance(address string, propertyId int) (tokenBalance models.TokenBalance, err error) {
	cmd := GetBalanceCommand{
		Address:    address,
		PropertyId: propertyId,
	}

	var result []byte
	if result, err = client.Exec(cmd); err != nil {
		return
	}

	err = json.Unmarshal(result, &tokenBalance)
	return
}

func (client *OmniClient) GetAllBalancesForId(propertyId int) (addrTokenBalanceList []models.AddressTokenBalance, err error) {
	cmd := GetAllbalancesForIdCommand{
		PropertyId:propertyId,
	}

	var result []byte
	if result, err = client.Exec(cmd); err != nil {
		return
	}

	err = json.Unmarshal(result, &addrTokenBalanceList)
	return
}

func (client *OmniClient) GetAllBalancesForAddress(address string) (propertyTokenBalanceList []models.PropertyTokenBalance, err error) {
	cmd := GetAllBalancesForAddressCommand{
		Address:address,
	}

	var result []byte
	if result, err = client.Exec(cmd); err != nil {
		return
	}

	err = json.Unmarshal(result, &propertyTokenBalanceList)
	return
}
