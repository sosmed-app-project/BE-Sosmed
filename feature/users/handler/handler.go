package handler

import (
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

	err := handler.userService.DeleteUserById(id)
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
	dataLogin, token, err := handler.userService.LoginUser(userInput.Email, userInput.Password)
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
