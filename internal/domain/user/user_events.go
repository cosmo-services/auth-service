package user

import "time"

type UserCreatedEvent struct {
	UserID    string    `json:"user_id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type UserActivateEvent struct {
	UserID      string    `json:"user_id"`
	ActivatedAt time.Time `json:"activated_at"`
}

type UserDeleteEvent struct {
	UserID    string    `json:"user_id"`
	DeletedAt time.Time `json:"deleted_at"`
}

type UserChangeEmailEvent struct {
	UserID    string    `json:"user_id"`
	NewEmail  string    `json:"new_email"`
	ChangedAt time.Time `json:"changed_at"`
}

type UserChangeUsernameEvent struct {
	UserID      string    `json:"user_id"`
	NewUsername string    `json:"new_username"`
	ChangedAt   time.Time `json:"changed_at"`
}
