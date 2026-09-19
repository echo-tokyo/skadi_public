package v1

import (
	"errors"
	"fmt"

	fiber "github.com/gofiber/fiber/v2"

	"skadi/backend/internal/app/user"
	"skadi/backend/internal/pkg/httperror"
	"skadi/backend/internal/pkg/serialize"
	"skadi/backend/internal/pkg/validator"
)

// UserControllerAdmin represents a controller for user routes accepted for admin only.
type UserControllerAdmin struct {
	valid       validator.Validator
	userUCAdmin user.UsecaseAdmin
}

// NewControllerAdmin returns a new instance of [UserControllerAdmin].
func NewControllerAdmin(userUCAdmin user.UsecaseAdmin,
	valid validator.Validator) *UserControllerAdmin {

	return &UserControllerAdmin{
		valid:       valid,
		userUCAdmin: userUCAdmin,
	}
}

// @summary		Создание нового юзера. [Только админ]
// @description	Создание нового юзера и профиля для него со всеми данными.
// @router			/user [post]
// @id				user-create
// @tags			user
// @accept			json
// @produce		json
// @security		JWTAccess
// @param			userBody	body		userBody	true	"userBody"
// @success		201			{object}	entity.User
// @failure		400			"группа не найдена | слабый пароль"
// @failure		401			"неверный токен (пустой, истекший или неверный формат)"
// @failure		409			"пользователь с введенным логином уже существует"
func (c *UserControllerAdmin) Create(ctx *fiber.Ctx) error {
	inputBody := &userBody{}
	if err := serialize.Deserialize(inputBody, ctx.BodyParser, c.valid.Validate); err != nil {
		return err
	}
	userObj := inputBody.ToEntityUser()

	// create a new user
	err := c.userUCAdmin.CreateWithProfile(userObj)
	if errors.Is(err, user.ErrAlreadyExists) {
		return &httperror.HTTPError{
			CauseErr:   err,
			StatusCode: fiber.StatusConflict,
			Message:    "пользователь с введенным логином уже существует",
		}
	}
	if errors.Is(err, user.ErrInvalidData) {
		return &httperror.HTTPError{
			CauseErr:   err,
			StatusCode: fiber.StatusBadRequest,
			Message:    "группа не найдена",
		}
	}
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusCreated).JSON(userObj)
}

// @summary		Получение юзера по id. [Только админ]
// @description	Получение юзера со всеми данными по его id.
// @router			/user/{id} [get]
// @id				user-read
// @tags			user
// @accept			json
// @produce		json
// @security		JWTAccess
// @param			id	path		int	true	"ID юзера"
// @success		200	{object}	entity.User
// @failure		401	"неверный токен (пустой, истекший или неверный формат)"
// @failure		404	"пользователь не найден"
func (c *UserControllerAdmin) Read(ctx *fiber.Ctx) error {
	inputPath := &userIDPath{}
	if err := serialize.Deserialize(inputPath, ctx.ParamsParser, c.valid.Validate); err != nil {
		return err
	}

	userResp, err := c.userUCAdmin.GetByID(inputPath.ID)
	if errors.Is(err, user.ErrNotFound) {
		return &httperror.HTTPError{
			CauseErr:   err,
			StatusCode: fiber.StatusNotFound,
			Message:    "пользователь не найден",
		}
	}
	if err != nil {
		return fmt.Errorf("read: %w", err)
	}
	return ctx.Status(fiber.StatusOK).JSON(userResp)
}

// @summary		Обновление юзера по id. [Только админ]
// @description	Полное обновление профиля юзера, его группы (если студент) и пароля (опционально) по его id.
// @router			/user/{id} [put]
// @id				user-update
// @tags			user
// @accept			json
// @produce		json
// @security		JWTAccess
// @param			id			path		int			true	"ID юзера"
// @param			updateBody	body		updateBody	true	"updateBody"
// @success		200			{object}	entity.User
// @failure		400			"неподдерживаемые данные для роли обновляемого пользователя | группа не найдена | слабый пароль"
// @failure		401			"неверный токен (пустой, истекший или неверный формат)"
// @failure		404			"пользователь не найден"
func (c *UserControllerAdmin) Update(ctx *fiber.Ctx) error {
	inputPath := &userIDPath{}
	if err := serialize.Deserialize(inputPath, ctx.ParamsParser, c.valid.Validate); err != nil {
		return err
	}
	inputBody := &updateBody{}
	if err := serialize.Deserialize(inputBody, ctx.BodyParser, c.valid.Validate); err != nil {
		return err
	}
	oldUser := inputBody.ToEntityUser()

	// update user
	newUser, err := c.userUCAdmin.Update(inputPath.ID, oldUser)
	if errors.Is(err, user.ErrInvalidData) {
		return &httperror.HTTPError{
			CauseErr:   err,
			StatusCode: fiber.StatusBadRequest,
			Message:    "группа не найдена",
		}
	}
	if errors.Is(err, user.ErrUnsupportedData) {
		return &httperror.HTTPError{
			CauseErr:   err,
			StatusCode: fiber.StatusBadRequest,
			Message:    "неподдерживаемые данные для роли обновляемого пользователя",
		}
	}
	if errors.Is(err, user.ErrNotFound) {
		return &httperror.HTTPError{
			CauseErr:   err,
			StatusCode: fiber.StatusNotFound,
			Message:    "пользователь не найден",
		}
	}
	if err != nil {
		return fmt.Errorf("update: %w", err)
	}
	return ctx.Status(fiber.StatusOK).JSON(newUser)
}

