package services

import (
	"context"
	"errors"
	"litigation_backend/config"
	"litigation_backend/models/requests"
	"litigation_backend/models/responses"
	"log"
	"os"
	"time"

	"firebase.google.com/go/v4/auth"
	"github.com/wneessen/go-mail"
)

var SECONDS_IN_3_DAYS = 259200

func CheckCasesAndSendReminders() error {
	cases, err := GetAllCases()
	if err != nil {
		return err
	}
	for _, caseModel := range cases {
		nextHearing := caseModel.NextHearing
		if nextHearing == nil {
			continue
		}
		userId := caseModel.UserId
		if userId == nil {
			log.Println("FOR CASE: "+caseModel.CaseNo+", USER ID NOT FOUND!", time.Now())
			continue
		}
		now := time.Now().Add(time.Hour * 24 * 3)
		if now.Day() == nextHearing.Day() && now.Month() == nextHearing.Month() && now.Year() == nextHearing.Year() {
			user, err := GetUserByUid(*userId)
			if err != nil {
				log.Println("FOR CASE: "+caseModel.CaseNo+", COULD NOT GET USER, ERROR: ", err.Error())
				continue
			}
			phoneNumber := user.CountryCode + user.PhoneNumber
			err = SendWhatsappReminder(phoneNumber)
			if err != nil {
				log.Println("FOR CASE: "+caseModel.CaseNo+", COULD NOT SEND REMINDER, ERROR: ", err.Error())
			}
		}
	}
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
