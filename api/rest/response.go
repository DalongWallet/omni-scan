package rest

import "github.com/gin-gonic/gin"

const (
	OK = 0
	BadRequest = 40002
	InternalServerError = 50000
)

var statusText = map[int]string {
	OK: 		"Success",
	BadRequest: "Bad request params",
	InternalServerError: "Server internal error",
}

type RespJsonObj struct {
	Code 	int  	`json:"code"`
	Msg  	string 	`json:"msg"`
	Data  	interface{} `json:"data"`
}

func StatusText(code int) string {
	return statusText[code]
}

func HttpStatusCode(code int) int {
	return code / 100
}

func RespJson(c *gin.Context, code int, data interface{}) {
	c.JSON(HttpStatusCode(code), RespJsonObj{
		Code: 		code,
		Msg: 		StatusText(code),
		Data: 		data,
	})
}
