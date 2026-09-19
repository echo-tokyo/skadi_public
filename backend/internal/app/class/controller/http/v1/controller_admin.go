package v1

import (
	"errors"
	"fmt"

	fiber "github.com/gofiber/fiber/v2"

	"skadi/backend/internal/app/class"
	"skadi/backend/internal/app/entity"
	"skadi/backend/internal/pkg/httperror"
	"skadi/backend/internal/pkg/serialize"
	"skadi/backend/internal/pkg/utils/slices"
	"skadi/backend/internal/pkg/validator"
)

// ClassControllerAdmin represents a controller for class routes accepted for admin only.
type ClassControllerAdmin struct {
	valid        validator.Validator
	classUCAdmin class.UsecaseAdmin
}

// NewControllerAdmin returns a new instance of [ClassControllerAdmin].
func NewControllerAdmin(classUCAdmin class.UsecaseAdmin,
	valid validator.Validator) *ClassControllerAdmin {

	return &ClassControllerAdmin{
		valid:        valid,
		classUCAdmin: classUCAdmin,
	}
}

// @summary		Создание новой группы. [Только админ]
// @description	Создание новой группы со всеми данными и включение в неё переданных студентов.
// @router			/class [post]
// @id				class-create
// @tags			class
// @accept			json
// @produce		json
// @security		JWTAccess
// @param			classBody	body		classBody	true	"classBody"
// @success		201			{object}	entity.Class
// @failure		400			"неверный ученик | неверный преподаватель | преподаватель не найден"
// @failure		401			"неверный токен (пустой, истекший или неверный формат)"
// @failure		409			"группа с введенным названием уже существует"
func (c *ClassControllerAdmin) Create(ctx *fiber.Ctx) error {
	inputBody := &classBody{}
	if err := serialize.Deserialize(inputBody, ctx.BodyParser, c.valid.Validate); err != nil {
		return err
	}
	classObj := inputBody.ToEntityClass()

	// create a new class
	err := c.classUCAdmin.Create(classObj, inputBody.StudentIDs)
	if errors.Is(err, class.ErrAlreadyExists) {
		return &httperror.HTTPError{
			CauseErr:   err,
			StatusCode: fiber.StatusConflict,
			Message:    "группа с введенным названием уже существует",
		}
	}
	if errors.Is(err, class.ErrNotFoundUser) {
		return &httperror.HTTPError{
			CauseErr:   err,
			StatusCode: fiber.StatusBadRequest,
			Message:    "преподаватель не найден",
		}
	}
	if errors.Is(err, class.ErrInvalidTeacher) {
		return &httperror.HTTPError{
			CauseErr:   err,
			StatusCode: fiber.StatusBadRequest,
			Message:    "неверный преподаватель",
		}
	}
	if errors.Is(err, class.ErrInvalidStud) {
		return &httperror.HTTPError{
			CauseErr:   err,
			StatusCode: fiber.StatusBadRequest,
			Message:    "неверный ученик",
		}
	}
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusCreated).JSON(classObj)
}

// @summary		Обновление группы по id. [Только админ]
// @description	Частичное обновление группы (только переданные поля) по её id.
// @router			/class/{id} [patch]
// @id				class-update
// @tags			class
// @accept			json
// @produce		json
// @security		JWTAccess
// @param			id			path		int			true	"ID группы"
// @param			updateBody	body		updateBody	true	"updateBody"
// @success		200			{object}	entity.Class
// @failure		400			"неверный ученик | неверный преподаватель | преподаватель не найден"
// @failure		401			"неверный токен (пустой, истекший или неверный формат)"
// @failure		404			"группа не найдена"
func (c *ClassControllerAdmin) Update(ctx *fiber.Ctx) error {
	inputPath := &classIDPath{}
	if err := serialize.Deserialize(inputPath, ctx.ParamsParser, c.valid.Validate); err != nil {
		return err
	}
	inputBody := &updateBody{}
	if err := serialize.Deserialize(inputBody, ctx.BodyParser, c.valid.Validate); err != nil {
		return err
	}

	// data reshaping
	newData := &entity.ClassUpdate{
		Name:            inputBody.Name,
		TeacherID:       inputBody.TeacherID,
		Schedule:        inputBody.Schedule,
		NewFullStudents: slices.DelDupls(inputBody.Students), // delete duplicates from list
	}
	classObj, err := c.classUCAdmin.Update(inputPath.ID, newData)
	if errors.Is(err, class.ErrNotFound) {
		return &httperror.HTTPError{
			CauseErr:   err,
			StatusCode: fiber.StatusConflict,
			Message:    "группа не найдена",
		}
	}
	if errors.Is(err, class.ErrNotFoundUser) {
		return &httperror.HTTPError{
			CauseErr:   err,
			StatusCode: fiber.StatusBadRequest,
			Message:    "преподаватель не найден",
		}
	}
	if errors.Is(err, class.ErrInvalidTeacher) {
		return &httperror.HTTPError{
			CauseErr:   err,
			StatusCode: fiber.StatusBadRequest,
			Message:    "неверный преподаватель",
		}
	}
	if errors.Is(err, class.ErrInvalidStud) {
		return &httperror.HTTPError{
			CauseErr:   err,
			StatusCode: fiber.StatusBadRequest,
			Message:    "неверный ученик",
		}
	}
	if err != nil {
		return err
	}
	return ctx.Status(fiber.StatusOK).JSON(classObj)
}

// @summary		Удаление группы по id. [Только админ]
// @description	Удаление группы по её id.
// @router			/class/{id} [delete]
// @id				class-delete
// @tags			class
// @accept			json
// @produce		json
// @security		JWTAccess
// @param			id	path	int	true	"ID юзера"
// @success		204	"No Content"
// @failure		401	"неверный токен (пустой, истекший или неверный формат)"
func (c *ClassControllerAdmin) Delete(ctx *fiber.Ctx) error {
	inputPath := &classIDPath{}
	if err := serialize.Deserialize(inputPath, ctx.ParamsParser, c.valid.Validate); err != nil {
		return err
	}

	// delete class
	err := c.classUCAdmin.DeleteByID(inputPath.ID)
	if err != nil {
		return fmt.Errorf("delete: %w", err)
	}
	return ctx.Status(fiber.StatusNoContent).JSON(nil)
}