// @summary		Удаление юзера по id. [Только админ]
// @description	Удаление юзера и его профиля по его id.
// @router			/user/{id} [delete]
// @id				user-delete
// @tags			user
// @accept			json
// @produce		json
// @security		JWTAccess
// @param			id	path	int	true	"ID юзера"
// @success		204	"No Content"
// @failure		401	"неверный токен (пустой, истекший или неверный формат)"
// @failure		404	"пользователь не найден"
func (c *UserControllerAdmin) Delete(ctx *fiber.Ctx) error {
	inputPath := &userIDPath{}
	if err := serialize.Deserialize(inputPath, ctx.ParamsParser, c.valid.Validate); err != nil {
		return err
	}

	err := c.userUCAdmin.DeleteByID(inputPath.ID)
	if errors.Is(err, user.ErrNotFound) {
		return &httperror.HTTPError{
			CauseErr:   err,
			StatusCode: fiber.StatusNotFound,
			Message:    "пользователь не найден",
		}
	}
	if err != nil {
		return fmt.Errorf("delete: %w", err)
	}
	return ctx.Status(fiber.StatusNoContent).JSON(nil)
}

// @summary		Получение списка юзеров. [Админ и преподаватель]
// @description	Получение списка юзеров со всеми данными (с настраиваемой пагинацией).
// @router			/user [get]
// @id				user-list
// @tags			user
// @accept			json
// @produce		json
// @security		JWTAccess
// @param			listUserQuery	query		listUserQuery	false	"listUserQuery"
// @success		200				{object}	listUserOut
// @failure		401				"неверный токен (пустой, истекший или неверный формат)"
func (c *UserControllerAdmin) List(ctx *fiber.Ctx) error {
	inputQuery := &listUserQuery{}
	err := serialize.Deserialize(inputQuery, ctx.QueryParser, inputQuery.Validate(c.valid))
	if err != nil {
		return err
	}
	// get pagination object OR nil
	pageParams := inputQuery.PaginationQuery.ToPagination()

	// get users
	userListResp, err := c.userUCAdmin.GetByRoles(inputQuery.Roles, inputQuery.Free,
		inputQuery.Search, pageParams)
	if err != nil {
		return fmt.Errorf("list: %w", err)
	}

	output := &listUserOut{
		Data:       userListResp,
		Pagination: pageParams,
	}
	return ctx.Status(fiber.StatusOK).JSON(output)
}

// @summary		Смена пароля юзера. [Только админ]
// @description	Смена пароля юзера по его id.
// @router			/user/{id}/password [put]
// @id				user-password-update
// @tags			user
// @accept			json
// @produce		json
// @security		JWTAccess
// @param			id						path	string					true	"ID юзера"
// @param			updatePasswordAdminBody	body	updatePasswordAdminBody	true	"updatePasswordAdminBody"
// @success		204						"No Content"
// @failure		401						"неверный токен (пустой, истекший или неверный формат)"
// @failure		404						"пользователь не найден"
func (c *UserControllerAdmin) ChangePassword(ctx *fiber.Ctx) error {
	inputPath := &userIDPath{}
	if err := serialize.Deserialize(inputPath, ctx.ParamsParser, c.valid.Validate); err != nil {
		return err
	}
	inputBody := &updatePasswordAdminBody{}
	if err := serialize.Deserialize(inputBody, ctx.BodyParser, c.valid.Validate); err != nil {
		return err
	}

	// change user password
	err := c.userUCAdmin.ChangePasswordAsAdmin(inputPath.ID, []byte(inputBody.NewPasswd))
	if errors.Is(err, user.ErrNotFound) {
		return &httperror.HTTPError{
			CauseErr:   err,
			StatusCode: fiber.StatusNotFound,
			Message:    "пользователь не найден",
		}
	}
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusNoContent).JSON(nil)
}
