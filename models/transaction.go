package models

import "omni-scan/storage"

/*
{
  "txid": "2f51deda0dc3d93ee3f05a0d5c8e339e1e290b00784d2617f86c8d0391747fed",
  "fee": "0.00011000",
  "sendingaddress": "1Pr75FNvtoWHeocNfc4zTQCfK5kMVakWcn",
  "referenceaddress": "1QG3Jw5Mf76sQTcvQoe1cqBbX4g5z6ixXb",
  "ismine": false,
  "version": 0,
  "type_int": 0,
  "type": "Simple Send",
  "propertyid": 2,
  "divisible": true,
  "amount": "0.01000000",
  "valid": true,
  "blockhash": "00000000000000004f85dc04a43c3c3bfd26c6c06acf3cd1e2bb39510745c144",
  "blocktime": 1399750635,
  "positioninblock": 189,
  "block": 300092,
  "confirmations": 64104
}
 */
type Transaction struct {
	TxId  			string `json:"txid"`
	Fee 			string `json:"fee"`
	SendingAddress 	string `json:"sendingaddress"`
	ReferenceAddress string `json:"referenceaddress"`
	IsMine 			bool `json:"ismine"`
	Version 		int `json:"version"`
	TypeInt 		int `json:"type_int"`
	Type 			string `json:"type"`
	PropertyId 		int `json:"propertyid"`
	Divisible  		bool `json:"divisible"`
	Amount 			string `json:"amount"`
	Valid  			bool `json:"valid"`
	BlockHash 		string `json:"blockhash"`
	BlockTime 		int64 `json:"blocktime"`
	PositionInBlock int `json:"positioninblock"`
	Block 			int64 `json:"block"`
	Confirmations  	int64 `json:"confirmations"`
}

func (m *Transaction) Encode() ([]byte, error) {
	return Encode(m)
}

func (m *Transaction) Decode(data []byte) error {
	return Decode(data, m)
}

func (m *Transaction) Save(store storage.Storage) error {
	data, err := m.Encode()
	if err != nil {
		return err
	}
	return store.Set(TxKey(m.TxId), string(data))
}

func (m *Transaction) Load(store storage.Storage, txid string) error {
	data, err := store.Get(TxKey(txid))
	if err != nil {
		return err
	}
	if data == "" {
		return ErrorNotFound
	}
	err = m.Decode([]byte(data))
	if err != nil {
		return err
	}
	return nil
}

type TxByTxidSlice []*Transaction

func (c TxByTxidSlice) Len() int {
	return len(c)
}
func (c TxByTxidSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c TxByTxidSlice) Less(i, j int) bool {
	return c[i].TxId < c[j].TxId
}

type TxByTimeSlice []*Transaction

func (c TxByTimeSlice) Len() int {
	return len(c)
}
func (c TxByTimeSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c TxByTimeSlice) Less(i, j int) bool {
	return c[i].BlockTime < c[j].BlockTime
}

type TxByTimeDescSlice []*Transaction

func (c TxByTimeDescSlice) Len() int {
	return len(c)
}
func (c TxByTimeDescSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c TxByTimeDescSlice) Less(i, j int) bool {
	return c[i].BlockTime > c[j].BlockTime
}