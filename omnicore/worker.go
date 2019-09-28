package omnicore

import (
	"context"
	"fmt"
	"github.com/DalongWallet/omni-scan/models"
	"github.com/DalongWallet/omni-scan/rpc"
	"github.com/DalongWallet/omni-scan/storage/leveldb"
	"github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"io"
	"os"
	"strconv"
	"time"
)
var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Worker struct {
	storage *leveldb.LevelStorage
	rpcClient *rpc.OmniClient
	ctx context.Context
	stop context.CancelFunc
}

func NewWorker(storage *leveldb.LevelStorage, rpcClient *rpc.OmniClient) *Worker {
	ctx, cancel := context.WithCancel(context.Background())
	return &Worker{
		storage:storage,
		rpcClient:rpcClient,
		ctx:ctx,
		stop:cancel,
	}
}


func (w *Worker) Run() {
	infoLogFile := mustOpenFile("./scan_info.log")
	defer infoLogFile.Close()
	infoLogger := newLogger(infoLogFile, logrus.InfoLevel)

	errLogFile := mustOpenFile("./scan_err.log")
	defer errLogFile.Close()
	errLogger := newLogger(errLogFile, logrus.ErrorLevel)

	db := w.storage

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

	var increment int64 = 1000
	startScanBlockHeight, endScanBlockHeight := hasScannedBlockHeight+1, hasScannedBlockHeight+increment

	for {
		if w.isDone() {
			return
		}

		// TODO: save LatestBlockInfo
		latestBlock, err := w.rpcClient.GetLatestBlockInfo()
		if err != nil {
			if err.Error() != "Work queue depth exceeded" {
				errLogger.Error(fmt.Sprintf("GetInfo Failed, %+v \n\n", err))
			}
			time.Sleep(1 * time.Second)
			continue
		}
		infoLogger.Info("latestBlockHeight:", latestBlock.BlockHeight)
		latestBlockData, err := json.Marshal(latestBlock)
		if err != nil {
			errLogger.Error(fmt.Sprintf("GetInfo Failed, %+v \n\n", err))
			time.Sleep(1 * time.Second)
			continue
		}
		if err = db.Set(models.LatestBlockInfoKey(), latestBlockData); err != nil {
			errLogger.Error(fmt.Sprintf("GetInfo Failed, %+v \n\n", err))
			time.Sleep(1 * time.Second)
			continue
		}

		if startScanBlockHeight > latestBlock.BlockHeight {
			time.Sleep(5 * time.Second)
			continue
		}

		recordNums := 0
		start := time.Now()

		infoLogger.Info("scan:", startScanBlockHeight,"-", endScanBlockHeight)
		txIdList, err := w.rpcClient.ListBlocksTransactions(startScanBlockHeight, endScanBlockHeight)
		if err != nil {
			if err.Error() != "Work queue depth exceeded" {
				errLogger.Error(fmt.Sprintf("ListBlocksTransactions [%d,%d] Failed,%+v \n\n", startScanBlockHeight, endScanBlockHeight, err))
			}
			time.Sleep(1 * time.Second)
			continue
		}
		infoLogger.Info("tx count:", len(txIdList))
		if len(txIdList) > 0 {
			batch := db.NewBatch()
			txQueue := NewTaskQueue(txIdList)
			for !txQueue.AllFinished() {
				if w.isDone() {
					return
				}
				txId := txQueue.GetTask()
				tx, err := w.rpcClient.GetTransaction(txId)
				if err != nil {
					if err.Error() != "Work queue depth exceeded" {
						errLogger.Error(fmt.Sprintf("GetTransaction [%s] Failed, %+v \n\n", txId,err))
					}
					time.Sleep(1)
					continue
				}

				var addrs []string
				if tx.SendingAddress != "" {
					addrs = append(addrs, tx.SendingAddress)
				}
				if tx.ReferenceAddress != "" {
					addrs = append(addrs, tx.ReferenceAddress)
				}

				txBytes, err := json.Marshal(tx)
				if err != nil {
					errLogger.Error(fmt.Sprintf("Marshal Tx [ %+v ] Failed,%+v \n\n", tx, err))
					time.Sleep(1 * time.Second)
					continue
				}
				batch.Set(models.TxKey(txId), txBytes)

				addrQueue := NewTaskQueue(addrs)
				for !addrQueue.AllFinished() {
					if w.isDone() {
						return
					}
					addr := addrQueue.GetTask()

					key := models.AddrPropertyTxKey(addr, tx.PropertyId, tx.TxId)
					batch.Set(key, txBytes)

					addrAllBalances, err := w.rpcClient.GetAllBalancesForAddress(addr)
					if err != nil {
						if err.Error() != "Work queue depth exceeded" {
							errLogger.Error(fmt.Sprintf("Txid [%s], Get Address [%s] Balance Failed, %+v \n\n",tx.TxId, addr, err))
						}
						time.Sleep(1 * time.Second)
						continue
					}
					for _, one := range addrAllBalances {
						key = models.AddrPropertyBalanceKey(addr, one.PropertyId)
						balanceBytes, err := json.Marshal(one)
						if err != nil {
							errLogger.Error(fmt.Sprintf("Marshal Balance [ %+v ] Failed, %+v \n\n", one,err))
							time.Sleep(1 * time.Second)
							continue
						}
						batch.Set(key, balanceBytes)
					}
					addrQueue.MarkTaskDone()
				}

				txQueue.MarkTaskDone()
			}

			recordNums = batch.Len()

			if err = batch.Commit(); err != nil {
				errLogger.Error(fmt.Sprintf("%+v \n\n", err))
				time.Sleep(1 * time.Second)
				continue
			}

			if err = db.Set("hasScannedBlockHeight", []byte(strconv.FormatInt(endScanBlockHeight, 10))); err != nil {
				errLogger.Error(fmt.Sprintf("Cache HasScannedBlockHeight Failed, %+v \n\n", err))
				time.Sleep(1 * time.Second)
				continue
			}
		}

		infoLogger.Info(fmt.Sprintf("hasScannedBlockHeight: %d, recordNums: %d, use: %s", endScanBlockHeight, recordNums, time.Since(start).String()))

		if latestBlock.BlockHeight - endScanBlockHeight < 10 {
			increment = 1
		}

		if endScanBlockHeight + increment >= latestBlock.BlockHeight {
			increment = latestBlock.BlockHeight - endScanBlockHeight
		}

		startScanBlockHeight, endScanBlockHeight = endScanBlockHeight+1, endScanBlockHeight+increment
	}
}

func (w *Worker) Stop() {
	w.stop()
}

func (w *Worker) isDone() bool {
	select {
	case <-w.ctx.Done():
		return true
	default:
		return false
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
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return file
}



