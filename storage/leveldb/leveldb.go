package leveldb

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	"github.com/syndtr/goleveldb/leveldb/util"
)

type Storage struct {
	DB *leveldb.DB
}

func Open(path string) (*Storage, error) {
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		return &Storage{}, err
	}
	return &Storage{DB: db}, nil
}

func (s *Storage) Close() error {
	return s.DB.Close()
}

func (s *Storage) Get(key string) ([]byte, error)  {
	return s.DB.Get([]byte(key), nil)
}

func (s *Storage) Set(key string, value []byte) error {
	return s.DB.Put([]byte(key), value,nil)
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

func (s *Storage) Range(start, end string)([][]byte, error) {
	iter := s.DB.NewIterator(&util.Range{Start:[]byte(start), Limit:[]byte(end)}, nil)
	defer iter.Release()
	return iterateData(iter)
}

func (s *Storage) GetWithPrefix(prefix string) ([][]byte, error) {
	iter := s.DB.NewIterator(util.BytesPrefix([]byte(prefix)), nil)
	defer iter.Release()
	return iterateData(iter)
}

func (s *Storage) NewBatch() *Batch {
	return &Batch{
		db:    s,
		batch: new(leveldb.Batch),
	}
}

type Batch struct {
	db *Storage
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
	b.batch.Reset()
	return b.batch.Len()
}

func (b *Batch) Commit() error {
	return b.db.DB.Write(b.batch, nil)
}