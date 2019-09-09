package main

import (
	"fmt"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"omni-scan/storage/leveldb"
	"github.com/judwhite/go-svc/svc"
	"omni-scan/api/rest"
	"time"
	"omni-scan/rpc"
	"strconv"
)

type program struct{

}

func main() {
	rest.NewHttpServer(80)
}

func (program) Init(env svc.Environment) error {
	fmt.Println("init...")
	return nil
}

func SaveData() {
	db, err := leveldb.Open("")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var curBlockHeight int64
	if gatherBlock, err := db.Get("gatherBlock"); err != nil {
		if err == errors.ErrNotFound {
			curBlockHeight = 0
		}else {
			panic(err)
		}
	}else {
		curBlockHeight, err  = strconv.ParseInt(string(gatherBlock), 10, 64)
		if err != nil {
			panic(err)
		}
	}

	for {
		time.Sleep(1)
		latestBlock, err := rpc.GetLatestBlockInfo()
		if err != nil {
			continue
		}
		if curBlockHeight < latestBlock.Height {
			batch := db.NewBatch()
			batch.Set("test",[]byte{})

			txs, err := rpc.GetBlockTransactions(curBlockHeight)
			if err != nil {
				continue
			}
			for _, tx := range txs {
				fmt.Println(tx)
				// 交易是双向的，from 记录一条, to 记录一条
				//key := fmt.Sprintf("%s-%s-%s", tx.SendingAddress, tx.PropertyId, tx.TxId)
				//value := json.Mu
			}

			if err = batch.Commit(); err != nil {
				continue
			}
		}

	}
}