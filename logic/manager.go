package logic

import (
	"github.com/DalongWallet/omni-scan/models"
	"github.com/DalongWallet/omni-scan/storage/leveldb"
	"sort"
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

func NewOmniMgr(st *leveldb.LevelStorage, mempoolSize int) (*OmniMgr, error) {
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

func (m *OmniMgr) GetAddressConfirmedTxs(address string, propertyId int, limit uint, offset uint) (confirmTxs []*models.Transaction, err error) {
	confirmTxs, err = m.ConfirmedTx.GetAddressTxs(address, propertyId)
	if err != nil || len(confirmTxs) == 0 {
		return
	}

	start, end := int(offset), int(offset+limit)
	if start > len(confirmTxs) {
		return []*models.Transaction{}, nil
	}
	if end > len(confirmTxs) {
		end = len(confirmTxs)
	}

	txs := models.TxByTimeDescSlice(confirmTxs)
	sort.Sort(txs)
	confirmTxs = txs[start:end]
	return
}

func (m *OmniMgr) GetAddressBalance(address string, propertyId int) (*models.AddressTokenBalance, error) {
	return nil, nil
}


