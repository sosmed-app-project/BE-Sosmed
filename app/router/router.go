package router

import (
	middlewares "hris-app-golang/app/middlewares"

	_userData "hris-app-golang/feature/users/data"
	_userHandler "hris-app-golang/feature/users/handler"
	_userService "hris-app-golang/feature/users/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, c *echo.Echo) {
	UserData := _userData.New(db)
	UserService := _userService.New(UserData)
	UserHandlerAPI := _userHandler.New(UserService)

	c.POST("/login", UserHandlerAPI.Login)

	c.POST("/users", UserHandlerAPI.Add)
	c.GET("/users", UserHandlerAPI.GetAll, middlewares.JWTMiddleware())
	c.POST("/users/:user_id", UserHandlerAPI.Update, middlewares.JWTMiddleware())
	c.GET("/users/:user_id", UserHandlerAPI.GetUserByID, middlewares.JWTMiddleware())
	c.DELETE("/users/:user_id", UserHandlerAPI.DeleteUser, middlewares.JWTMiddleware())
}
