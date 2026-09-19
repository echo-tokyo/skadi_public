// Package jwt (utils) contains help-functions for JWT-token parsing.
package jwt

import (
	fiber "github.com/gofiber/fiber/v2"

	"skadi/backend/internal/app/entity"
)

const (
	_tokenCtxKey      = "token"      // key for the JWT-token string in the fiber ctx
	_userClaimsCtxKey = "userClaims" // key for the user claims in the fiber ctx
)

// ParseTokenStringFromRequest parses token string from request context.
// Token string must be saved to context by middleware.
func ParseTokenStringFromRequest(ctx *fiber.Ctx) string {
	return ctx.Locals(_tokenCtxKey).(string)
}

// ParseUserClaimsFromRequest parses user claims from request context.
// User claims must be saved to context by middleware.
func ParseUserClaimsFromRequest(ctx *fiber.Ctx) *entity.UserClaims {
	userClaims, ok := ctx.Locals(_userClaimsCtxKey).(*entity.UserClaims)
	if !ok {
		return nil
	}
	return userClaims
}
