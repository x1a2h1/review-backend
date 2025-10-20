package user

import (
	"net/http"
	"review/internal/dto/req"
	"review/internal/dto/res"
	"review/internal/service"

	"github.com/gin-gonic/gin"
)

func Route(ctx *gin.RouterGroup) {
	group := ctx.Group("user")
	// 获取用户列表-筛选-排序
	group.GET("/list")
	// 创建用户
	group.POST("", create)
	// 更新用户
	group.PUT("/:id")
	// 删除用户
	group.DELETE("/:id")
}

func create(ctx *gin.Context) {
	upsertReq := &req.UserUpsertReq{}
	if err := ctx.ShouldBindJSON(upsertReq); err != nil {
		res.Fail(ctx, http.StatusBadRequest, "参数有误！")
		return
	}
	if err := service.NewUserService().Create(upsertReq); err != nil {
		res.Fail(ctx, http.StatusBadRequest, err.Error())
		return
	}
	res.Success(ctx, nil)
}
