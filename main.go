package main

import (
	"encoding/json"
	"fmt"
	"github.com/judwhite/go-svc/svc"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"omni-scan/api/rest"
	"omni-scan/rpc"
	"omni-scan/storage/leveldb"
	"os"
	"strconv"
	"time"
)

type program struct {
}

func main() {
	if len(os.Args) > 1 {
		ScanData()
	}else {
		rest.NewHttpServer(80)
	}
}

func (program) Init(env svc.Environment) error {
	fmt.Println("init...")
	return nil
}

func ScanData() {
	logFile, err := os.OpenFile("omni_data.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	db, err := leveldb.Open("./omni_db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var lastScanBlockHeight int64
	lastScanBlockIndex, err := db.Get("lastScanBlockIndex")
	switch err {
	case errors.ErrNotFound:
		lastScanBlockHeight = 0
	case nil:
		if lastScanBlockHeight, err = strconv.ParseInt(string(lastScanBlockIndex), 10, 64); err != nil {
			panic(err)
		}
	default:
		panic(err)
	}

OUT:
	for {
		time.Sleep(1)
		latestBlock, err := rpc.GetLatestBlockInfo()
		if err != nil {
			fmt.Fprintf(logFile, "%+v \n", err)
			continue
		}

		if lastScanBlockHeight >= latestBlock.Height {
			continue
		}

		var recordNums int
		start := time.Now()

		batch := db.NewBatch()
		txHashList, err := rpc.ListBlockTransactions(lastScanBlockHeight)
		if err != nil {
			fmt.Fprintf(logFile, "%+v \n", err)
			continue
		}
		for _, txHash := range txHashList {
			tx, err := rpc.GetTransaction(txHash)
			if err != nil {
				fmt.Fprintf(logFile, "%+v \n", err)
				continue OUT
			}
			// 1. 存交易
			key1 := fmt.Sprintf("%s-%d-%s", tx.SendingAddress, tx.PropertyId, tx.TxId)
			key2 := fmt.Sprintf("%s-%d-%s", tx.ReferenceAddress, tx.PropertyId, tx.TxId)
			value, err := json.Marshal(tx)
			if err != nil {
				fmt.Fprintf(logFile, "%+v \n", err)
				continue OUT
			}
			batch.Set(key1, value).Set(key2, value)

			// 2. 查余额，存余额
			for _, addr := range []string{
				tx.SendingAddress,
				tx.ReferenceAddress,
			} {
				addrAllBalances, err := rpc.GetAllBalancesForAddress(addr)
				if err != nil {
					fmt.Fprintf(logFile, "%+v \n", err)
					continue OUT
				}
				for _, one := range addrAllBalances {
					key1 = fmt.Sprintf("%s-%d", addr, one.PropertyId)
					if value, err = json.Marshal(one); err != nil {
						fmt.Fprintf(logFile, "%+v \n", err)
						continue OUT
					}
					batch.Set(key1, value)
				}
			}
		}
		recordNums = batch.Len()

		if err = batch.Commit(); err != nil {
			fmt.Fprintf(logFile, "%+v \n", err)
			continue
		}

		if err = db.Set("lastScanBlockIndex", []byte(strconv.FormatInt(lastScanBlockHeight+1, 10))); err != nil {
			fmt.Fprintf(logFile, "%+v \n", err)
			continue
		}
		lastScanBlockHeight++
		fmt.Fprintf(logFile, "================== hasScanBlockHeight: %d, recordNums: %d, use: %d \n", lastScanBlockHeight, recordNums, time.Since(start))
	}
}
