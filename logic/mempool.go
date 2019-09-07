package logic

import (
	"container/list"
	"omni-scan/models"
	"sync"
)

type MemPoolMgr struct {
	txQueue   	*list.List
	txGetter 	TxGetter
	txs 		map[string]*list.Element
	txAddrs 	map[string]map[string]struct{}
	addrTxs 	map[string]map[string]*models.Transaction
	mutex  		sync.RWMutex

	mempoolSize int
}

func NewMemPool(txGeter TxGetter, mempoolSize int) *MemPoolMgr {
	m := &MemPoolMgr{txGetter: txGeter,
		txQueue:     list.New(),
		txs:         make(map[string]*list.Element),
		txAddrs:     make(map[string]map[string]struct{}),
		addrTxs:     make(map[string]map[string]*models.Transaction),
		mempoolSize: mempoolSize,
	}
	return m
}
