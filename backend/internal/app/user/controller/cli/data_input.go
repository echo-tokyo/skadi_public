package cli

import (
	"errors"
	"fmt"
	"skadi/backend/internal/pkg/password"
	"strconv"
)

const _usernameMaxLen = 50 // max length of user username

// userBody represents a user data.
type userBody struct {
	// user username
	Username string
	// user password
	Password []byte
}

// ParseUsername parses user username from given string and validates it.
func (u *userBody) ParseUsername(raw string) error {
	if len(raw) > _usernameMaxLen {
		return errors.New("username is too long: accepted 50 (or less) simbols length")
	}
	u.Username = raw
	return nil
}

// ParsePassword parses user password from given string and validates it.
func (u *userBody) ParsePassword(raw string) error {
	if !password.Strong(raw) {
		return errors.New("weak password")
	}
	u.Password = []byte(raw)
	return nil
}

// userID represents a user ID.
type userID struct {
	// user ID
	ID int
}

// ParseID parses user ID from given string and validates it.
func (u *userID) ParseID(raw string) error {
	id, err := strconv.Atoi(raw)
	if err != nil {
		return fmt.Errorf("invalid id: %w", err)
	}
	u.ID = id
	return nil
}
