package router

import (
	"app-sosmed/feature/users/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// InitRoute initializes the routes
func InitRoute(e *echo.Echo, uc *handler.UserHandler) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	routeUser(e, uc)
}

func routeUser(e *echo.Echo, uc *handler.UserHandler) {
	e.POST("/register", uc.Register)
	e.POST("/login", uc.Login)
}
