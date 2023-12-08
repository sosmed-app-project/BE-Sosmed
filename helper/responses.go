// helper/response.go
package helper

import (
	"github.com/labstack/echo/v4"
)

// RespondWithError creates a JSON response with an error message
func RespondWithError(c echo.Context, code int, message string) error {
	return c.JSON(code, map[string]string{"error": message})
}

// RespondWithJSON creates a JSON response with the provided data
func RespondWithJSON(c echo.Context, code int, data interface{}) error {
	return c.JSON(code, data)
}
