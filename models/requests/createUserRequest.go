package requests

type CreateUserRequest struct {
	Email       string   `json:"email" validate:"required,email"`
	Password    string   `json:"password" validate:"required,gte=8"`
	Name        string   `json:"name" validate:"required"`
	Role        UserRole `json:"role" validate:"required"`
	CountryCode string   `json:"country_code" validate:"required"`
	PhoneNumber string   `json:"phone_number" validate:"required"`
}

type UserRole string

const (
	Admin      UserRole = "admin"      // Admin, Do whatever it wants!
	Editor     UserRole = "editor"     // Can create cases, hearings, and add documents
	Commentor  UserRole = "commentor"  // Cannot create cases, can create hearings and add documents
	Documentor UserRole = "documentor" // Cannot create cases and hearings, can only add documents
	Viewer     UserRole = "viewer"     // Only view cases, hearings, and documents
)
