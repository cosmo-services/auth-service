package tokens

import "errors"

var ErrInvalidToken = errors.New("invalid verification token")
var ErrTokenExpired = errors.New("expired token")
var ErrTokenNotFound = errors.New("token not found")
