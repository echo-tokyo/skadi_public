package class

import "errors"

var (
	ErrInvalidStud    = errors.New("invalid student")       // code 400
	ErrInvalidTeacher = errors.New("invalid teacher")       // code 400
	ErrNotFound       = errors.New("record not found")      // code 404
	ErrNotFoundUser   = errors.New("user not found")        // code 404
	ErrAlreadyExists  = errors.New("record already exists") // code 409
)
