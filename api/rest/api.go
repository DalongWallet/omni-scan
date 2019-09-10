package rest

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

const OmniPropertyUSDT = 31

func (s *Server) GetBlocksTxHashList(c *gin.Context) {
	startStr := c.Query("start")
	endStr := c.Query("end")
	start, _ := strconv.Atoi(startStr)
	end, _ := strconv.Atoi(endStr)

	hashList, err := s.omniCli.RpcClient.ListBlocksTransactions(int64(start), int64(end))
	if err != nil {
		RespJson(c, InternalServerError, err.Error())
		return
	}

	RespJson(c, OK, hashList)
}

func (s *Server) GetTransactionById(c *gin.Context) {
	txId := c.Query("tx")
	if txId == "" {
		RespJson(c, BadRequest, "require tx")
		return
	}

	tx, err := s.omniCli.RpcClient.GetTransaction(txId)
	if err != nil {
		RespJson(c, InternalServerError, err.Error())
		return
	}

	RespJson(c, OK, tx)
}

func (s *Server) GetAddressBalance(c *gin.Context) {
	addr := c.Query("address")
	if addr == "" {
		RespJson(c, BadRequest, "require address")
		return
	}

	balance, err := s.mgr.GetAddressBalance(addr, OmniPropertyUSDT)
	if err != nil {
		RespJson(c, InternalServerError, err.Error())
		return
	}

	RespJson(c, OK, balance)
}

func (s *Server) GetConfirmedAddressTransactions(c *gin.Context) {
	addr := c.Query("address")
	if addr == "" {
		RespJson(c, BadRequest, "require address")
		return
	}

	limitStr := c.Query("limit")
	if limitStr == "" {
		limitStr = "10"
	}
	offsetStr := c.Query("offset")

	limit, _ := strconv.Atoi(limitStr)
	offset, _ := strconv.Atoi(offsetStr)

	txs, err := s.mgr.GetAddressConfirmedTxs(addr, uint(limit), uint(offset))
	if err != nil {
		RespJson(c, InternalServerError, err.Error())
		return
	}

	RespJson(c, OK, txs)
}

func (s *Server) PushRawTransaction(c *gin.Context) {
	hex := c.Query("txHex")
	hex = strings.TrimSpace(hex)
	hex = strings.TrimPrefix(hex, "0x")
	if hex == "" {
		RespJson(c, BadRequest, "txHex invalid")
		return
	}

	addr := c.Query("addr")
	if addr == "" {
		RespJson(c, BadRequest, "require sender address")
		return
	}

	txHash, err := s.omniCli.RpcClient.SendRawTransaction(addr, hex)
	if err != nil {
		RespJson(c, BadRequest, err.Error())
		return
	}

	tx, err := s.omniCli.RpcClient.DecodeTransaction(txHash)
	if err != nil {
		RespJson(c, BadRequest, err.Error())
		return
	}

	RespJson(c, OK, tx)
}
