package tokens

import "errors"

var (
	ErrInvalidToken  = errors.New("INVALID_TOKEN")
	ErrTokenExpired  = errors.New("TOKEN_EXPIRED")
	ErrTokenNotFound = errors.New("TOKEN_NOT_FOUND")
)
