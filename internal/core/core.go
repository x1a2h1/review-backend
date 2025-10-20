package core

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type Core struct {
	core *gin.Engine
}

func New() *Core {
	router := gin.Default()

	// 配置 session 中间件
	store := cookie.NewStore([]byte("reviewappsecretkeyx1a2h1"))
	store.Options(sessions.Options{
		Path:     "/",
		Domain:   "",
		MaxAge:   86400 * 7, // 7天
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
	router.Use(sessions.Sessions("LOGIN_SESSION", store))

	return &Core{core: router}
}

func (c *Core) Use(middleware ...func() gin.HandlerFunc) {
	for _, fn := range middleware {
		c.core.Use(fn())
	}
}

func (c *Core) Router(path string, routes ...func(*gin.RouterGroup)) {
	api := c.core.Group(fmt.Sprintf("/api/%s", path))
	for _, fn := range routes {
		fn(api)
	}
}

func (c *Core) Run() {
	srv := &http.Server{
		Addr:    "[::]:" + "2333",
		Handler: c.core,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("服务启动失败，监听: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("Server Shutdown:", err)
	}
	fmt.Println("http api 服务退出！")
}
