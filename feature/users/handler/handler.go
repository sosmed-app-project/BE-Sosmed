package handler

import (
	middleware "hris-app-golang/app/middlewares"
	"hris-app-golang/feature/users"
	"hris-app-golang/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService users.UserServiceInterface
}

func New(service users.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (handler *UserHandler) Add(c echo.Context) error {
	var input UserRequest
	// role_id := middleware.ExtractTokenUserRoleId(c)
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "operation failed, request resource not valid"+errBind.Error(), nil))
	}

	// if role_id == "2" {
	// 	if input.RoleID == "1" || input.RoleID == "2" {
	// 		return c.JSON(http.StatusUnauthorized, helper.WebResponse(http.StatusUnauthorized, "operation failed, request resource not allowed", nil))
	// 	}
	// }

	input.Password = "qwerty"
	var userCore = UserRequestToCore(input)
	err := handler.userService.Add(userCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "operation failed, internal server error", nil))
	}
	return c.JSON(http.StatusOK, helper.WebResponse(http.StatusOK, "operation success, data added", nil))
}

func (handler *UserHandler) GetAll(c echo.Context) error {
	role_id := middleware.ExtractTokenUserRoleId(c)
	division_id := middleware.ExtractTokenUserDivisionId(c)
	result, err := handler.userService.GetAll(role_id, division_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "operation failed, internal server error", nil))
	}
	var userResp []UserResponseAll
	for _, value := range result {
		var user = UserCoreToResponseAll(value)
		userResp = append(userResp, user)
	}
	return c.JSON(http.StatusOK, helper.WebResponse(http.StatusOK, "success", userResp))
}

func (handler *UserHandler) Update(c echo.Context) error {
	id := c.Param("user_id")
	var input UserRequest
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "operation failed, request resource not valid", nil))
	}
	var userCore = UserRequestToCore(input)
	err := handler.userService.Update(id, userCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "operation failed, internal server error", nil))
	}
	return c.JSON(http.StatusOK, helper.WebResponse(http.StatusOK, "success", nil))
}

func (handler *UserHandler) Login(c echo.Context) error {
	var input UserRequest
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "operation failed, request resource not valid", nil))
	}
	var inputCore = UserRequestToCore(input)
	result, token, err := handler.userService.Login(inputCore.Email, inputCore.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "operation failed, internal server error", nil))
	}
	dataResponse := LoginResponse{
		Role:     result.Role.Name,
		Division: result.Division.Name,
		Token:    token,
	}
	return c.JSON(http.StatusOK, helper.WebResponse(http.StatusOK, "success", dataResponse))
}
