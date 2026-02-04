package user

import (
	"errors"
)

var ErrUserNotFound = errors.New("user not found")
var ErrNoPermission = errors.New("user does not have permission")
var ErrUsernameAlreadyTaken = errors.New("username already taken")
var ErrEmailAlreadyTaken = errors.New("email already taken")
var ErrActivateDeleted = errors.New("cannot activate deleted user")
var ErrAlreadyActivated = errors.New("user is already activated")
