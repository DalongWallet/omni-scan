package jwt

import (
	"github.com/DalongWallet/omni-scan/api/rest/response"
	"github.com/DalongWallet/omni-scan/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		var code int
		var data interface{}

		if context.Request.URL.Path == "/api/v1/token" {
			context.Next()
			return
		}

		code = response.OK
		token := context.Query("token")
		if token == "" {
			code = response.BadRequest
		}else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = response.ErrorAuthCheckTokenInvalid
			}else if time.Now().Unix() > claims.ExpiresAt {
				code = response.ErrorAuthCheckTokenTimetout
			}
		}

		if code != response.OK {
			context.JSON(http.StatusUnauthorized, gin.H{
				"code":  	code,
				"msg": 		"error JWT",
				"data": 	data,
			})

			context.Abort()
			return
		}

		context.Next()
	}
}