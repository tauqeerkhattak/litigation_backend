package responses

import (
	"litigation_backend/models/requests"
	"log"
	"time"
)

type User struct {
	Uid       string            `json:"uid" firestore:"uid"`
	Email     string            `json:"email" firestore:"email"`
	Name      string            `json:"name" firestore:"name"`
	Role      requests.UserRole `json:"role" firestore:"role"`
	Disabled  bool              `json:"disabled" firestore:"disabled"`
	CreatedAt time.Time         `json:"created_at" firestore:"created_at"`
	UpdatedAt time.Time         `json:"updated_at" firestore:"updated_at"`
}

func UserFromJson(data map[string]any) User {
	log.Println("USER: ", data)
	return User{
		Uid:       data["uid"].(string),
		Email:     data["email"].(string),
		Name:      data["name"].(string),
		Disabled:  data["disabled"].(bool),
		Role:      requests.UserRole(data["role"].(string)),
		CreatedAt: data["created_at"].(time.Time),
		UpdatedAt: data["updated_at"].(time.Time),
	}
}
