package logic

import (
	"github.com/DalongWallet/omni-scan/models"
	"github.com/DalongWallet/omni-scan/storage/leveldb"
	"github.com/DalongWallet/omni-scan/utils"
	jsoniter "github.com/json-iterator/go"
)

type ConfirmedTxMgr struct {
	storage  		*leveldb.LevelStorage
	ctx  			*models.Context
}

type TxAddr struct {
	addr 		string
	Type 		uint32
}

func NewConfirmedBlockMgr(storage *leveldb.LevelStorage) (*ConfirmedTxMgr, error) {
	m := ConfirmedTxMgr {
		storage:  storage,
	}
	//ctx := &models.Context{}
	//err := ctx.Load(storage)
	//if err != nil && err != models.ErrorNotFound {
	//	return nil, err
	//}
	//m.ctx = ctx

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

func (m *ConfirmedTxMgr) GetAddressTxs(addr string, propertyId int) (txs []*models.Transaction, err error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var data [][]byte

	if data, err = m.storage.GetWithPrefix(models.AddrPropertyTxsKey(addr, propertyId)); err != nil {
		if utils.IsErrorNotFound(err) {
			err = nil
		}
		return
	}

	for _, one := range data {
		tx := &models.Transaction{}
		if err = json.Unmarshal(one, tx); err != nil {
			return
		}
		txs = append(txs, tx)
	}
	return
}

