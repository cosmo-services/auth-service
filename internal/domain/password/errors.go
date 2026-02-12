package password

import "errors"

var (
	ErrPasswordTooShort    = errors.New("PASSWORD_TOO_SHORT")
	ErrPasswordNoUppercase = errors.New("PASSWORD_MISSING_UPPERCASE")
	ErrPasswordNoLowercase = errors.New("PASSWORD_MISSING_LOWERCASE")
	ErrPasswordNoDigit     = errors.New("PASSWORD_MISSING_DIGIT")
	ErrPasswordNoSpecial   = errors.New("PASSWORD_MISSING_SPECIAL")
)
