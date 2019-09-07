package rest

import (
	"context"
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Server struct {
	httpServer  	*http.Server
}

func NewHttpServer(port int) *Server {
	gin.SetMode(gin.ReleaseMode)
	engin := gin.Default()

	addr := fmt.Sprintf("127.0.0.1:%d", port)
	httpServer := &http.Server{
		Addr: 		addr,
		Handler: 	engin,
	}
	server := &Server{
		httpServer: 		httpServer,
	}
	server.initRouter(engin)

	return server
}

func (s *Server) Run() {
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

