package password

import (
	"unicode"
)


type PasswordService struct {
	pswdHasher PasswordHasher
}

func NewPasswordService(pswdHasher PasswordHasher) *PasswordService {
	return &PasswordService{pswdHasher: pswdHasher}
}

func (s *PasswordService) ValidatePassword(password string) error {
	if len(password) < 8 {
		return ErrPasswordTooShort
	}

	hasUppercase := false
	hasLowercase := false
	hasDigit := false
	hasSpecial := false

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUppercase = true
		case unicode.IsLower(char):
			hasLowercase = true
		case unicode.IsDigit(char):
			hasDigit = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char) || !unicode.IsLetter(char) && !unicode.IsNumber(char):
			hasSpecial = true
		}
	}

	if !hasUppercase {
		return ErrPasswordNoUppercase
	}

	if !hasLowercase {
		return ErrPasswordNoLowercase
	}

	if !hasDigit {
		return ErrPasswordNoDigit
	}

	if !hasSpecial {
		return ErrPasswordNoSpecial
	}

	return nil
}

func (s *PasswordService) HashPassword(password string) (string, error) {
	return s.pswdHasher.HashPassword(password)
}

func (s *PasswordService) CheckPassword(password string, hash string) bool {
	return s.pswdHasher.CheckPassword(password, hash)
}
