package logic

import (
	"github.com/DalongWallet/omni-scan/models"
	"sync"
)

type Block struct {
	Height 		int64
	Hash  		string
	Txs 		[]*models.Transaction
}

type txPosition struct {
	BlockHeihgt 	int64
	Index 			int
}

type UnconfirmedBlockMgr struct {
	txGetter  		TxGetter
	blocks  		map[int64]*Block
	txIndex 		map[string]*txPosition
	txAddrs 		map[string]map[string]struct{}
	addrIndex  		map[string]map[string]*models.Transaction
	maxBlockHeight 	int64
	mutex 			sync.RWMutex
}

func NewUnconfirmedBlockMgr(txGetter TxGetter) *UnconfirmedBlockMgr {
	return &UnconfirmedBlockMgr{
		txGetter: 	txGetter,
		blocks: 	make(map[int64]*Block),
		txIndex: 	make(map[string]*txPosition),
		txAddrs: 	make(map[string]map[string]struct{}),
		addrIndex: 	make(map[string]map[string]*models.Transaction),
	}
}

//func (m *UnconfirmedBlockMgr) RemoveBlock(height int64) error {
//	m.mutex.Lock()
//	defer m.mutex.Unlock()
//
//	m.removeBlock(height)
//
//	return nil
//}
//
//func (m *UnconfirmedBlockMgr) AddBlock(block *Block) error {
//	m.mutex.Lock()
//	defer m.mutex.Unlock()
//
//	temp
//
//}
