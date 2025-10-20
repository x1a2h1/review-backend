package authc

import (
	"fmt"
	"net/http"
	"path"
	"review/internal/dto/res"
	"review/internal/utils"
	"strconv"
	"strings"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gobwas/glob"
)

var ignorePatterns = []string{"/api/v1/comm/login", "/api/v1/comm/upload"}
var websocketPatterns []string

func AuthenticationHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if ignored(c.Request.URL.Path) {
			c.Next()
		} else {
			isWs := isWebsocket(c.Request.URL.Path)
			if isWs {
				token := c.Query("token")
				decryptToken, _ := utils.DecryptString("reviewappsecretkeyx1a2h1", token, "")
				index := strings.LastIndex(decryptToken, "$")
				if index != -1 {
					loginToken, timestampStr := decryptToken[:index], decryptToken[index+1:]
					timestamp, _ := strconv.Atoi(timestampStr)
					if time.Now().UnixMilli()-int64(timestamp) < 15*1000 {
						c.Request.AddCookie(&http.Cookie{
							Name:  "LOGIN_SESSION",
							Value: loginToken,
						})
					}
				}
			}
			session := sessions.Default(c)
			userId := session.Get("_userId_")
			if userId == nil {
				if isWs {
					c.AbortWithStatus(http.StatusUnauthorized)
				} else {
					res.Fail(c, http.StatusUnauthorized, "授权无效，请重新登录！")
				}
				return
			}
			c.Set("_userId_", userId)
			c.Set("_isRootUser_", session.Get("_isRootUser_"))
			c.Next()
		}
	}
}
func ignored(requestPath string) bool {
	for _, pattern := range ignorePatterns {
		g, err := glob.Compile(pattern, '/')
		if err != nil {
			panic(fmt.Errorf("pattern '%s' compile error: %w", pattern, err))
		}
		if g.Match(path.Clean(requestPath)) {
			return true
		}
	}
	return false
}
func isWebsocket(requestPath string) bool {
	for _, pattern := range websocketPatterns {
		g, err := glob.Compile(pattern, '/')
		if err != nil {
			panic(fmt.Errorf("pattern '%s' compile error: %w", pattern, err))
		}
		if g.Match(path.Clean(requestPath)) {
			return true
		}
	}
	return false
}
