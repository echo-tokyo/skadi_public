package file

import "errors"

var (
	ErrForbidden = errors.New("forbidden")        // code 403
	ErrNotFound  = errors.New("record not found") // code 404
)
