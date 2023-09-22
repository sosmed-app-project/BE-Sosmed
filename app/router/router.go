package router

import (
	middlewares "hris-app-golang/app/middlewares"
	"hris-app-golang/helper"
	"net/http"

	_divisionData "hris-app-golang/feature/divisions/data"
	_divisionHandler "hris-app-golang/feature/divisions/handler"
	_divisionService "hris-app-golang/feature/divisions/service"
	_roleData "hris-app-golang/feature/roles/data"
	_roleHandler "hris-app-golang/feature/roles/handler"
	_roleService "hris-app-golang/feature/roles/service"
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

	RoleData := _roleData.NewRoleQuery(db)
	RoleService := _roleService.NewRoleService(RoleData)
	RoleHandlerAPI := _roleHandler.NewRoleHandler(RoleService)

	c.GET("/test", func(c echo.Context) error {
		return c.JSON(http.StatusOK, helper.WebResponse(http.StatusOK, "get test success", nil))
	})

	c.POST("/login", UserHandlerAPI.Login)

	c.POST("/users", UserHandlerAPI.Add, middlewares.JWTMiddleware())
	c.GET("/users", UserHandlerAPI.GetAll, middlewares.JWTMiddleware())
	c.PUT("/users/:user_id", UserHandlerAPI.Update, middlewares.JWTMiddleware())
	c.GET("/users/:user_id", UserHandlerAPI.GetUserByID, middlewares.JWTMiddleware())
	c.DELETE("/users/:user_id", UserHandlerAPI.DeleteUser, middlewares.JWTMiddleware())

	c.GET("/managers", UserHandlerAPI.GetAllManager, middlewares.JWTMiddleware())

	c.GET("/roles", RoleHandlerAPI.GetAllRoles)

	c.GET("/divisions", DivisionHandlerAPI.GetAllDivisions)

	c.GET("/dashboard/", UserHandlerAPI.Dashboard, middlewares.JWTMiddleware())

	// c.GET("/dashboard/jumlahmanager", UserHandlerAPI.GetManagerCount, middlewares.JWTMiddleware())
	// c.GET("/dashboard/jumlahmale_users", UserHandlerAPI.GetMaleUserCount, middlewares.JWTMiddleware())
	// c.GET("/dashboard/jumlahfemale_users", UserHandlerAPI.GetFemaleUserCount, middlewares.JWTMiddleware())

}
