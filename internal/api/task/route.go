package task

import "github.com/gin-gonic/gin"

func Route(ctx *gin.RouterGroup) {
	route := ctx.Group("/task")
	route.GET("/list")
}
