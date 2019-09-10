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
	log "github.com/sirupsen/logrus"
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
	infoLogger := log.New()
	infoLogFile, err := os.OpenFile("scan_info.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer infoLogFile.Close()
	infoLogger.SetOutput(infoLogFile)
	infoLogger.SetLevel(log.InfoLevel)

	errLogger := log.New()
	errLogFile, err := os.OpenFile("scan_err.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer errLogFile.Close()
	errLogger.SetOutput(errLogFile)
	errLogger.SetLevel(log.ErrorLevel)

	formatter := &log.TextFormatter{
		ForceColors:     true,
		TimestampFormat: "2006-01-02 03:04:05",
		FullTimestamp:true,
	}
	infoLogger.SetFormatter(formatter)
	errLogger.SetFormatter(formatter)


	db, err := leveldb.Open("./omni_db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var lastScanBlockHeight int64
	lastScanBlockIndex, err := db.Get("hasScanedBlockHeight")
	switch err {
	case errors.ErrNotFound:
		lastScanBlockHeight = 250000
	case nil:
		if lastScanBlockHeight, err = strconv.ParseInt(string(lastScanBlockIndex), 10, 64); err != nil {
			panic(err)
		}
	default:
		panic(err)
	}



	client := rpc.DefaultOmniClient

	var increment int64 = 1000
OUT:
	for {
		time.Sleep(1)
		latestBlock, err := client.GetLatestBlockInfo()
		if err != nil {
			errLogger.Error(fmt.Sprintf("%+v \n\n", err))
			continue
		}

		if lastScanBlockHeight > latestBlock.BlockHeight {
			continue
		}

		batch := db.NewBatch()
		recordNums := 0	
		start := time.Now()
		startScanBlockHeight, endScanBlockHeight := lastScanBlockHeight, lastScanBlockHeight + increment

		txIdList, err := client.ListBlocksTransactions(startScanBlockHeight, endScanBlockHeight)
		if err != nil {
			errLogger.Error(fmt.Sprintf("%+v \n\n", err))
			continue
		}

		if len(txIdList) > 0 {
			for _, txId := range txIdList {
				tx, err := client.GetTransaction(txId)
				if err != nil {
					errLogger.Error(fmt.Sprintf("%+v \n\n", err))
					continue OUT
				}

				key1 := fmt.Sprintf("%s-%d-%s", tx.SendingAddress, tx.PropertyId, tx.TxId)
				key2 := fmt.Sprintf("%s-%d-%s", tx.ReferenceAddress, tx.PropertyId, tx.TxId)
				value, err := json.Marshal(tx)
				if err != nil {
					errLogger.Error(fmt.Sprintf("%+v \n\n", err))
					continue OUT
				}
				batch.Set(key1, value).Set(key2, value)

				for _, addr := range []string{
					tx.SendingAddress,
					tx.ReferenceAddress,
				} {
					addrAllBalances, err := client.GetAllBalancesForAddress(addr)
					if err != nil {
						errLogger.Error(fmt.Sprintf("%+v \n\n", err))
						continue OUT
					}
					for _, one := range addrAllBalances {
						key1 = fmt.Sprintf("%s-%d", addr, one.PropertyId)
						if value, err = json.Marshal(one); err != nil {
							errLogger.Error(fmt.Sprintf("%+v \n\n", err))
							continue OUT
						}
						batch.Set(key1, value)
					}
				}
			}
			recordNums = batch.Len()

			if err = batch.Commit(); err != nil {
				errLogger.Error(fmt.Sprintf("%+v \n\n", err))
				continue
			}

			if err = db.Set("hasScanedBlockHeight", []byte(strconv.FormatInt(endScanBlockHeight, 10))); err != nil {
				errLogger.Error(fmt.Sprintf("%+v \n\n", err))
				continue
			}
		}

		infoLogger.Info(fmt.Sprintf("hasScanedBlockHeight: %d, recordNums: %d, use: %s \n", endScanBlockHeight, recordNums, time.Since(start).String()))

		if endScanBlockHeight + increment - latestBlock.BlockHeight > 0 {
			increment = latestBlock.BlockHeight - endScanBlockHeight
		}

		lastScanBlockHeight = endScanBlockHeight + 1
	}
}
