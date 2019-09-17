package utils

import (
	"github.com/DalongWallet/omni-scan/storage/leveldb"
	jsoniter "github.com/json-iterator/go"
	"github.com/syndtr/goleveldb/leveldb/errors"
)

func Load(store *leveldb.LevelStorage, key string, v interface{}) (err error) {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var data []byte
	if data, err = store.Get(key); err != nil {
		if IsErrorNotFound(err) {
			err = nil
		}
		return
	}

	return json.Unmarshal(data, v)
}

func IsErrorNotFound(err error) bool {
	switch err {
	case errors.ErrNotFound:
		return true
	default:
		return false
	}
}