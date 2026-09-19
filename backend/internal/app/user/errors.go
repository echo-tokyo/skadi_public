package user

import "errors"

var (
	ErrInvalidData     = errors.New("invalid data")          // code 400
	ErrUnsupportedData = errors.New("unsupported data")      // code 400
	ErrNotFound        = errors.New("record not found")      // code 404
	ErrAlreadyExists   = errors.New("record already exists") // code 409
	ErrConflict        = errors.New("conflict")              // code 409
)
