package services

import (
	"context"
	"errors"
	"litigation_backend/config"
	"litigation_backend/models/requests"
	"litigation_backend/models/responses"
	"log"
	"os"

	"firebase.google.com/go/v4/auth"
	"github.com/wneessen/go-mail"
)

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

func CreateUser(request *requests.CreateUserRequest) (*responses.User, error) {
	if request.Role == requests.UserRole("admin") {
		return nil, errors.New("Cannot create admin!")
	}
	userToCreate := (&auth.UserToCreate{})
	userToCreate.Email(request.Email)
	userToCreate.Password(request.Password)
	userToCreate.DisplayName(request.Name)
	authUser, err := config.Auth.CreateUser(context.Background(), userToCreate)
	if err != nil {
		return nil, err
	}
	user, err := CreateUserInDb(authUser.UID, request)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func DisableUser(uid string) (*responses.User, error) {
	userToUpdate := (&auth.UserToUpdate{}).Disabled(true)
	_, err := config.Auth.UpdateUser(context.Background(), uid, userToUpdate)
	if err != nil {
		return nil, err
	}
	return DisableUserInDb(uid)
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

func AdminDashboardData() (*responses.DashboardData, error) {
	userCount, err := GetUserCount()
	if err != nil {
		return nil, err
	}
	casesCount, err := GetCasesCount()
	data := responses.DashboardData{
		TotalUsers:    *userCount,
		ActiveCases:   *casesCount,
		HearingsToday: 0,
		UrgentTasks:   0,
	}
	return &data, nil
}
