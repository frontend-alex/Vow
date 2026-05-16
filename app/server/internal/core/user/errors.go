package user

import "errors"

var (
	ErrNotFound        = errors.New("user not found")
	ErrEmailTaken      = errors.New("email already taken")
	ErrInvalidEmail    = errors.New("invalid email")
	ErrInvalidPassword = errors.New("invalid password")
)
