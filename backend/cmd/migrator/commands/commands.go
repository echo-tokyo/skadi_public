// Package commands contains command handlers for migrator cmd binary.
package commands

import "errors"

// Validator func for cmd flags.
// Returns error if flag value is less than 0.
func positiveFlagValidator(n int) error {
	if n < 0 {
		return errors.New("flag value must be a positive number")
	}
	return nil
}
