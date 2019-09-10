package rest

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"omni-scan/logic"
	"omni-scan/omnicore"
)

type Server struct {
	httpServer 	*http.Server
	mgr        	*logic.OmniMgr
	omniCli 	*omnicore.Client
}

func NewHttpServer(port int) *Server {
	gin.SetMode(gin.ReleaseMode)
	engin := gin.Default()

	addr := fmt.Sprintf("127.0.0.1:%d", port)
	httpServer := &http.Server{
		Addr:    addr,
		Handler: engin,
	}
	server := &Server{
		httpServer: httpServer,
	}
	server.initRouter(engin)

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
}
