package models

import "github.com/golang/protobuf/proto"

const (
	TxKeyPrefix        = "tx-"
	ContextKey         = "context"
)

//key:  tx-{tx.txid}
func TxKey(txid string) string {
	return TxKeyPrefix + txid
}

func Encode(v proto.Message) ([]byte, error) {
	return proto.Marshal(v)
}

func Decode(data []byte, v proto.Message) error {
	return proto.Unmarshal(data, v)
}

