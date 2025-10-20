package comm

import (
	"fmt"
	"net/http"
	"review/internal/dto/req"
	"review/internal/dto/res"
	"review/internal/middleware/recorder"
	"review/internal/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func login(ctx *gin.Context) {
	loginReq := &req.LoginReq{}
	if err := ctx.ShouldBindJSON(loginReq); err != nil {
		res.Fail(ctx, http.StatusBadRequest, "参数有误！")
		return
	}
	user, err := service.NewAuthService().Login(loginReq)
	if err != nil {
		res.Fail(ctx, http.StatusBadRequest, err.Error())
		return
	}
	// session存储
	session := sessions.Default(ctx)
	session.Set("_userId_", user.ID)
	session.Set("_isRootUser_", user.Root)
	if err := session.Save(); err != nil {
		res.Fail(ctx, http.StatusBadRequest, err.Error())
		return
	}
	recorder.RecordLoginLog(ctx, user)
	res.Success(ctx, user)
}

// me 获取个人信息
func me(ctx *gin.Context) {
	userId := ctx.GetUint("_userId_")
	fmt.Println("用户id", userId)
	user := service.NewUserService().FindOneById(userId)
	if user == nil {
		res.Fail(ctx, http.StatusUnauthorized, "登录失效")
		return
	}
	res.Success(ctx, user)
}
