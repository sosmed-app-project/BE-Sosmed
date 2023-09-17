package handler

import (
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
