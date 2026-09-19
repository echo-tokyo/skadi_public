// Package validator provides interface to validate struct data by tags.
package validator

import "errors"

var ErrValidate = errors.New("validate") // validate data error

// Validator describes tool to validate any struct.
type Validator interface {
	// Validate validates given struct s (s is a pointer to struct).
	Validate(s any) error
}
