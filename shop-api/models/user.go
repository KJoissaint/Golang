package models

import "time"

type Role string

const (
	RoleSuperAdmin Role = "SuperAdmin"
	RoleAdmin      Role = "Admin"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"` // Never expose password in JSON
	Role      Role      `json:"role"`
	ShopID    int       `json:"shop_id"`
	CreatedAt time.Time `json:"created_at"`
}

// UserResponse is used for API responses (without sensitive data)
type UserResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      Role      `json:"role"`
	ShopID    int       `json:"shop_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Role:      u.Role,
		ShopID:    u.ShopID,
		CreatedAt: u.CreatedAt,
	}
}
