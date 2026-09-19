// Package auth contains all repos, usecases and controllers for auth.
// Sub-package repo contains RepoCache implementation.
// Sub-package usecase contains UsecaseClient and UsecaseMiddleware implementations.
package auth

import (
	"skadi/backend/internal/app/entity"
)

// UsecaseClient describes all auth usecases for client.
type UsecaseClient interface {
	// LogIn returns authenticated user with token pair (access and refresh).
	// Password is a raw (not hashed) password.
	LogIn(username string, password []byte) (*entity.UserWithToken, error)
	// LogOut sets given refresh token to blacklist.
	LogOut(refreshToken string) error
	// ObtainAccess obtains a new access token using given user data.
	ObtainAccess(userClaims *entity.UserClaims) (*entity.Token, error)
}

// UsecaseMiddleware describes all auth usecases for middlewares.
type UsecaseMiddleware interface {
	// BlockIfBlacklist returns error if the given token is in blacklist.
	BlockIfBlacklist(token string) error
}
