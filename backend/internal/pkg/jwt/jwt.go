// Package jwt provides JWT-token builder.
// Also it contains functions to parse JWT-tokens.
package jwt

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

const (
	_defaultAccessTTL  = 5 * time.Minute     // default access token ttl  (5 minutes)
	_defaultRefreshTTL = 24 * 10 * time.Hour // default refresh token ttl (10 days)
)

// Builder represents a JWT-token builder for obtaining tokens.
type Builder struct {
	accessSecret  []byte
	refreshSecret []byte
	accessTTL     time.Duration
	refreshTTL    time.Duration
}

// Option represents an option for [Builder] initializing.
type Option func(*Builder)

// NewBuilder returns a new instance of [Builder].
func NewBuilder(accessSecret, refreshSecret []byte, options ...Option) *Builder {
	builder := &Builder{
		accessSecret:  accessSecret,
		refreshSecret: refreshSecret,
		accessTTL:     _defaultAccessTTL,
		refreshTTL:    _defaultRefreshTTL,
	}

	// apply all options to customize Builder
	for _, opt := range options {
		opt(builder)
	}
	return builder
}

// WithAccessTTL sets custom access token ttl.
func WithAccessTTL(exp time.Duration) Option {
	return func(b *Builder) {
		b.accessTTL = exp
	}
}

// WithRefreshTTL sets custom refresh token ttl.
func WithRefreshTTL(exp time.Duration) Option {
	return func(b *Builder) {
		b.refreshTTL = exp
	}
}

// ObtainAccess obtains a new access token with given user claims and returns it.
func (b *Builder) ObtainAccess(claims any) (string, error) {
	return obtain(b.accessTTL, b.accessSecret, claims)
}

// ObtainRefresh obtains a new refresh token with given user claims and returns it.
func (b *Builder) ObtainRefresh(claims any) (string, error) {
	return obtain(b.refreshTTL, b.refreshSecret, claims)
}

// obtain obtains a new token with given user claims and returns it.
func obtain(ttl time.Duration, secret []byte, claims any) (string, error) {
	tokenClaims := &TokenClaims[any]{
		ExtraClaims: claims,
		Exp:         time.Now().UTC().Add(ttl).Unix(),
	}
	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	tokenStr, err := tokenObj.SignedString(secret)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}
