package responses

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ErrorResponse(c echo.Context, statusCode int, message string) error {
	return c.JSON(statusCode, echo.Map{
		"status":  statusCode,
		"message": message,
	})
}

func SuccessResponse(c echo.Context, data any) error {
	return c.JSON(http.StatusOK, echo.Map{
		"status":  http.StatusOK,
		"message": "Success",
		"data":    data,
	})
}
