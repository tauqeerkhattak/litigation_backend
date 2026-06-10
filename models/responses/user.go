package responses

import (
	"litigation_backend/models/requests"
	"time"
)

type User struct {
	Uid       string            `json:"uid"`
	Email     string            `json:"email"`
	Name      string            `json:"name"`
	Role      requests.UserRole `json:"role"`
	Disabled  bool              `json:"disabled"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

func UserFromJson(data map[string]any) User {
	return User{
		Uid:       data["uid"].(string),
		Email:     data["email"].(string),
		Name:      data["name"].(string),
		Role:      data["role"].(requests.UserRole),
		CreatedAt: data["created_at"].(time.Time),
		UpdatedAt: data["updated_at"].(time.Time),
	}
}
