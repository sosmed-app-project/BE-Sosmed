package router

import (
	middlewares "hris-app-golang/app/middlewares"

	_userData "hris-app-golang/feature/users/data"
	_userHandler "hris-app-golang/feature/users/handler"
	_userService "hris-app-golang/feature/users/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	userD "hris-app-golang/feature/users/data"
	userH "hris-app-golang/feature/users/handler"
	userS "hris-app-golang/feature/users/service"
)

func InitRouter(db *gorm.DB, c *echo.Echo) {
<<<<<<< HEAD
	UserData := _userData.New(db)
	UserService := _userService.New(UserData)
	UserHandlerAPI := _userHandler.New(UserService)

	c.POST("/login", UserHandlerAPI.Login)

	c.POST("/users", UserHandlerAPI.Add)
	c.GET("/users", UserHandlerAPI.GetAll, middlewares.JWTMiddleware())
	c.POST("/users/:user_id", UserHandlerAPI.Update, middlewares.JWTMiddleware())
=======
	userData := userD.New(db)
	userService := userS.New(userData)
	userHandlerAPI := userH.New(userService)

	c.GET("/users/:user_id", userHandlerAPI.GetUserByID)
	c.DELETE("/users/:user_id", userHandlerAPI.DeleteUser)
	c.POST("/login", userHandlerAPI.Login)
>>>>>>> d65206d7f3cdd592e18a4c27ff9f11dcaa5ebf9b
}
