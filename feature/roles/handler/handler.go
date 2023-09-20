package handler

import (
	"hris-app-golang/feature/roles"
)

type RoleHandler struct {
	roleService roles.RoleServiceInterface
}

func New(service roles.RoleServiceInterface) *RoleHandler {
	return &RoleHandler{
		roleService: service,
	}
}

/*
func (handler *RoleHandler) GetAll(c echo.Context) {
	var result, err = handler.roleService.GetAllRoles()

	return
}
*/
