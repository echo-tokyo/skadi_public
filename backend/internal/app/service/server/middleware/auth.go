package middleware

import (
	"errors"
	"fmt"

	fiber "github.com/gofiber/fiber/v2"
	gojwt "github.com/golang-jwt/jwt/v5"

	"skadi/backend/config"
	"skadi/backend/internal/app/auth"
	"skadi/backend/internal/app/entity"
	"skadi/backend/internal/app/service/server/errhandler"
	"skadi/backend/internal/pkg/httperror"
	"skadi/backend/internal/pkg/jwt"
)

const (
	_accessTokenCookie  = "access"  // key to store access token in cookies
	_refreshTokenCookie = "refresh" // key to store refresh token in cookies

	_tokenCtxKey      = "token"      // key for the JWT-token string in the fiber ctx
	_userClaimsCtxKey = "userClaims" // key for the user claims in the fiber ctx
)

// JWTAccess parses access token from request header to context and validates it.
func JWTAccess(cfg *config.Config) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// parse token string from refresh cookie
		token := ctx.Cookies(_accessTokenCookie)
		if token == "" {
			return handleTokenErr(ctx, errors.New("missing"))
		}
		// parse user claims from token
		userClaims, err := parseUserClaims(cfg.Server.Auth.AccessToken.Secret, token)
		if err != nil {
			return handleTokenErr(ctx, err)
		}
		// save token string and user claims to fiber context
		ctx.Locals(_tokenCtxKey, token)
		ctx.Locals(_userClaimsCtxKey, userClaims)
		return ctx.Next()
	}
}

// JWTRefresh parses refresh token from request cookies to context and validates it.
func JWTRefresh(cfg *config.Config, authUC auth.UsecaseMiddleware) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// parse token string from refresh cookie
		token := ctx.Cookies(_refreshTokenCookie)
		if token == "" {
			return handleTokenErr(ctx, errors.New("missing"))
		}
		// parse user claims from token
		userClaims, err := parseUserClaims(cfg.Server.Auth.RefreshToken.Secret, token)
		if err != nil {
			return handleTokenErr(ctx, err)
		}

		// check token is blacklisted
		err = authUC.BlockIfBlacklist(token)
		if errors.Is(err, auth.ErrForbidden) {
			return errhandler.CustomErrorHandler(ctx, &httperror.HTTPError{
				CauseErr:   err,
				StatusCode: fiber.StatusForbidden,
				Message:    err.Error(),
			})
		}
		if err != nil {
			return errhandler.CustomErrorHandler(ctx, err)
		}

		// save token string and user claims to fiber context
		ctx.Locals(_tokenCtxKey, token)
		ctx.Locals(_userClaimsCtxKey, userClaims)
		return ctx.Next()
	}
}

// handleTokenErr handles a token error using custom error handler.
func handleTokenErr(ctx *fiber.Ctx, err error) error {
	return errhandler.CustomErrorHandler(ctx,
		fiber.NewError(fiber.StatusUnauthorized, fmt.Sprintf("token: %v", err)),
	)
}

// parseTokenWithClaims returns parsed token user claims.
func parseUserClaims(tokenSecret []byte, token string) (*entity.UserClaims, error) {
	claims := &jwt.TokenClaims[*entity.UserClaims]{}
	// parse token object with claims from token string
	_, err := gojwt.ParseWithClaims(token, claims, func(t *gojwt.Token) (any, error) {
		if t.Method != gojwt.SigningMethodHS256 {
			return nil, errors.New("invalid signature")
		}
		return tokenSecret, nil
	})
	// if token is expired
	if errors.Is(err, gojwt.ErrTokenExpired) {
		return nil, errors.New("parse: token is expired")
	}
	// other error
	if err != nil {
		return nil, fmt.Errorf("parse: %w", err)
	}
	return claims.ExtraClaims, nil
}
