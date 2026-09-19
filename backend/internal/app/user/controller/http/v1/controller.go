package v1

import (
	"errors"
	"fmt"

	fiber "github.com/gofiber/fiber/v2"

	"skadi/backend/internal/app/user"
	"skadi/backend/internal/pkg/httperror"
	"skadi/backend/internal/pkg/serialize"
	utilsjwt "skadi/backend/internal/pkg/utils/jwt"
	"skadi/backend/internal/pkg/validator"
)

// UserController represents a controller for user routes accepted for all clients.
type UserController struct {
	valid        validator.Validator
	userUCClient user.UsecaseClient
}

// NewController returns a new instance of [UserController].
func NewController(userUCClient user.UsecaseClient,
	valid validator.Validator) *UserController {

	return &UserController{
		valid:        valid,
		userUCClient: userUCClient,
	}
}

// @summary		Получение информации о себе.
// @description	Получение всей информации о себе (юзер, профиль).
// @router			/user/me [get]
// @id				user-me-get
// @tags			user
// @accept			json
// @produce		json
// @security		JWTAccess
// @success		200	{object}	entity.User
// @failure		401	"неверный токен (пустой, истекший или неверный формат)"
func (c *UserController) GetMe(ctx *fiber.Ctx) error {
	// parse user claims
	userClaims := utilsjwt.ParseUserClaimsFromRequest(ctx)

	userResp, err := c.userUCClient.GetByID(userClaims.ID)
	if err != nil {
		return fmt.Errorf("read: %w", err)
	}
	return ctx.Status(fiber.StatusOK).JSON(userResp)
}

// @summary		Обновление своего профиля.
// @description	Полное обновление своего профиля.
// @router			/user/me/profile [put]
// @id				user-me-profile-update
// @tags			user
// @accept			json
// @produce		json
// @security		JWTAccess
// @param			profileBody	body		profileBody	true	"profileBody"
// @success		200			{object}	entity.User
// @failure		401			"неверный токен (пустой, истекший или неверный формат)"
func (c *UserController) UpdateMeProfile(ctx *fiber.Ctx) error {
	// parse user claims
	userClaims := utilsjwt.ParseUserClaimsFromRequest(ctx)

	inputBody := &profileBody{}
	if err := serialize.Deserialize(inputBody, ctx.BodyParser, c.valid.Validate); err != nil {
		return err
	}
	profile := inputBody.ToEntityProfile()

	// update user profile
	userObj, err := c.userUCClient.UpdateProfile(userClaims.ID, profile)
	if err != nil {
		return fmt.Errorf("update profile: %w", err)
	}
	return ctx.Status(fiber.StatusOK).JSON(userObj)
}

// @summary		Смена своего пароля.
// @description	Смена своего пароля.
// @router			/user/me/password [put]
// @id				user-me-password-update
// @tags			user
// @accept			json
// @produce		json
// @security		JWTAccess
// @param			updatePasswordBody	body	updatePasswordBody	true	"updatePasswordBody"
// @success		204					"No Content"
// @failure		400					"неверный старый пароль"
// @failure		401					"неверный токен (пустой, истекший или неверный формат)"
// @failure		409					"пароли не должны совпадать"
func (c *UserController) ChangePassword(ctx *fiber.Ctx) error {
	inputBody := &updatePasswordBody{}
	if err := serialize.Deserialize(inputBody, ctx.BodyParser, c.valid.Validate); err != nil {
		return err
	}
	// parse user claims
	userClaims := utilsjwt.ParseUserClaimsFromRequest(ctx)

	// change user password
	err := c.userUCClient.ChangePasswordAsClient(userClaims.ID,
		[]byte(inputBody.OldPasswd), []byte(inputBody.NewPasswd))
	if errors.Is(err, user.ErrInvalidData) {
		return &httperror.HTTPError{
			CauseErr:   err,
			StatusCode: fiber.StatusBadRequest,
			Message:    "неверный старый пароль",
		}
	}
	if errors.Is(err, user.ErrConflict) {
		return &httperror.HTTPError{
			CauseErr:   err,
			StatusCode: fiber.StatusConflict,
			Message:    "пароли не должны совпадать",
		}
	}
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusNoContent).JSON(nil)
}
