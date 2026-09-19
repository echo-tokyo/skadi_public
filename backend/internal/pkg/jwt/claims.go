package jwt

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

// TokenClaims represents all token claims for JWT-token.
// It includes token expiration datetime in UNIX-format and generic extra claims.
type TokenClaims[T any] struct {
	Exp         int64 `json:"exp"`
	ExtraClaims T     `json:"extra"`
}

// GetExpirationTime implements [jwt.Claims].
func (t *TokenClaims[T]) GetExpirationTime() (*jwt.NumericDate, error) {
	return &jwt.NumericDate{Time: time.Unix(t.Exp, 0)}, nil
}

// GetSubject implements [jwt.Claims].
func (t *TokenClaims[T]) GetSubject() (string, error) { return "", nil }

// GetIssuedAt implements [jwt.Claims].
func (*TokenClaims[T]) GetIssuedAt() (*jwt.NumericDate, error) { return nil, nil }

// GetNotBefore implements [jwt.Claims].
func (*TokenClaims[T]) GetNotBefore() (*jwt.NumericDate, error) { return nil, nil }

// GetIssuer implements [jwt.Claims].
func (*TokenClaims[T]) GetIssuer() (string, error) { return "", nil }

// GetAudience implements [jwt.Claims].
func (*TokenClaims[T]) GetAudience() (jwt.ClaimStrings, error) { return nil, nil }
