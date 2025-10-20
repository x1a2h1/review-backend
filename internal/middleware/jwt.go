package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func AuthAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("当前用户权限：", "")
	}

}
