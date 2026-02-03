package user

import "time"

type UserRepository interface {
	Create(user *User) error
	GetByID(id string) (*User, error)
	GetByEmail(email string) (*User, error)
	GetByUsername(username string) (*User, error)
	Update(user *User) error
	Delete(userID string) error
	DeleteInactiveUsers(before time.Time) error
	IsEmailAvailable(email string) (bool, error)
	IsUsernameAvailable(username string) (bool, error)
}
