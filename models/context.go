package models

import (
	"github.com/DalongWallet/omni-scan/storage/leveldb"
	"github.com/golang/protobuf/proto"
)

// key:  context
type Context struct {
	OmniCoreVersion    string `json:"omnicoreversion"`
	BitcoinCoreVersion string `json:"bitcoincoreversion"`
	BlockHeight        int64  `json:"block"`
	BlockTime          int64  `json:"blocktime"`
	BlockTransactions  int64  `json:"blocktransactions"`
	TotalTrades        int64  `json:"totaltrades"`
	TotalTransactions  int64  `json:"totaltransactions"`
}

func (m *Context) Reset() {
	*m = Context{}
}

func (m *Context) String() string {
	return proto.CompactTextString(m)
}

func (m *Context) ProtoMessage() {}

func (m *Context) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{1}
}

func (m *Context) GetBlockHeight() int64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *Context) GetOmniCoreVersion() string {
	if m != nil {
		return m.OmniCoreVersion
	}
	return ""
}

func (m *Context) GetBitcoinCoreVersion() string {
	if m != nil {
		return m.BitcoinCoreVersion
	}
	return ""
}

func (m *Context) GetBlockTime() int64 {
	if m != nil {
		return m.BlockTime
	}
	return 0
}

func (m *Context) GetBlockTransactions() int64 {
	if m != nil {
		return m.BlockTransactions
	}
	return 0
}

func (m *Context) GetTotalTransactions() int64 {
	if m != nil {
		return m.TotalTransactions
	}
	return 0
}

func (m *Context) GetTotalTrades() int64 {
	if m != nil {
		return m.TotalTrades
	}
	return 0
}

func (m *Context) Encode() ([]byte, error) {
	return Encode(m)
}

func (m *Context) Decode(data []byte) error {
	return Decode(data, m)
}

func (m *Context) Save(store *leveldb.LevelStorage) error {
	data, err := m.Encode()
	if err != nil {
		return err
	}
	return store.Set(ContextKey, data)
}

func (m *Context) Load(store *leveldb.LevelStorage) error {
	data, err := store.Get(ContextKey)
	if err != nil {
		return err
	}
	if len(data) == 0 {
		return ErrorNotFound
	}
	err = m.Decode(data)
	if err != nil {
		return err
	}
	return nil
}

func (m *Context) Clone() *Context {
	newCtx := &Context{
		OmniCoreVersion:    m.OmniCoreVersion,
		BitcoinCoreVersion: m.BitcoinCoreVersion,
		BlockHeight:        m.BlockHeight,
		BlockTime:          m.BlockTime,
		BlockTransactions:  m.BlockTransactions,
		TotalTrades:        m.TotalTrades,
		TotalTransactions:  m.TotalTransactions,
	}
	return newCtx
}
