package rest

import (
	"errors"
	. "github.com/DalongWallet/omni-scan/api/rest/response"
	"github.com/DalongWallet/omni-scan/models"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

const OmniPropertyUSDT = 31

func (s *Server) GetBlocksTxHashList(c *gin.Context) {
	startStr := c.Query("start")
	endStr := c.Query("end")

	if !isUintStr(startStr) || !isUintStr(endStr) {
		RespJson(c, BadRequest, errors.New("start or end must >= 0"))
	}

	start, _ := strconv.Atoi(startStr)
	end, _ := strconv.Atoi(endStr)

	if start > end {
		RespJson(c, BadRequest, errors.New("end must be greater than start"))
	}

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

	var tx models.Transaction
	if err := tx.Load(s.storage, txId); err != nil {
		RespJson(c, InternalServerError, err.Error())
		return
	}

	RespJson(c, OK, tx)
}

func (s *Server) GetAddressPropertyBalance(c *gin.Context) {
	addr := c.Query("address")
	if addr == "" {
		RespJson(c, BadRequest, "require address")
		return
	}
	propertyIdStr := c.Query("propertyId")
	if propertyIdStr == "" {
		RespJson(c, BadRequest, "require propertyId")
		return
	}
	if !isUintStr(propertyIdStr) {
		RespJson(c, BadRequest, "propertyId must >= 0")
		return
	}
	propertyId, _ := strconv.Atoi(propertyIdStr)

	var balance models.PropertyTokenBalance
	if err := balance.Load(s.storage, addr, propertyId); err != nil {
		RespJson(c, InternalServerError, err.Error())
		return
	}

	RespJson(c, OK, balance)
}

func (s *Server) GetAddressUsdtBalance(c *gin.Context) {
	addr := c.Query("address")
	if addr == "" {
		RespJson(c, BadRequest, "require address")
		return
	}

	var balance models.PropertyTokenBalance
	if err := balance.Load(s.storage, addr, OmniPropertyUSDT); err != nil {
		RespJson(c, InternalServerError, err.Error())
		return
	}

	RespJson(c, OK, balance)
}

func (s *Server) GetConfirmedAddressUsdtTransactions(c *gin.Context) {
	addr := c.Query("address")
	if addr == "" {
		RespJson(c, BadRequest, "require address")
		return
	}

	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	if !isUintStr(limitStr) {
		RespJson(c, BadRequest, "limit must >= 0")
		return
	}
	if !isUintStr(offsetStr) {
		RespJson(c, BadRequest, "offset must >= 0")
		return
	}

	limit, _ := strconv.Atoi(limitStr)
	offset, _ := strconv.Atoi(offsetStr)

	txs, err := s.mgr.GetAddressConfirmedTxs(addr, OmniPropertyUSDT, uint(limit), uint(offset))
	if err != nil {
		RespJson(c, InternalServerError, err.Error())
		return
	}

	RespJson(c, OK, txs)
}

func (s *Server) GetConfirmedAddressPropertyTransactions(c *gin.Context) {
	addr := c.Query("address")
	if addr == "" {
		RespJson(c, BadRequest, "require address")
		return
	}

	propertyIdStr := c.Query("propertyId")
	if propertyIdStr == "" {
		RespJson(c, BadRequest, "require propertyId")
		return
	}
	if !isUintStr(propertyIdStr) {
		RespJson(c, BadRequest, "propertyId must >= 0")
		return
	}
	propertyId, _ := strconv.Atoi(propertyIdStr)

	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	if !isUintStr(limitStr) {
		RespJson(c, BadRequest, "limit must >= 0")
		return
	}
	if !isUintStr(offsetStr) {
		RespJson(c, BadRequest, "offset must >= 0")
		return
	}

	limit, _ := strconv.Atoi(limitStr)
	offset, _ := strconv.Atoi(offsetStr)

	txs, err := s.mgr.GetAddressConfirmedTxs(addr,propertyId, uint(limit), uint(offset))
	if err != nil {
		RespJson(c, InternalServerError, err.Error())
		return
	}

	RespJson(c, OK, txs)
}

func (s *Server) SendRawTransaction(c *gin.Context) {
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

	RespJson(c, OK, txHash)
}

func (s *Server) DecodeRawTransaction(c *gin.Context) {
	txHex := c.Query("txHex")
	txHex = strings.TrimSpace(txHex)
	txHex = strings.TrimPrefix(txHex, "0x")
	if txHex == "" {
		RespJson(c, BadRequest, "txHex invalid")
		return
	}

	tx, err := s.omniCli.RpcClient.DecodeTransaction(txHex)
	if err != nil {
		RespJson(c, BadRequest, err.Error())
		return
	}

	RespJson(c, OK, tx)
}
