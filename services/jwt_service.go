package services

import (
	"litigation_backend/models/responses"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func GenerateJWTToken() (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"api_version": "1.0.0",
		"expires_at":  time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return &tokenString, err
}

func VerifyJWTToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := c.Request().Header.Get("Authorization")
		if header == "" {
			return responses.ErrorResponse(c, http.StatusUnauthorized, "Unauthorized")
		}
		tokenString := strings.TrimPrefix(header, "Bearer ")
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		}, jwt.WithValidMethods([]string{jwt.SigningMethodHS512.Alg()}))
		if err != nil {
			return responses.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		}
		if claims, ok := token.Claims.(jwt.MapClaims); !ok {
			return responses.ErrorResponse(c, http.StatusUnauthorized, "Invalid token")
		} else {
			if claims["api_version"] != os.Getenv("API_VERSION") {
				return responses.ErrorResponse(c, http.StatusUnauthorized, "Invalid API version")
			}
			now := time.Now().Unix()
			if now > int64(claims["expires_at"].(float64)) {
				return responses.ErrorResponse(c, 600, "Token is expires")
			}
		}
		return next(c)
	}
}
