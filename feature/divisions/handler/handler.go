package handler

import (
	"hris-app-golang/feature/divisions"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type DivisionHandler struct {
	divisionService divisions.DivisionDataInterface
}

func NewDivisionHandler(divisionService divisions.DivisionDataInterface) *DivisionHandler {
	return &DivisionHandler{
		divisionService: divisionService,
	}
}

func (h *DivisionHandler) CreateDivision(c echo.Context) error {
	var division divisions.DivisionCore
	if err := c.Bind(&division); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err := h.divisionService.Insert(division)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "Division created successfully"})
}

func (h *DivisionHandler) GetAllDivisions(c echo.Context) error {
	divisions, err := h.divisionService.SelectAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, divisions)
}

func (h *DivisionHandler) GetDivisionByID(c echo.Context) error {
	id := c.Param("id")
	divisionID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid division ID"})
	}

	division, err := h.divisionService.SelectByID(uint(divisionID))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, division)
}

func (h *DivisionHandler) UpdateDivision(c echo.Context) error {
	id := c.Param("id")
	divisionID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid division ID"})
	}

	var division divisions.DivisionCore
	if err := c.Bind(&division); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err = h.divisionService.Update(uint(divisionID), division)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Division updated successfully"})
}

func (h *DivisionHandler) DeleteDivision(c echo.Context) error {
	id := c.Param("id")
	divisionID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid division ID"})
	}

	err = h.divisionService.Delete(uint(divisionID))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Division deleted successfully"})
}