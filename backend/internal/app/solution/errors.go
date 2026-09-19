package solution

import "errors"

var (
	ErrInvalidData     = errors.New("invalid data")     // code 400
	ErrUnsupportedData = errors.New("unsupported data") // code 400
	ErrForbidden       = errors.New("forbidden")        // code 403
	ErrNotFound        = errors.New("record not found") // code 404
)
