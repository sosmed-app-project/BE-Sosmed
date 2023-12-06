package handler

import (
	"hris-app-golang/feature/divisions"
	"hris-app-golang/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DivisionHandler struct {
	divisionService divisions.DivisionServiceInterface
}

func NewDivisionsHandler(service divisions.DivisionServiceInterface) *DivisionHandler {
	return &DivisionHandler{
		divisionService: service,
	}
}

func (handler *DivisionHandler) GetAllDivisions(c echo.Context) error {

	result, err := handler.divisionService.GetDiv()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, helper.WebResponse(http.StatusInternalServerError, "error read data", nil))
	}
	var divResponse []DivisionsResponse
	for _, value := range result {
		divResponse = append(divResponse, DivisionsResponse{
			ID:   value.ID,
			Name: value.Name,
		})

	}
	return c.JSON(http.StatusOK, helper.WebResponse(http.StatusOK, "success read data", divResponse))
}
