package handler

import (
	levels "hris-app-golang/feature/roles"
	"hris-app-golang/helper"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type RoleHandler struct {
	roleService levels.RoleServiceInterface
}

func New(service levels.RoleServiceInterface) *RoleHandler {
	return &RoleHandler{
		roleService: service,
	}
}

func (handler *RoleHandler) GetAll(c echo.Context) error {
	roleIDStr := c.QueryParam("id")
	roleName := c.QueryParam("name")

	var roleID uint
	if roleIDStr != "" {
		id, err := strconv.ParseUint(roleIDStr, 10, 32)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "Invalid id parameter", nil))
		}
		roleID = uint(id)
	}

	result, err := handler.roleService.GetAll(roleID, roleName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "Operation failed, internal server error", nil))
	}

	var roleResp []RoleResponseAll
	for _, value := range result {
		var role = RoleCoreToResponseAll(value)
		roleResp = append(roleResp, role)
	}

	return c.JSON(http.StatusOK, helper.WebResponse(http.StatusOK, "Success", roleResp))
}

func (handler *RoleHandler) UpdateById(c echo.Context) error {
	roleInput := new(RoleRequest)

	id := c.Param("roles_id")
	idParam, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "error data id. data not valid", nil))
	}
	errBind := c.Bind(&roleInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}
	roleCore := RequestToCore(*roleInput)
	result, err := handler.roleService.UpdateRoleById(uint(idParam), roleCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "error insert data", nil))
		}
	}

	updateResponse := RoleResponse{
		ID:   uint(idParam),
		Name: result.Name,
	}
	return c.JSON(http.StatusOK, helper.WebResponse(http.StatusOK, "success update data", updateResponse))
}
