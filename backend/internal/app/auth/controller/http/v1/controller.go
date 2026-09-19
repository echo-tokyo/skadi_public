package v1

import (
	"errors"
	"fmt"

	fiber "github.com/gofiber/fiber/v2"

	"skadi/backend/config"
	"skadi/backend/internal/app/auth"
	"skadi/backend/internal/app/user"
	"skadi/backend/internal/pkg/cookie"
	"skadi/backend/internal/pkg/httperror"
	"skadi/backend/internal/pkg/serialize"
	utilsjwt "skadi/backend/internal/pkg/utils/jwt"
	"skadi/backend/internal/pkg/validator"
)

const _accessTokenCookie = "access"   // key to store access token in cookies
const _refreshTokenCookie = "refresh" // key to store refresh token in cookies

// AuthController represents a controller for all auth routes.
type AuthController struct {
	valid                validator.Validator
	authUCClient         auth.UsecaseClient
	accessCookieBuilder  *cookie.Builder
	refreshCookieBuilder *cookie.Builder
}

// NewController returns a new instance of [AuthController].
func NewController(cfg *config.Config, authUCClient auth.UsecaseClient,
	valid validator.Validator) *AuthController {

	return &AuthController{
		valid:        valid,
		authUCClient: authUCClient,
		accessCookieBuilder: cookie.NewBuilder(cfg.Auth.AccessToken.TTL,
			cookie.WithPath(cfg.Auth.AccessToken.Cookie.Path),
			cookie.WithSecure(cfg.Auth.AccessToken.Cookie.Secure),
			cookie.WithHTTPOnly(cfg.Auth.AccessToken.Cookie.HTTPOnly),
			cookie.WithSameSite(cfg.Auth.AccessToken.Cookie.SameSite),
		),
		refreshCookieBuilder: cookie.NewBuilder(cfg.Auth.RefreshToken.TTL,
			cookie.WithPath(cfg.Auth.RefreshToken.Cookie.Path),
			cookie.WithSecure(cfg.Auth.RefreshToken.Cookie.Secure),
			cookie.WithHTTPOnly(cfg.Auth.RefreshToken.Cookie.HTTPOnly),
			cookie.WithSameSite(cfg.Auth.RefreshToken.Cookie.SameSite),
		),
	}
}

// @summary		Вход для юзера.
// @description	Вход для существующего юзера по логину и паролю.
// @router			/auth/login [post]
// @id				auth-login
// @tags			auth
// @accept			json
// @produce		json
// @param			authBody	body		authBody	true	"authBody"
// @success		200			{object}	entity.User
// @failure		400			"неверный логин или пароль"
func (c *AuthController) LogIn(ctx *fiber.Ctx) error {
	inputBody := &authBody{}
	if err := serialize.Deserialize(inputBody, ctx.BodyParser, c.valid.Validate); err != nil {
		return err
	}

	// log in existing user
	userWithToken, err := c.authUCClient.LogIn(inputBody.Username, []byte(inputBody.Password))
	if errors.Is(err, auth.ErrInvalidPassword) || errors.Is(err, user.ErrNotFound) {
		return &httperror.HTTPError{
			CauseErr:   err,
			StatusCode: fiber.StatusBadRequest,
			Message:    "неверный логин или пароль",
		}
	}
	if err != nil {
		return err
	}
	// add access and refresh tokens to cookies
	ctx.Cookie(c.accessCookieBuilder.Create(_accessTokenCookie, userWithToken.Token.Access))
	ctx.Cookie(c.refreshCookieBuilder.Create(_refreshTokenCookie, userWithToken.Token.Refresh))
	return ctx.Status(fiber.StatusOK).JSON(userWithToken.User)
}

// @summary		Получение access токена.
// @description	Получение нового access токена по данному refresh токену из cookie.
// @router			/auth/private/obtain [post]
// @id				auth-private-obtain
// @tags			auth
// @produce		json
// @security		JWTRefresh
// @success		204 "No Content"
// @failure		401	"неверный токен (пустой, истекший или неверный формат)"
// @failure		403	"токен в черном списке"
func (c *AuthController) Obtain(ctx *fiber.Ctx) error {
	// parse user claims
	userClaims := utilsjwt.ParseUserClaimsFromRequest(ctx)
	// generate access token
	token, err := c.authUCClient.ObtainAccess(userClaims)
	if err != nil {
		return fmt.Errorf("obtain access token: %w", err)
	}
	// add access token to cookies
	ctx.Cookie(c.accessCookieBuilder.Create(_accessTokenCookie, token.Access))
	return ctx.Status(fiber.StatusNoContent).JSON(nil)
}

// @summary		Выход юзера.
// @description	Выход юзера (помещение refresh токена юзера в черный список).
// @router			/auth/private/logout [post]
// @id				auth-private-logout
// @tags			auth
// @security		JWTRefresh
// @success		204	"No Content"
// @failure		401	"неверный токен (пустой, истекший или неверный формат)"
// @failure		403	"токен в черном списке"
func (c *AuthController) LogOut(ctx *fiber.Ctx) error {
	// parse access token
	token := utilsjwt.ParseTokenStringFromRequest(ctx)

	// put token to blacklist
	if err := c.authUCClient.LogOut(token); err != nil {
		return err
	}
	// remove refresh token from cookies
	ctx.Cookie(c.accessCookieBuilder.Clear(_accessTokenCookie))
	ctx.Cookie(c.refreshCookieBuilder.Clear(_refreshTokenCookie))
	return ctx.Status(fiber.StatusNoContent).Send(nil)
}
