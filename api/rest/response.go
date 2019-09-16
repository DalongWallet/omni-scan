package rest

import (
	"github.com/DalongWallet/omni-scan/api/rest/response"
	"github.com/gin-gonic/gin"
)

func RespJson(c *gin.Context, code int, data interface{}) {
	c.JSON(response.HttpStatusCode(code), response.RespJsonObj{
		Code: 		code,
		Msg: 		response.StatusText(code),
		Data: 		data,
	})
}
