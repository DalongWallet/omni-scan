package logic

import (
	"omni-scan/models"
	"omni-scan/storage"
)

type ConfirmedTxMgr struct {
	storage  		storage.Storage
	ctx  			*models.Context
}

type TxAddr struct {
	addr 		string
	Type 		uint32
}

func NewConfirmedBlockMgr(storage storage.Storage) (*ConfirmedTxMgr, error) {
	m := ConfirmedTxMgr {
		storage:  storage,
	}
	ctx := &models.Context{}
	err := ctx.Load(storage, "")
	if err != nil && err != models.ErrorNotFound {
		return nil, err
	}
	m.ctx = ctx

	return &m, nil
}

func (m *ConfirmedTxMgr) GetContext() *models.Context {
	return m.ctx
}

func (m *ConfirmedTxMgr) GetTx(txid string) (*models.Transaction, error) {
	t := &models.Transaction{}
	err := t.Load(m.storage, txid)
	return t, err
}

func (m *ConfirmedTxMgr) SaveTx(tx *models.Transaction) error {
	err := tx.Save(m.storage)
	return err
}

func (m *ConfirmedTxMgr) GetAddressTxs(addr string, limit uint, order int, preKey string) {

}

