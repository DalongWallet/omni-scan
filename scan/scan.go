package scan

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"io"
	"omni-scan/rpc"
	"os"
	"strconv"
	"time"
	"omni-scan/storage/leveldb"
	"github.com/mitchellh/cli"
)


func New() *cmd {
	return &cmd{}
}

type cmd struct{}

func (c *cmd) Run(args []string) int {
	ScanData1()
	return cli.RunResultHelp
}

func (c *cmd) Synopsis() string {
	return ""
}

func (c *cmd) Help() string {
	return ""
}

func ScanData() {
	infoLogFile := mustOpenFile("./scan_info.log")
	defer infoLogFile.Close()
	infoLogger := newLogger(infoLogFile, logrus.InfoLevel)

	errLogFile := mustOpenFile("./scan_err.log")
	defer errLogFile.Close()
	errLogger := newLogger(errLogFile, logrus.ErrorLevel)

	db, err := leveldb.Open("./omni_db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var hasScannedBlockHeight int64
	hasScannedBlockIndex, err := db.Get("hasScannedBlockHeight")
	switch err {
	case errors.ErrNotFound:
		hasScannedBlockHeight = 250000
	case nil:
		if hasScannedBlockHeight, err = strconv.ParseInt(string(hasScannedBlockIndex), 10, 64); err != nil {
			panic(err)
		}
	default:
		panic(err)
	}

	client := rpc.DefaultOmniClient

	var increment int64 = 1000
	startScanBlockHeight, endScanBlockHeight := hasScannedBlockHeight, hasScannedBlockHeight + increment
OUT:
	for {
		latestBlock, err := client.GetLatestBlockInfo()
		if err != nil {
			errLogger.Error(fmt.Sprintf("%+v \n\n", err))
			time.Sleep(1)
			continue
		}

		if startScanBlockHeight > latestBlock.BlockHeight {
			continue
		}

		recordNums := 0
		start := time.Now()

		fmt.Println("scan:", startScanBlockHeight,"-", endScanBlockHeight)
		txIdList, err := client.ListBlocksTransactions(startScanBlockHeight, endScanBlockHeight)
		if err != nil {
			errLogger.Error(fmt.Sprintf("%+v \n\n", err))
			time.Sleep(1)
			continue
		}
		fmt.Println("got txIdList length:", len(txIdList))
		if len(txIdList) > 0 {
			batch := db.NewBatch()
			for _, txId := range txIdList {
				tx, err := client.GetTransaction(txId)
				if err != nil {
					errLogger.Error(fmt.Sprintf("%+v \n\n", err))
					time.Sleep(1)
					continue OUT
				}

				fmt.Println("got Tx:", tx.TxId)
				key1 := fmt.Sprintf("%s-%d-%s", tx.SendingAddress, tx.PropertyId, tx.TxId)
				key2 := fmt.Sprintf("%s-%d-%s", tx.ReferenceAddress, tx.PropertyId, tx.TxId)
				value, err := json.Marshal(tx)
				if err != nil {
					errLogger.Error(fmt.Sprintf("%+v \n\n", err))
					time.Sleep(1)
					continue OUT
				}
				batch.Set(key1, value).Set(key2, value)
				fmt.Println("set Tx Key finished")
				for _, addr := range []string{
					tx.SendingAddress,
					tx.ReferenceAddress,
				} {
					fmt.Println(addr)
					addrAllBalances, err := client.GetAllBalancesForAddress(addr)
					if err != nil {
						fmt.Println(err)
						errLogger.Error(fmt.Sprintf("%+v \n\n", err))
						time.Sleep(1)
						continue OUT
					}
					fmt.Println("got Balance:",addr)
					for _, one := range addrAllBalances {
						key1 = fmt.Sprintf("%s-%d", addr, one.PropertyId)
						if value, err = json.Marshal(one); err != nil {
							errLogger.Error(fmt.Sprintf("%+v \n\n", err))
							time.Sleep(1)
							continue OUT
						}
						batch.Set(key1, value)
					}
				}
			}
			recordNums = batch.Len()

			if err = batch.Commit(); err != nil {
				errLogger.Error(fmt.Sprintf("%+v \n\n", err))
				time.Sleep(1)
				continue
			}
			fmt.Println("Commmit Data")

			if err = db.Set("hasScannedBlockHeight", []byte(strconv.FormatInt(endScanBlockHeight, 10))); err != nil {
				errLogger.Error(fmt.Sprintf("%+v \n\n", err))
				time.Sleep(1)
				continue
			}
		}

		infoLogger.Info(fmt.Sprintf("hasScannedBlockHeight: %d, recordNums: %d, use: %s", endScanBlockHeight, recordNums, time.Since(start).String()))

		if latestBlock.BlockHeight < endScanBlockHeight + increment  {
			increment = latestBlock.BlockHeight - endScanBlockHeight
		}

		startScanBlockHeight, endScanBlockHeight  = endScanBlockHeight + 1, endScanBlockHeight + increment
	}
}


func newLogger(writer io.Writer, level logrus.Level) *logrus.Logger  {
	logger := logrus.New()
	logger.SetOutput(writer)
	logger.SetLevel(level)
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		TimestampFormat: "2006-01-02 03:04:05",
		FullTimestamp:true,
	})
	return logger
}

func mustOpenFile(path string) *os.File {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	return file
}

func ScanData1() {
	startScanBlockHeight, endScanBlockHeight := int64(256000), int64(257000)
	client := rpc.DefaultOmniClient
	txIdList, err := client.ListBlocksTransactions(startScanBlockHeight, endScanBlockHeight)
	if err != nil {
		panic(err)
	}
	if len(txIdList) > 0 {
		for _, txId := range txIdList {
			tx, err := client.GetTransaction(txId)
			if err != nil {
			}
			fmt.Printf("Tx: %+v \n", tx)
			key1 := fmt.Sprintf("%s-%d-%s", tx.SendingAddress, tx.PropertyId, tx.TxId)
			key2 := fmt.Sprintf("%s-%d-%s", tx.ReferenceAddress, tx.PropertyId, tx.TxId)
			fmt.Printf("Transaction sending Key[%s]: %s \n  reference Key[%s]: %s", tx.SendingAddress, key1, tx.ReferenceAddress, key2)
			for _, addr := range []string{
				tx.SendingAddress,
				tx.ReferenceAddress,
			} {
				addrAllBalances, err := client.GetAllBalancesForAddress(addr)
				if err != nil {
				panic(err)
				}
				for _, one := range addrAllBalances {
					key1 = fmt.Sprintf("%s-%d", addr, one.PropertyId)
					fmt.Printf("Balance: %+v \n Balance key: %s \n", one, key1)
				}
			}
			}
		}
}