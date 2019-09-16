package leveldb

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
)

type LevelStorage struct {
	DB *leveldb.DB
}

var LevelStoragePool = map[string]*LevelStorage{}


func GetLevelDbStorage(path string, opt *opt.Options) *LevelStorage {
	if db, err := leveldb.OpenFile(path, opt); err == nil {
		return newLevelStorage(db)
	}else {
		// db has been locked, try to recover
		if err.Error() == "resource temporarily unavailable" {
			if db, err = leveldb.RecoverFile(path, opt); err != nil {
				panic(err)
			}
			return newLevelStorage(db)
		}
		panic(err)
	}
}


func newLevelStorage(db *leveldb.DB) *LevelStorage {
	return &LevelStorage{
		DB:db,
	}
}

func (s *LevelStorage) Close() error {
	return s.DB.Close()
}

func (s *LevelStorage) Get(key string) ([]byte, error)  {
	return s.DB.Get([]byte(key), nil)
}

func (s *LevelStorage) GetString(key string) (string, error) {
	value, err := s.DB.Get([]byte(key), nil)
	return string(value), err
}

func (s *LevelStorage) Set(key string, value []byte) error {
	return s.DB.Put([]byte(key), value,nil)
}

func (s *LevelStorage) SetString(key string, value string) error {
	return s.DB.Put([]byte(key), []byte(value), nil)
}

func (s *LevelStorage) Delete(key string) error {
	return s.DB.Delete([]byte(key), nil)
}

func iterateData(iter iterator.Iterator)  (data [][]byte, err error) {
	for iter.Next() {
		value := make([]byte, len(iter.Value()))
		copy(value, iter.Value())
		data = append(data,value)
	}
	err = iter.Error()
	return
}

// range [start, end)
func (s *LevelStorage) Range(start, end string)([][]byte, error) {
	iter := s.DB.NewIterator(&util.Range{Start:[]byte(start), Limit:[]byte(end)}, nil)
	defer iter.Release()
	return iterateData(iter)
}

func (s *LevelStorage) GetWithPrefix(prefix string) ([][]byte, error) {
	iter := s.DB.NewIterator(util.BytesPrefix([]byte(prefix)), nil)
	defer iter.Release()
	return iterateData(iter)
}

func (s *LevelStorage) NewBatch() *Batch {
	return &Batch{
		db:    s,
		batch: new(leveldb.Batch),
	}
}

type Batch struct {
	db *LevelStorage
	batch *leveldb.Batch
}

func (b *Batch) Set(key string, value []byte) *Batch {
	b.batch.Put([]byte(key), value)
	return b
}

func (b *Batch) Delete(key string) *Batch {
	b.batch.Delete([]byte(key))
	return b
}

func (b *Batch) Len() int {
	return b.batch.Len()
}

func (b *Batch) Commit() error {
	return b.db.DB.Write(b.batch, nil)
}