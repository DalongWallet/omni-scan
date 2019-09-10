package logic

import (
	"omni-scan/models"
	"omni-scan/storage"
)

type TxGetter interface {
	GetTx(txid string) (*models.Transaction, error)
}

// TODO pending tx collect
type OmniMgr struct {
	//UnconfirmedBlock 	*UnconfirmedBlockMgr
	ConfirmedTx *ConfirmedTxMgr
	MemPool     *MemPoolMgr
}

func NewOmniMgr(st storage.Storage, mempoolSize int) (*OmniMgr, error) {
	m := &OmniMgr{}
	cfBlock, err := NewConfirmedBlockMgr(st)
	if err != nil {
		return nil, err
	}
	m.ConfirmedTx = cfBlock
	m.MemPool = NewMemPool(m, mempoolSize)
	return m, nil
}

func (m *OmniMgr) GetTx(txid string) (*models.Transaction, error) {
	tx, err := m.ConfirmedTx.GetTx(txid)
	if err == nil {
		return tx, nil
	}
	return nil, models.ErrorNotFound
}

func (m *OmniMgr) GetAddressConfirmedTxs(address string, limit uint, offset uint) ([]*models.Transaction, error) {
	var out []*models.Transaction
	return out, nil
}

func (m *OmniMgr) GetAddressBalance(address string, propertyId int) (*models.AddressTokenBalance, error) {
	return nil, nil
}


