package comm

import "github.com/gin-gonic/gin"

func Route(rg *gin.RouterGroup) {
	group := rg.Group("comm")
	// 登录
	group.POST("/login", login)
	// 获取个人信息
	group.GET("/me", me)
	// 上传接口
	group.POST("/upload", upload)
}
