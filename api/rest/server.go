package rest

import (
	"context"
	"fmt"
	"github.com/DalongWallet/omni-scan/api/rest/middleware/jwt"
	"github.com/DalongWallet/omni-scan/logic"
	"github.com/DalongWallet/omni-scan/omnicore"
	"github.com/DalongWallet/omni-scan/rpc"
	"github.com/DalongWallet/omni-scan/storage/leveldb"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Server struct {
	storage    *leveldb.LevelStorage
	httpServer *http.Server
	mgr        *logic.OmniMgr
	omniCli    *omnicore.Client
	worker     *omnicore.Worker
}

func NewServer(port int) *Server {
	gin.SetMode(gin.ReleaseMode)

	engine := gin.Default()
	engine.Use(jwt.JWT())

	addr := fmt.Sprintf("127.0.0.1:%d", port)
	httpServer := &http.Server{
		Addr:    addr,
		Handler: engine,
	}

	storage := leveldb.GetLevelDbStorage("./omni_db", nil)
	omniCli := &omnicore.Client{
		RpcClient: rpc.DefaultOmniClient,
	}
	worker := omnicore.NewWorker(storage, omniCli.RpcClient)
	mgr, err := logic.NewOmniMgr(storage, 64)
	if err != nil {
		panic(err)
	}

	server := &Server{
		storage:    storage,
		httpServer: httpServer,
		omniCli:    omniCli,
		worker:     worker,
		mgr:        mgr,
	}
	server.initRouter(engine)

	return server
}

func (s *Server) Run() int {
	var wg sync.WaitGroup
	startRestService := func() {
		wg.Add(1)
		defer wg.Done()
		logrus.Println("Start omni-scan REST api service")
		if err := s.httpServer.ListenAndServe(); err != nil {
			logrus.Errorf("RestServer.Run %s", err)
		}
		logrus.Println("rest service shutdown")
	}
	startScanService := func() {
		wg.Add(1)
		defer wg.Done()
		logrus.Println("Start omni-scan Scan service")
		s.worker.Run()
		logrus.Println("scan service shutdown")
	}

	waitToStop := func() {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
		<-signalChan
	}

	stopAllServices := func() {
		s.worker.Stop()
		s.httpServer.Shutdown(context.Background())
		s.storage.Close()
		wg.Wait()
	}

	go startRestService()
	go startScanService()
	waitToStop()
	stopAllServices()

	return 1
}

func (s *Server) initRouter(r gin.IRouter) {
	r.GET("/api/v1/token", s.GenerateToken)
	r.GET("/api/v1/usdtTxs", s.GetConfirmedAddressUsdtTransactions)
	r.GET("/api/v1/propertyTxs", s.GetConfirmedAddressPropertyTransactions)
	r.GET("/api/v1/txHashListByBlocks", s.GetBlocksTxHashList)
	r.GET("/api/v1/tx", s.GetTransactionById)
	r.GET("/api/v1/usdtBalance", s.GetAddressUsdtBalance)
	r.GET("/api/v1/propertyBalance", s.GetAddressPropertyBalance)
	r.POST("/api/v1/sendRawTx", s.SendRawTransaction)
	r.POST("/api/v1/decodeRawTx", s.DecodeRawTransaction)
	r.GET("/api/v1/usdtBalanceRPC", s.GetAddressUsdtBalanceByRpc)
	r.GET("/api/v1/latestBlockInfo", s.GetLatestBlockInfo)
	r.GET("/api/v1/latestBlockInfoRPC", s.GetLatestBlockInfoByRpc)
}
