package recorder

import (
	"log/slog"
	"review/internal/models"
	"review/internal/service"
	"time"

	"github.com/gin-gonic/gin"
)

func RecordLoginLog(ctx *gin.Context, user *models.User) {
	// 是否记录登录日志
	// if !config.GetConfig().Log.LoginLog {
	// 	return
	// }
	logService := service.NewLogService()
	userAgent := ctx.Request.UserAgent()
	clientIp := ctx.ClientIP()
	log := &models.LoginLog{
		UserAgent: userAgent,
		Ip:        clientIp,
	}
	log.UserId = user.ID
	log.Nickname = user.Nickname
	log.Base = models.Base{CreatedAt: time.Now(), UpdatedAt: time.Now()}
	if err := logService.AddLoginLog(log); err != nil {
		slog.Error("-", slog.Any("err", err))
	}
}
