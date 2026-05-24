package routes

import (
	"litigation_backend/controllers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Group) {
	e.POST("/test-whatsapp", controllers.TestWhatsapp)
}
