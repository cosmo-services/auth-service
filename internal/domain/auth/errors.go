package auth

import "errors"

var ErrInvalidRefreshToken = errors.New("invalid refresh token")
var ErrInvalidCredentials = errors.New("invalid credentials")
var ErrInactiveUser = errors.New("user account is not acitvated")
