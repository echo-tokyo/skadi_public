package auth

import "errors"

var (
	ErrInvalidPassword = errors.New("invalid password") // code 400
	ErrForbidden       = errors.New("forbidden")        // code 403
	ErrNotFound        = errors.New("record not found") // code 404
)
