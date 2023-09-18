package handler

import (
	middleware "hris-app-golang/app/middlewares"
	"hris-app-golang/feature/users"
	"hris-app-golang/helper"
	"net/http"
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

/*
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
*/

func (handler *UserHandler) GetUserByID(c echo.Context) error {
	id := c.Param("user_id")

	result, err := handler.userService.GetById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Id Harus berupa string")
	}

	resultDivision := DivisionResponse{
		ID:   result.Division.ID,
		Name: result.Division.Name,
	}

	resultRole := RoleResponse{
		ID:   result.Role.ID,
		Name: result.Role.Name,
	}

	resultImportantData := UserImportantDataResponse{
		ID:          result.UserImport.UserID,
		Birth_Place: result.UserImport.BirthPlace,
		Birth_Date:  result.UserImport.BirthDate,
		Religion:    result.UserImport.Religion,
	}

	resultResponse := UserResponse{
		ID:                  result.ID,
		First_Name:          result.FirstName,
		Email:               result.Email,
		Phone_Number:        result.PhoneNumber,
		Address:             result.Address,
		Division:            resultDivision,
		Role:                resultRole,
		User_Important_Data: resultImportantData,
	}

	return c.JSON(http.StatusOK, helper.WebResponse(http.StatusOK, "Success get user data", resultResponse))
}

func (handler *UserHandler) DeleteUser(c echo.Context) error {
	id := c.Param("user_id")

	err := handler.userService.Delete(id)
	if err != nil {
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
	response := map[string]any{
		"token": token,
		"role":  dataLogin.Role,
		"id":    dataLogin.ID,
		"email": dataLogin.Email,
	}
	return c.JSON(http.StatusOK, helper.WebResponse(http.StatusOK, "success login", response))
}
