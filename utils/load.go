package utils

import (
	"github.com/DalongWallet/omni-scan/storage/leveldb"
	jsoniter "github.com/json-iterator/go"
	"github.com/syndtr/goleveldb/leveldb/errors"
)

func Load(store *leveldb.LevelStorage, key string, v interface{}) (err error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var data []byte
	data, err = store.Get(key)
	switch err {
	case errors.ErrNotFound:
		return nil
	case nil:
		return json.Unmarshal(data, v)
	default:
		return
	}
}

