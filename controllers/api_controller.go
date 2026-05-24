package controllers

import (
	"litigation_backend/services"

	"github.com/labstack/echo/v4"
)

func TestWhatsapp(c echo.Context) error {
	err := services.TestWhatsapp()
	if err != nil {
		return c.JSON(500, echo.Map{
			"status":  "error",
			"message": "WhatsApp integration test failed",
		})
	}
	return c.JSON(200, echo.Map{
		"status":  "success",
		"message": "WhatsApp integration test successful",
	})
}
