package main

import (
	"omni-scan/storage/leveldb"
)

func main() {
	db, err := leveldb.Open("./testdb")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
