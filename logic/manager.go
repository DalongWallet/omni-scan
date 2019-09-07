package logic

import (
	"omni-scan/models"
	"omni-scan/storage"
)

type TxGetter interface {
	GetTx(txid string) (*models.Transaction, error)
}

// TODO pending tx collect
type TxMgr struct {
	//UnconfirmedBlock 	*UnconfirmedBlockMgr
	ConfirmedTx *ConfirmedTxMgr
	MemPool     *MemPoolMgr
}

func NewTxMgr(st storage.Storage, mempoolSize int) (*TxMgr, error) {
	m := &TxMgr{}
	cfBlock, err := NewConfirmedBlockMgr(st)
	if err != nil {
		return nil, err
	}
	m.ConfirmedTx = cfBlock
	m.MemPool = NewMemPool(m, mempoolSize)
	return m, nil
}

func (m *TxMgr) GetTx(txid string) (*models.Transaction, error) {
	tx, err := m.ConfirmedTx.GetTx(txid)
	if err != nil {
		return nil, models.ErrorNotFound
	}
	return tx, nil
}

