package models

import (
	"fmt"
	"github.com/golang/protobuf/proto"
)

const (
	TxKeyPrefix        = "tx-"
	BalanceKeyPrefix   = "balance-"
	ContextKey         = "context"
)

// key:  tx-{tx.txid}
func TxKey(txid string) string {
	return TxKeyPrefix + txid
}

// key: tx-{addr}-{propertyId}-{txId}
func AddrPropertyTxKey(addr string, propertyId int, txId string) string {
	return fmt.Sprintf("%s-%s-%d-%s",TxKeyPrefix, addr, propertyId, txId)
}

// key: tx-{addr}-{propertyId}
func AddrPropertyTxsKey(addr string, propertyId int) string {
	return fmt.Sprintf("%s-%s-%d",TxKeyPrefix, addr, propertyId)
}

// key: balance-{addr}-{propertyId}
func AddrPropertyBalanceKey(addr string, propertyId int) string {
	return fmt.Sprintf("%s-%s-%d", BalanceKeyPrefix, addr, propertyId)
}

func Encode(v proto.Message) ([]byte, error) {
	return proto.Marshal(v)
}

func Decode(data []byte, v proto.Message) error {
	return proto.Unmarshal(data, v)
}

