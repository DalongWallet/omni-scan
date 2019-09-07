package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func (s *Server) GetConfirmedAddressTransactions(c *gin.Context) {
	addr := c.Query("address")
	if addr == "" {
		RespJson(c, BadRequest, nil)
		return
	}

	limitStr := c.Query("limit")
	if limitStr == "" {
		limitStr = "10"
	}
	offsetStr := c.Query("offset")

	limit, _ := strconv.Atoi(limitStr)
	offset, _ := strconv.Atoi(offsetStr)

	fmt.Println("debug joy", limit, offset)
	RespJson(c, OK, "")
}
