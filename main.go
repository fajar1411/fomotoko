package main

import (
	"fmt"
	"test/config"
	"test/driver/mysql"
	fr "test/faktoryandroute"
	"test/migrasi"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitDB(cfg)

	migrasi.MigrateDB(db)

	e := echo.New()

	fr.FaktoryAndRoute(e, db)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%d", cfg.SERVER_PORT)))
}
