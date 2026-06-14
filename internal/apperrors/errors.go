package apperrors

import "errors"

var (
	ErrNotFound           = errors.New("not found")
	ErrHashPassword       = errors.New("failed to hash password")
	ErrDuplicateEmail     = errors.New("email already exists")
	ErrInvalidPassword    = errors.New("password must contain uppercase, lowercase, and number")
	ErrInvalidCredentials = errors.New("invalid email or password")
)
