package scan

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/cli"
	"github.com/sirupsen/logrus"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"io"
	"github.com/DalongWallet/omni-scan/rpc"
	"github.com/DalongWallet/omni-scan/storage/leveldb"
	"os"
	"strconv"
	"time"
)

func New() *cmd {
	return &cmd{}
}

type cmd struct{}

func (c *cmd) Run(args []string) int {
	ScanData()
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

	db := leveldb.GetLevelStorage("./omni_db", nil)
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
	startScanBlockHeight, endScanBlockHeight := hasScannedBlockHeight, hasScannedBlockHeight+increment
OUT:
	for {
		latestBlock, err := client.GetLatestBlockInfo()
		if err != nil {
			if err.Error() != "Work queue depth exceeded" {
				errLogger.Error(fmt.Sprintf("%+v \n\n", err))
			}
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
			if err.Error() != "Work queue depth exceeded" {
				errLogger.Error(fmt.Sprintf("%+v \n\n", err))
			}
			time.Sleep(1)
			continue
		}
		if len(txIdList) > 0 {
			batch := db.NewBatch()
			for _, txId := range txIdList {
				tx, err := client.GetTransaction(txId)
				if err != nil {
					if err.Error() != "Work queue depth exceeded" {
						errLogger.Error(fmt.Sprintf("%+v \n\n", err))
					}
					time.Sleep(1)
					continue OUT
				}

				key1 := fmt.Sprintf("tx-%s-%d-%s", tx.SendingAddress, tx.PropertyId, tx.TxId)
				key2 := fmt.Sprintf("tx-%s-%d-%s", tx.ReferenceAddress, tx.PropertyId, tx.TxId)
				value, err := json.Marshal(tx)
				if err != nil {
					errLogger.Error(fmt.Sprintf("%+v \n\n", err))
					time.Sleep(1)
					continue OUT
				}
				batch.Set(key1, value).Set(key2, value)

				for _, addr := range []string{
					tx.SendingAddress,
					tx.ReferenceAddress,
				} {
					addrAllBalances, err := client.GetAllBalancesForAddress(addr)
					if err != nil {
						if err.Error() != "Work queue depth exceeded" {
							errLogger.Error(fmt.Sprintf("%+v \n\n", err))
						}
						time.Sleep(1)
						continue OUT
					}

					for _, one := range addrAllBalances {
						key1 = fmt.Sprintf("balance-%s-%d", addr, one.PropertyId)
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

			if err = db.Set("hasScannedBlockHeight", []byte(strconv.FormatInt(endScanBlockHeight, 10))); err != nil {
				errLogger.Error(fmt.Sprintf("%+v \n\n", err))
				time.Sleep(1)
				continue
			}
		}

		infoLogger.Info(fmt.Sprintf("hasScannedBlockHeight: %d, recordNums: %d, use: %s", endScanBlockHeight, recordNums, time.Since(start).String()))

		if latestBlock.BlockHeight < endScanBlockHeight+increment {
			increment = latestBlock.BlockHeight - endScanBlockHeight
		}

		startScanBlockHeight, endScanBlockHeight = endScanBlockHeight+1, endScanBlockHeight+increment
	}
}

func newLogger(writer io.Writer, level logrus.Level) *logrus.Logger {
	logger := logrus.New()
	logger.SetOutput(writer)
	logger.SetLevel(level)
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		TimestampFormat: "2006-01-02 03:04:05",
		FullTimestamp:   true,
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
