package project

import "github.com/gin-gonic/gin"

func Route(ctx *gin.RouterGroup) {
	route := ctx.Group("/project")
	route.GET("/list")   //返回用户创建于加入的项目列表 通过type来区分创建或者加入
	route.GET("/:id")    //获取单个项目详情
	route.PUT(":/id")    //修改项目信息
	route.DELETE("/:id") //删除项目
	// 项目邀请相关
	route.POST("/:id/invites")         //生成项目邀请码 截止时间｜最大可邀请人数
	route.GET("/:id/invites")          //查看项目邀请码列表
	route.GET("/:id/invites/:code")    //查看项目邀请码详情
	route.DELETE("/:id/invites/:code") //删除项目单个邀请码
	route.POST("/invites/:code")       //加入项目

}
