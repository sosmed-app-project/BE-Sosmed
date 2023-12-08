package router

import (
	"app-sosmed/app/middlewares"
	_userData "app-sosmed/features/users/data"
	_userHandler "app-sosmed/features/users/handler"
	_userService "app-sosmed/features/users/service"
	"app-sosmed/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, c *echo.Echo) {
	UserData := _userData.NewUsersQuery(db)
	UserService := _userService.NewUsersLogic(UserData)
	UserHandlerAPI := _userHandler.NewUsersHandler(UserService)

	c.GET("/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "get test success", nil))
	})

	c.POST("/login", UserHandlerAPI.Login)
	c.POST("/register", UserHandlerAPI.CreateUser)
	c.DELETE("/user/:user_id", UserHandlerAPI.DeleteUser, middlewares.JWTMiddleware())
	c.GET("/profile", UserHandlerAPI.GetUser, middlewares.JWTMiddleware())
	c.PUT("/update", UserHandlerAPI.UpdateUser, middlewares.JWTMiddleware())
}
