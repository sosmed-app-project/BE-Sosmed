package router

import (
	middlewares "hris-app-golang/app/middlewares"

	_userData "hris-app-golang/feature/users/data"
	_userHandler "hris-app-golang/feature/users/handler"
	_userService "hris-app-golang/feature/users/service"

	_roleData "hris-app-golang/feature/roles/data"
	_roleHandler "hris-app-golang/feature/roles/handler"
	_roleService "hris-app-golang/feature/roles/service"

	_divisionData "hris-app-golang/feature/divisions/data"
	_divisionHandler "hris-app-golang/feature/divisions/handler"
	_divisionService "hris-app-golang/feature/divisions/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, c *echo.Echo) {
	UserData := _userData.New(db)
	UserService := _userService.New(UserData)
	UserHandlerAPI := _userHandler.New(UserService)

	RoleData := _roleData.New(db)
	RoleService := _roleService.New(RoleData)
	RoleHandlerAPI := _roleHandler.New(RoleService)

	DivisionData := _divisionData.NewDivisionQuery(db)
	DivisionService := _divisionService.NewDivisionService(DivisionData)
	DivisionHandlerAPI := _divisionHandler.NewDivisionsHandler(DivisionService)

	c.POST("/login", UserHandlerAPI.Login)

	c.POST("/users", UserHandlerAPI.Add)
	c.GET("/users", UserHandlerAPI.GetAll, middlewares.JWTMiddleware())
	c.POST("/users/:user_id", UserHandlerAPI.Update, middlewares.JWTMiddleware())

	c.GET("/roles", RoleHandlerAPI.GetAllRole, middlewares.JWTMiddleware())
	c.PUT("/roles/:roles_id", RoleHandlerAPI.UpdateById)
	c.GET("/users/:user_id", UserHandlerAPI.GetUserByID, middlewares.JWTMiddleware())
	c.DELETE("/users/:user_id", UserHandlerAPI.DeleteUser, middlewares.JWTMiddleware())

	c.GET("/divisions", DivisionHandlerAPI.GetAllDivisions, middlewares.JWTMiddleware())
}
