package leveldb

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/util"
)

type levelStorage struct {
	DB *leveldb.DB
}

var LevelStoragePool = map[string]*levelStorage{}


func GetLevelDbStorage(path string, opt*opt.Options) (storage *levelStorage) {
	var db *leveldb.DB

	var err error
	storage = LevelStoragePool[path]

	// storage has been created
	if storage != nil {
		if _, err = storage.DB.GetSnapshot(); err == nil {
			// db has opened, return directly
			return
		}else {
			// reopen if db has been closed
			if err.Error() == "leveldb: closed" {
				if db, err = leveldb.OpenFile(path, opt); err != nil {
					panic(err)
				}
				storage = newLevelStorage(db)
				LevelStoragePool[path] = storage
				return
			}
			panic(err)
		}
	}

	// storage has not been created
	if db, err = leveldb.OpenFile(path, opt); err == nil {
		storage = newLevelStorage(db)
		LevelStoragePool[path] = storage
		return
	}else {
		// db has been locked, try to recover
		if err.Error() == "resource temporarily unavailable" {
			if db, err = leveldb.RecoverFile(path, opt); err != nil {
				panic(err)
			}
			storage = newLevelStorage(db)
			LevelStoragePool[path] = storage
			return
		}
		panic(err)
	}
}

func newLevelStorage(db *leveldb.DB) *levelStorage {
	return &levelStorage{
		DB:db,
	}
}

func (s *levelStorage) Close() error {
	return s.DB.Close()
}

func (s *levelStorage) Get(key string) ([]byte, error)  {
	return s.DB.Get([]byte(key), nil)
}

func (s *levelStorage) GetString(key string) (string, error) {
	value, err := s.DB.Get([]byte(key), nil)
	return string(value), err
}

func (s *levelStorage) Set(key string, value []byte) error {
	return s.DB.Put([]byte(key), value,nil)
}

func (s *levelStorage) SetString(key string, value string) error {
	return s.DB.Put([]byte(key), []byte(value), nil)
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

func (s *levelStorage) Range(start, end string)([][]byte, error) {
	iter := s.DB.NewIterator(&util.Range{Start:[]byte(start), Limit:[]byte(end)}, nil)
	defer iter.Release()
	return iterateData(iter)
}

func (s *levelStorage) GetWithPrefix(prefix string) ([][]byte, error) {
	iter := s.DB.NewIterator(util.BytesPrefix([]byte(prefix)), nil)
	defer iter.Release()
	return iterateData(iter)
}

func (s *levelStorage) NewBatch() *Batch {
	return &Batch{
		db:    s,
		batch: new(leveldb.Batch),
	}
}

type Batch struct {
	db *levelStorage
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