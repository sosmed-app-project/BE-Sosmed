package router

import (
	middlewares "hris-app-golang/app/middlewares"

	_divisionData "hris-app-golang/feature/divisions/data"
	_divisionHandler "hris-app-golang/feature/divisions/handler"
	_divisionService "hris-app-golang/feature/divisions/service"
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
	DivisionData := _divisionData.NewDivisionQuery(db)
	DivisionService := _divisionService.NewDivisionService(DivisionData)
	DivisionHandlerAPI := _divisionHandler.NewDivisionsHandler(DivisionService)

	c.POST("/login", UserHandlerAPI.Login)

	c.POST("/users", UserHandlerAPI.Add)
	c.GET("/users", UserHandlerAPI.GetAll, middlewares.JWTMiddleware())
	c.POST("/users/:user_id", UserHandlerAPI.Update, middlewares.JWTMiddleware())
	c.GET("/users/:user_id", UserHandlerAPI.GetUserByID, middlewares.JWTMiddleware())
	c.DELETE("/users/:user_id", UserHandlerAPI.DeleteUser, middlewares.JWTMiddleware())

	c.GET("/divisions", DivisionHandlerAPI.GetAllDivisions)
}
