// Package password provides functions for passwords and its hashes.
package password

import (
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

// strong password rules
var _strongPassRules = []*regexp.Regexp{
	regexp.MustCompile(`.{8,}`), // length
	regexp.MustCompile(`[a-z]`), // lower
	regexp.MustCompile(`\d`),    // digit
	regexp.MustCompile(`[A-Z]`), // upper
	regexp.MustCompile(`\W`),    // special
}

// Encode returns hashed password.
func Encode(password []byte) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		// probably, password is too long
		return nil, err
	}
	return hash, nil
}

// IsCorrect returns true if the given password equals to its hash from DB.
func IsCorrect(password, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, password)
	return err == nil
}

// Strong returns true if password is strong. Strength rules:
// 1) At least 8 letters.
// 2) At least 1 number.
// 3) At least 1 upper case.
// 4) At least 1 special character.
func Strong(password string) bool {
	for _, rule := range _strongPassRules {
		if !rule.MatchString(password) {
			return false
		}
	}
	return true
}
