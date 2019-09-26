package rpc

import (
	"github.com/json-iterator/go"
	"github.com/pkg/errors"
	"github.com/DalongWallet/omni-scan/models"
	"reflect"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func (client *OmniClient) GetLatestBlockInfo() (block models.OmniInfoResult, err error) {
	cmd := GetInfoCommand{}

	var result []byte
	if result, err = client.Exec(cmd); err != nil {
		return block, err
	}

	err = unmarshal(result, &block, "GetLatestBlockInfo")
	return
}

// index can be block height or block index
func (client *OmniClient) ListBlockTransactions(index int) (txIdList []string, err error) {
	cmd := ListBlockTransactionsCommand{
		Index: index,
	}

	var result []byte
	if result, err = client.Exec(cmd); err != nil {
		return txIdList, err
	}

	err = unmarshal(result, &txIdList, "ListBlockTransactions")
	return
}

func (client *OmniClient) ListBlocksTransactions(firstBlock, lastBlock int64) (txIdList []string, err error) {
	cmd := ListBlocksTransactionsCommand{
		FirstBlock: firstBlock,
		LastBlock:  lastBlock,
	}

	var result []byte
	if result, err = client.Exec(cmd); err != nil {
		return txIdList, err
	}

	err = unmarshal(result, &txIdList, "ListBlocksTransactions")
	return
}

func (client *OmniClient) GetTransaction(txId string) (tx models.Transaction, err error) {
	cmd := GetTransactionCommand{
		TxId: txId,
	}

	var result []byte
	if result, err = client.Exec(cmd); err != nil {
		return
	}

	err = unmarshal(result, &tx, "GetTransaction")
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

	err = unmarshal(result, &tokenBalance, "GetBalance")
	return
}

func (client *OmniClient) GetAllBalancesForId(propertyId int) (addrTokenBalanceList []models.AddressTokenBalance, err error) {
	cmd := GetAllbalancesForIdCommand{
		PropertyId: propertyId,
	}

	var result []byte
	if result, err = client.Exec(cmd); err != nil {
		return
	}

	err = unmarshal(result, &addrTokenBalanceList, "GetAllBalancesForId")
	return
}

func (client *OmniClient) GetAllBalancesForAddress(address string) (propertyTokenBalanceList []models.PropertyTokenBalance, err error) {
	cmd := GetAllBalancesForAddressCommand{
		Address: address,
	}

	var result []byte
	if result, err = client.Exec(cmd); err != nil {
		if rpcErr, ok := errors.Cause(err).(*rpcError); ok {
			if reflect.DeepEqual(rpcErr, ErrAddressNotFound) {
				err = nil
			}
		}
		return
	}

	err = unmarshal(result, &propertyTokenBalanceList, "GetAllBalancesForAddress")
	return
}

func (client *OmniClient) GetPropertyBalanceForAddress(address string, propertyId int) (propertyBalance models.PropertyTokenBalance, err error) {
	var allPropertyBalances []models.PropertyTokenBalance
	if allPropertyBalances, err = client.GetAllBalancesForAddress(address); err != nil {
		return
	}
	for _, one := range allPropertyBalances {
		if one.PropertyId == propertyId {
			propertyBalance = one
			return
		}
	}
	return
}

func (client *OmniClient) SendRawTransaction(from string, hex string) (txHash string, err error) {
	cmd := SendRawTransactionCommand {
		FromAddress: 	from,
		Hex: 			hex,
	}

	var result []byte
	if result, err = client.Exec(cmd); err != nil {
		return
	}
	txHash = string(result)
	return
}

func (client *OmniClient) DecodeTransaction(rawtx string) (tx models.Transaction, err error){
	cmd := DecodeRawTransactionCommand {
		RawTx: 				rawtx,
	}

	var result []byte
	if result, err = client.Exec(cmd); err != nil {
		return
	}

	err = json.Unmarshal(result, &tx)
	if err != nil {
		return
	}

	return tx, nil
}

func unmarshal(data []byte, v interface{}, errMsg string) error {
	return errors.Wrap(json.Unmarshal(data, v), errMsg)
}
