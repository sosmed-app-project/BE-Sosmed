package handler

import (
	"fmt"
	middleware "hris-app-golang/app/middlewares"
	"hris-app-golang/feature/users"
	"hris-app-golang/helper"
	"net/http"
	"reflect"
	"strconv"
	"strings"

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
	fmt.Println("lead_id before mapping to core:", input.UserLeadID, reflect.TypeOf(input.UserLeadID))
	// if role_id == "2" {
	// 	if input.RoleID == "1" || input.RoleID == "2" {
	// 		return c.JSON(http.StatusUnauthorized, helper.WebResponse(http.StatusUnauthorized, "operation failed, request resource not allowed", nil))
	// 	}
	// }

	input.Password = "qwerty"
	var userCore = UserRequestToCore(input)
	fmt.Println("lead_id after mapping to core:", userCore.UserLeadID, reflect.TypeOf(input.UserLeadID))
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
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "operation failed, request resource not valid", nil))
	}
	var input UserRequest
	errBind := c.Bind(&input)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "operation failed, request resource not valid", nil))
	}
	var userCore = UserRequestToCore(input)
	err := handler.userService.Update(uint(idConv), userCore)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "operation failed, internal server error", nil))
	}
	return c.JSON(http.StatusOK, helper.WebResponse(http.StatusOK, "success", nil))
}

func (handler *UserHandler) GetUserByID(c echo.Context) error {
	id := c.Param("user_id")
	role_id := middleware.ExtractTokenUserRoleId(c)
	user_id := middleware.ExtractTokenUserId(c)

	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "operation failed, request resource not valid", nil))
	}

	if role_id == 3 || role_id == 4 {
		if uint(idConv) != user_id {
			return c.JSON(http.StatusUnauthorized, helper.WebResponse(http.StatusUnauthorized, "operation failed, request resource not allowed", nil))
		}
	}

	result, err := handler.userService.GetById(uint(idConv))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Id Harus berupa string")
	}

	resultResponse := UserCoreToResponse(result)

	return c.JSON(http.StatusOK, helper.WebResponse(http.StatusOK, "Success get user data", resultResponse))
}

func (handler *UserHandler) DeleteUser(c echo.Context) error {
	role_id := middleware.ExtractTokenUserRoleId(c)
	if role_id != 2 && role_id != 1 {
		return c.JSON(http.StatusUnauthorized, helper.WebResponse(http.StatusUnauthorized, "opeartion failed, request resource not allowed", nil))
	}
	id := c.Param("user_id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "operation failed, request resource not valid", nil))
	}

	err := handler.userService.Delete(uint(idConv))
	if err != nil {
		if strings.Contains(err.Error(), "no row affected") {
			return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "error delete data, data not found", nil))
		}
		return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "error delete data", nil))
	}
	return c.JSON(http.StatusOK, helper.WebResponse(http.StatusOK, "success delete data", nil))
}

func (handler *UserHandler) Login(c echo.Context) error {
	userInput := new(LoginRequest)
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}
	dataLogin, token, err := handler.userService.Login(userInput.Email, userInput.Password)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helper.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "error login", nil))

		}
	}
	var response = LoginResponse{
		Role:     dataLogin.Role.Name,
		Division: dataLogin.Division.Name,
		Token:    token,
	}
	return c.JSON(http.StatusOK, helper.WebResponse(http.StatusOK, "success login", response))
}
