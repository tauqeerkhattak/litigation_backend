package services

import (
	"context"
	"litigation_backend/config"
	"litigation_backend/models/requests"
	"litigation_backend/models/responses"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"firebase.google.com/go/v4/auth"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/wneessen/go-mail"
	"google.golang.org/api/iterator"
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

func TestWhatsapp() error {
	phoneNumberId := os.Getenv("WHATSAPP_PHONE_NUMBER_ID")
	url := "https://graph.facebook.com/v25.0/" + phoneNumberId + "/messages"
	body := map[string]any{
		"messaging_product": "whatsapp",
		"to":                "+923702260602",
		"type":              "template",
		"template": map[string]any{
			"name":     "hello_world",
			"language": map[string]any{"code": "en_US"},
		},
	}
	err, response := POST(url, body)
	if err != nil {
		return err
	}
	log.Println("RESPONSE: ", response)
	return nil
}

func GetAllUsers(pageToken string) ([]*auth.ExportedUserRecord, string) {
	users := make([]*auth.ExportedUserRecord, 0)
	iter := config.Auth.Users(context.Background(), pageToken)
	for {
		user, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Println("ERROR: ", user)
		}
		users = append(users, user)
	}
	return users, iter.PageInfo().Token
}

func CreateUser(request *requests.CreateUserRequest) (*auth.UserRecord, error) {
	userToCreate := (&auth.UserToCreate{})
	userToCreate.Email(request.Email)
	userToCreate.Password(request.Password)
	userToCreate.DisplayName(request.Name)
	user, err := config.Auth.CreateUser(context.Background(), userToCreate)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func DisableUser(uid string) (*auth.UserRecord, error) {
	userToUpdate := (&auth.UserToUpdate{}).Disabled(true)
	return config.Auth.UpdateUser(context.Background(), uid, userToUpdate)
}

func ForgotPasswordUser(uid string) (*string, error) {
	user, err := config.Auth.GetUser(context.Background(), uid)
	if err != nil {
		return nil, err
	}
	link, err := config.Auth.PasswordResetLink(context.Background(), user.Email)
	if err != nil {
		return nil, err
	}
	log.Printf("email=%s password_len=%d",
		os.Getenv("SMTP_EMAIL"),
		len(os.Getenv("SMTP_PASSWORD")),
	)
	message := mail.NewMsg()
	if err := message.From(os.Getenv("SMTP_EMAIL")); err != nil {
		return nil, err
	}
	if err := message.To(user.Email); err != nil {
		return nil, err
	}
	message.Subject("Password Reset Email")
	message.SetBodyString(mail.TypeTextHTML,
		`<h1>LITIGATION MANAGEMENT</h1>
    <p>Use the following link to reset your password:</p>
    <p><a href="`+link+`">Reset Password</a></p>`,
	)
	client, err := mail.NewClient("smtp.gmail.com",
		mail.WithPort(587),
		mail.WithTLSPolicy(mail.TLSMandatory),
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(os.Getenv("SMTP_EMAIL")),
		mail.WithPassword(os.Getenv("SMTP_PASSWORD")),
	)
	if err != nil {
		return nil, err
	}
	if err := client.DialAndSend(message); err != nil {
		log.Println("ERROR: ", err)
		return nil, err
	}
	return &user.Email, nil
}
