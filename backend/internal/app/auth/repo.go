package auth

import (
	"time"
)

// RepositoryCache describes all cache methods for auth.
type RepositoryCache interface {
	// Set token to blacklist expiring after exp duration.
	SetTokenToBlacklist(token string, exp time.Duration) error
	// TokenIsBlacklisted returns true, if given token is blacklisted.
	TokenIsBlacklisted(token string) (bool, error)
}
