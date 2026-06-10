package controllers

import (
	"litigation_backend/models/requests"
	"litigation_backend/models/responses"
	"litigation_backend/services"
	"os"

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

func AdminLogin(c echo.Context) error {
	var request requests.AdminLoginRequest
	err := c.Bind(&request)
	if err != nil {
		return responses.ErrorResponse(c, 400, err.Error())
	}
	adminEmail := os.Getenv("ADMIN_EMAIL")
	adminPassword := os.Getenv("ADMIN_PASSWORD")
	if request.Email != adminEmail || request.Password != adminPassword {
		return responses.ErrorResponse(c, 401, "Invalid email or password")
	}
	token, err := services.GenerateJWTToken()
	if err != nil {
		return responses.ErrorResponse(c, 500, "Failed to generate token")
	}
	return c.JSON(200, echo.Map{
		"status":  "success",
		"message": "Login successful",
		"token":   *token,
	})
}

func GetAllUsers(c echo.Context) error {
	users, err := services.GetAllUserFromDb()
	if err != nil {
		return responses.ErrorResponse(c, 500, err.Error())
	}
	return responses.SuccessResponse(c, users)
}

func CreateUser(c echo.Context) error {
	var request requests.CreateUserRequest
	err := c.Bind(&request)
	if err != nil {
		return responses.ErrorResponse(c, 400, err.Error())
	}
	user, err := services.CreateUser(&request)
	if err != nil {
		return responses.ErrorResponse(c, 500, err.Error())
	}
	return responses.SuccessResponse(c, user)
}

func DisableUser(c echo.Context) error {
	uid := c.Param("uid")
	user, err := services.DisableUser(uid)
	if err != nil {
		return responses.ErrorResponse(c, 500, err.Error())
	}
	return responses.SuccessResponse(c, user)
}

func ForgotPasswordUser(c echo.Context) error {
	uid := c.Param("uid")
	email, err := services.ForgotPasswordUser(uid)
	if err != nil {
		return responses.ErrorResponse(c, 500, err.Error())
	}
	return responses.SuccessResponse(c, "Email send to: "+*email)
}

func AdminDashboadData(c echo.Context) error {
	data, err := services.AdminDashboardData()
	if err != nil {
		return responses.ErrorResponse(c, 500, err.Error())
	}
	return responses.SuccessResponse(c, data)
}

func GetAllCases(c echo.Context) error {
	data, err := services.GetAllCases()
	if err != nil {
		return responses.ErrorResponse(c, 500, err.Error())
	}
	return responses.SuccessResponse(c, data)
}
