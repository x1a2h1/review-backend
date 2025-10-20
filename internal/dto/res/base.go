package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, value any) {
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": value,
		"msg":  "请求成功!",
	})
}

// Fail godoc
//
// @param ctx *gin.Context
// @param code int
// @param message string
// @return void
func Fail(ctx *gin.Context, code int, message string) {
	ctx.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": code,
		"data": nil,
		"msg":  message,
	})
}
