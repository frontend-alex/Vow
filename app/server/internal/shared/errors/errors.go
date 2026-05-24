package errors

import "errors"

var (
	ErrInvalidRegisterInput = errors.New("invalid register input")
	ErrEmailAlreadyExists   = errors.New("email already exists")
)
