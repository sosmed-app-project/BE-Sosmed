package handler

import (
	"app-sosmed/app/middlewares"
	"app-sosmed/features/users"
	"app-sosmed/helpers"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService users.UserServiceInterface
}

func NewUsersHandler(service users.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (handler *UserHandler) Login(c echo.Context) error {
	var userInput LoginRequest
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}
	dataLogin, token, err := handler.userService.LoginService(userInput.Email, userInput.Password)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error login", nil))

		}
	}
	var response = LoginResponse{
		ID:       dataLogin.ID,
		Username: dataLogin.Username,
		Token:    token,
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success login", response))
}

func (handler *UserHandler) CreateUser(c echo.Context) error {

	var userInput UserRequest
	errBind := c.Bind(&userInput)

	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error bind data. data not valid", nil))
	}

	validate := validator.New()
	if err := validate.Struct(userInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
	}

	userCore := UserRequestToCore(userInput)
	err := handler.userService.Create(userCore)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error insert data", nil))
		}
	}

	return c.JSON(http.StatusCreated, helpers.WebResponse(http.StatusCreated, "success insert data", nil))
}

func (handler *UserHandler) DeleteUser(c echo.Context) error {

	er := middlewares.ExtractTokenUserId(c)
	if er == 0 {
		return c.JSON(http.StatusForbidden, helpers.WebResponse(http.StatusForbidden, "access forbiden", nil))
	}

	id := c.Param("user_id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "operation failed, request resource not valid", nil))
	}

	err := handler.userService.Delete(uint(idConv))
	if err != nil {
		if strings.Contains(err.Error(), "no row affected") {
			return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, "operation failed, requested resource not found", nil))
		}
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error delete data", nil))
	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success delete data", nil))
}

func (handler *UserHandler) GetUser(c echo.Context) error {

	er := middlewares.ExtractTokenUserId(c)
	if er == 0 {
		return c.JSON(http.StatusForbidden, helpers.WebResponse(http.StatusForbidden, "access forbidden", nil))
	}

	result, err := handler.userService.Get(uint(er))
	if err != nil {

		return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, "user not found", nil))

	}

	var userResponse []UserResponse
	for _, value := range result {
		userResponse = append(userResponse, UserResponse{
			ID:           value.ID,
			Username:     value.Username,
			TanggalLahir: value.TanggalLahir,
			Email:        value.Email,
			NoHandphone:  value.NoHandphone,
		})

	}
	return c.JSON(http.StatusOK, helpers.WebResponse(http.StatusOK, "success read data", userResponse))
}

func (handler *UserHandler) UpdateUser(c echo.Context) error {

	er := middlewares.ExtractTokenUserId(c)
	if er == 0 {
		return c.JSON(http.StatusForbidden, helpers.WebResponse(http.StatusForbidden, "access forbidden", nil))
	}

	var userInput UserRequest

	if err := c.Bind(&userInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, "error binding data", nil))
	}

	validate := validator.New()
	if err := validate.Struct(userInput); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.WebResponse(http.StatusBadRequest, err.Error(), nil))
	}

	userCore := UserRequestToCore(userInput)
	if err := handler.userService.Update(er, userCore); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return c.JSON(http.StatusNotFound, helpers.WebResponse(http.StatusNotFound, "user not found", nil))
		}
		return c.JSON(http.StatusInternalServerError, helpers.WebResponse(http.StatusInternalServerError, "error updating user: "+err.Error(), nil))
	}

	return c.JSON(http.StatusCreated, helpers.WebResponse(http.StatusCreated, "success update data", nil))
}
