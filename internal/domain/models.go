package domain

import "time"

type User struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
	IsActive     bool   `json:"is_active"`
	IsDeleted    bool   `json:"is_deleted"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
