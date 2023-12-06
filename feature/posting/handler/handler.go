package handler

import (
	"hris-app-golang/feature/roles"
	"hris-app-golang/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RoleHandler struct {
	roleService roles.RoleServiceInterface
}

func NewRoleHandler(service roles.RoleServiceInterface) *RoleHandler {
	return &RoleHandler{
		roleService: service,
	}
}

func (handler *RoleHandler) GetAllRoles(c echo.Context) error {

	result, err := handler.roleService.GetAllRoles()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "error read data", nil))
	}
	var roleResponse []RoleResponse
	for _, value := range result {
		roleResponse = append(roleResponse, RoleResponse{
			ID:   value.ID,
			Name: value.Name,
		})

	}
	return c.JSON(http.StatusOK, helper.WebResponse(http.StatusOK, "success read data", roleResponse))
}
