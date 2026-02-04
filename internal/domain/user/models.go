package user

import "time"

type User struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	PasswordHash string `json:"-"`
	Email        string `json:"email"`
	IsActive     bool   `json:"is_active"`
	IsDeleted    bool   `json:"is_deleted"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) ValidateActivation() error {
	if u.IsDeleted {
		return ErrActivateDeleted
	}

	if u.IsActive {
		return ErrAlreadyActivated
	}

	return nil
}

func (u *User) Activate() error {
	if err := u.ValidateActivation(); err != nil {
		return err
	}

	u.IsActive = true

	return nil
}
