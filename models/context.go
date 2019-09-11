package models

import (
	"github.com/golang/protobuf/proto"
	"github.com/DalongWallet/omni-scan/storage"
)

// key:  context
type Context struct {
	BlockHeight 		int64 `json:"block_height"`
	BlockHash  			string `json:"block_hash"`
	BlockTotalTx   		int32 `json:"block_total_tx"`
	BlockProcessedTx 	int32 `json:"block_processed_tx"`
	TotalTx 			int64 `json:"total_tx"`
	TotalTxNode 		int64 `json:"total_tx_node"`
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

func (m *Context) GetBlockHash() string {
	if m != nil {
		return m.BlockHash
	}
	return ""
}

func (m *Context) GetBlockTotalTx() int32 {
	if m != nil {
		return m.BlockTotalTx
	}
	return 0
}

func (m *Context) GetBlockProcessedTx() int32 {
	if m != nil {
		return m.BlockProcessedTx
	}
	return 0
}

func (m *Context) GetTotalTx() int64 {
	if m != nil {
		return m.TotalTx
	}
	return 0
}

func (m *Context) GetTotalTxNode() int64 {
	if m != nil {
		return m.TotalTxNode
	}
	return 0
}

func (m *Context) Encode() ([]byte, error) {
	return Encode(m)
}

func (m *Context) Decode(data []byte) error {
	return Decode(data, m)
}

func (m *Context) Save(store storage.Storage, key string) error {
	data, err := m.Encode()
	if err != nil {
		return err
	}
	if key == "" {
		key = ContextKey
	}
	return store.Set(key, string(data))
}

func (m *Context) Load(store storage.Storage, key string) error {
	if key == "" {
		key = ContextKey
	}
	data, err := store.Get(key)
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

func (m *Context) Clone() *Context {
	newCtx := &Context{
		BlockHeight:      m.BlockHeight,
		BlockHash:        m.BlockHash,
		BlockTotalTx:     m.BlockTotalTx,
		BlockProcessedTx: m.BlockProcessedTx,
		TotalTx:          m.TotalTx,
		TotalTxNode:      m.TotalTxNode,
	}
	return newCtx
}