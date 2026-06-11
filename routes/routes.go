package routes

import (
	"litigation_backend/controllers"
	"litigation_backend/services"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Group) {
	admin := e.Group("/admin")
	{
		admin.POST("/login", controllers.AdminLogin)
		authAdmin := admin.Group("", services.VerifyJWTToken)
		{
			authAdmin.GET("/test-whatsapp", controllers.TestWhatsapp)
			users := authAdmin.Group("/users")
			{
				users.GET("", controllers.GetAllUsers)
				users.POST("/create", controllers.CreateUser)
				users.DELETE("/:uid", controllers.DisableUser)
				users.POST("/:uid/forgot-password", controllers.ForgotPasswordUser)
			}
			authAdmin.GET("/dashboard", controllers.AdminDashboadData)
			authAdmin.GET("/cases", controllers.GetAllCases)
		}
	}
}
