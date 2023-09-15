package main

import (
	"hris-app-golang/app/config"
	"hris-app-golang/app/database"
	"hris-app-golang/app/router"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.InitConfig()
	mysql := database.InitMysql(cfg)
	database.InittialMigration(mysql)

	e := echo.New()
	// e.Pre(middleware.RemoveTrailingSlash())
	// e.Use(middleware.CORS())
	// e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
	// Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	// }))
	router.InitRouter(mysql, e)
	e.Logger.Fatal(e.Start(":80"))
}
