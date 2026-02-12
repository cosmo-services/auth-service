package user

import "errors"

var (
	ErrUserNotFound         = errors.New("USER_NOT_FOUND")
	ErrNoPermission         = errors.New("NO_PERMISSION")
	ErrUsernameAlreadyTaken = errors.New("USERNAME_ALREADY_TAKEN")
	ErrEmailAlreadyTaken    = errors.New("EMAIL_ALREADY_TAKEN")
	ErrActivateDeleted      = errors.New("CANNOT_ACTIVATE_DELETED_USER")
	ErrAlreadyActivated     = errors.New("USER_ALREADY_ACTIVATED")
	ErrEmailNotChanged      = errors.New("EMAIL_NOT_CHANGED")
	ErrUsernameNotChanged   = errors.New("USERNAME_NOT_CHANGED")
)
