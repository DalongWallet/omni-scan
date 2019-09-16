package rest

import (
	"context"
	"fmt"
	"github.com/DalongWallet/omni-scan/api/rest/middleware/jwt"
	"github.com/DalongWallet/omni-scan/logic"
	"github.com/DalongWallet/omni-scan/omnicore"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Server struct {
	httpServer *http.Server
	mgr        *logic.OmniMgr
	omniCli    *omnicore.Client
}

func NewHttpServer(port int) *Server {
	gin.SetMode(gin.ReleaseMode)

	engine := gin.Default()
	engine.Use(jwt.JWT())

	addr := fmt.Sprintf("127.0.0.1:%d", port)
	httpServer := &http.Server{
		Addr:    addr,
		Handler: engine,
	}
	server := &Server{
		httpServer: httpServer,
	}
	server.initRouter(engine)

	return server
}

func (s *Server) Run() {
	logrus.Print("Start omni-scan REST api service")

	err := s.httpServer.ListenAndServe()
	if err != nil {
		logrus.Error("RestServer.Run %s", err)
	}
}

func (s *Server) Stop() {
	s.httpServer.Shutdown(context.Background())
}

func (s *Server) initRouter(r gin.IRouter) {
	r.GET("/api/v1/txs", s.GetConfirmedAddressTransactions)
	r.GET("/api/v1/txHashListByBlocks", s.GetBlocksTxHashList)
	r.GET("/api/v1/tx", s.GetTransactionById)
	r.GET("/api/v1/balance", s.GetAddressBalance)
	r.POST("/api/v1/sendRawTx", s.SendRawTransaction)

}
