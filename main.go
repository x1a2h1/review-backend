package main

import (
	"review/internal/api/comm"
	"review/internal/api/user"
	"review/internal/core"
	"review/internal/middleware/authc"
	_ "review/internal/pkg/config"
	"review/internal/pkg/database"
)

func main() {
	cmd := core.New()
	cmd.Use(authc.AuthenticationHandler)
	cmd.Router("v1", comm.Route, user.Route)
	closeMysql := database.InitializeMysql()
	defer closeMysql()
	cmd.Run()
}
