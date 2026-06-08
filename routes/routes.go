package routes

import (
	"litigation_backend/controllers"
	"litigation_backend/services"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Group) {
	e.POST("/test-whatsapp", controllers.TestWhatsapp)
	admin := e.Group("/admin")
	{
		admin.POST("/login", controllers.AdminLogin)
		authAdmin := admin.Group("", services.VerifyJWTToken)
		{
			authAdmin.GET("/users", controllers.GetAllUsers)
		}
	}
}
