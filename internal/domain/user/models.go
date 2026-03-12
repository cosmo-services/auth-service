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
		return ErrUserDeleted
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

func (u *User) Deactivate() error {
	if u.IsDeleted {
		return ErrUserDeleted
	}

	u.IsActive = false

	return nil
}

func (u *User) ChangeEmail(newEmail string) error {
	if u.Email == newEmail {
		return ErrEmailNotChanged
	}
	u.Email = newEmail
	return nil
}

func (u *User) ChangePassword(newPasswordHash string) error {
	u.PasswordHash = newPasswordHash
	return nil
}

func (u *User) ChangeUsername(newUsername string) error {
	if u.Username == newUsername {
		return ErrUsernameNotChanged
	}
	u.Username = newUsername
	return nil
}
