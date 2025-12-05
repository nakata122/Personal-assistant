package users

type UserRole string

const (
    RoleAdmin UserRole = "admin"
    RoleUser  UserRole = "user"
    RoleGuest UserRole = "guest"
)

type User struct {
	UserID   int      `json:"user_id"`
	Email    string   `json:"email"`
	GoogleID string   `json:"google_id"`
	Role	 UserRole `json:"role"`
	Name     string   `json:"name"`
	Picture  string   `json:"picture"`
}