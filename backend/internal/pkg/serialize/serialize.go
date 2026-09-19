// Package serialize contains func to deserialize any request data.
package serialize

// Parser represents a func to parse data.
type Parser func(out any) error

// Validator represents a func to validate data.
type Validator func(s any) error

// Deserialize parses and validates given data.
func Deserialize(data any, parser Parser, validator Validator) error {
	if err := parser(data); err != nil {
		return err
	}
	if err := validator(data); err != nil {
		return err
	}
	return nil
}
