package models

import (
	"github.com/DalongWallet/omni-scan/storage/leveldb"
	"github.com/DalongWallet/omni-scan/utils"
)

type TokenBalance struct {
	Balance string `json:"balance"`
	Reserved string `json:"reserved"`
	Frozen string `json:"frozen"`
}

type AddressTokenBalance struct {
	Address string  `json:"address"`
	TokenBalance
}

type PropertyTokenBalance struct {
	PropertyId int `json:"propertyid"`
	Name string `json:"name"`
	TokenBalance
}


func (b *PropertyTokenBalance) Load(store *leveldb.LevelStorage, addr string, propertyId int) (error) {
	return utils.Load(store, AddrPropertyBalanceKey(addr, propertyId), b)
}