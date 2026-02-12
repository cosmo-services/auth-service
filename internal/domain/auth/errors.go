package auth

import "errors"

var ErrInvalidRefreshToken = errors.New("INVALID_REFRESH_TOKEN")
var ErrInvalidCredentials = errors.New("INVALID_CREDENTIALS")
var ErrInactiveUser = errors.New("INACTIVE_USER")
