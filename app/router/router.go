package router

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	userD "hris-app-golang/feature/users/data"
	userH "hris-app-golang/feature/users/handler"
	userS "hris-app-golang/feature/users/service"
)

func InitRouter(db *gorm.DB, c *echo.Echo) {
	userData := userD.New(db)
	userService := userS.New(userData)
	userHandlerAPI := userH.New(userService)

	c.GET("/users/:user_id", userHandlerAPI.GetUserByID)
}
