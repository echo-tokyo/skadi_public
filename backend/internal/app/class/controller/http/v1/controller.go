package v1

import (
	"errors"
	"fmt"

	fiber "github.com/gofiber/fiber/v2"

	"skadi/backend/internal/app/class"
	"skadi/backend/internal/app/entity"
	"skadi/backend/internal/pkg/httperror"
	"skadi/backend/internal/pkg/serialize"
	"skadi/backend/internal/pkg/validator"
)

// ClassController represents a controller for class routes accepted for all clients.
type ClassController struct {
	valid         validator.Validator
	classUCClient class.UsecaseClient
}

// NewController returns a new instance of [ClassController].
func NewController(classUCClient class.UsecaseClient,
	valid validator.Validator) *ClassController {

	return &ClassController{
		valid:         valid,
		classUCClient: classUCClient,
	}
}

// @summary		Получение группы по id.
// @description	Получение всех данных о группе с инфой о преподе и учениках (ID и полные имена).
// @router			/class/{id} [get]
// @id				class-read
// @tags			class
// @accept			json
// @produce		json
// @security		JWTAccess
// @param			id	path		int	true	"ID группы"
// @success		200	{object}	entity.Class
// @failure		401	"неверный токен (пустой, истекший или неверный формат)"
// @failure		404	"группа не найдена"
func (c *ClassController) Read(ctx *fiber.Ctx) error {
	inputPath := &classIDPath{}
	if err := serialize.Deserialize(inputPath, ctx.ParamsParser, c.valid.Validate); err != nil {
		return err
	}

	classResp, err := c.classUCClient.GetByID(inputPath.ID)
	if errors.Is(err, class.ErrNotFound) {
		return &httperror.HTTPError{
			CauseErr:   err,
			StatusCode: fiber.StatusNotFound,
			Message:    "группа не найдена",
		}
	}
	if err != nil {
		return fmt.Errorf("read: %w", err)
	}
	return ctx.Status(fiber.StatusOK).JSON(classResp)
}

// @summary		Получение списка групп (кратко).
// @description	Получение списка групп (только ID и названия).
// @router			/class/short [get]
// @id				class-list-short
// @tags			class
// @accept			json
// @produce		json
// @security		JWTAccess
// @param			listClassQuery	query	listClassQuery	false	"listClassQuery"
// @success		200				{array}	entity.Class
// @failure		401				"неверный токен (пустой, истекший или неверный формат)"
func (c *ClassController) ListShort(ctx *fiber.Ctx) error {
	return getClassList(ctx, c.valid, c.classUCClient.ListShort)
}

// @summary		Получение списка групп.
// @description	Получение списка групп со всеми данными (ID и имена препода и учеников).
// @router			/class [get]
// @id				class-list
// @tags			class
// @accept			json
// @produce		json
// @security		JWTAccess
// @param			listClassQuery	query		listClassQuery	false	"listClassQuery"
// @success		200				{object}	listClassOut
// @failure		401				"неверный токен (пустой, истекший или неверный формат)"
func (c *ClassController) ListFull(ctx *fiber.Ctx) error {
	return getClassList(ctx, c.valid, c.classUCClient.ListFull)
}

// getListFunc represents a func to get classes list.
type getListFunc func(search string, page *entity.Pagination) ([]entity.Class, error)

// getClassList handles short/full class list endpoints.
func getClassList(ctx *fiber.Ctx, valid validator.Validator, getList getListFunc) error {
	inputQuery := &listClassQuery{}
	if err := serialize.Deserialize(inputQuery, ctx.QueryParser, valid.Validate); err != nil {
		return err
	}
	// get pagination object OR nil
	pageParams := inputQuery.PaginationQuery.ToPagination()

	// get classes
	classListResp, err := getList(inputQuery.Search, pageParams)
	if err != nil {
		return fmt.Errorf("list: %w", err)
	}
	output := &listClassOut{
		Data:       classListResp,
		Pagination: pageParams,
	}
	return ctx.Status(fiber.StatusOK).JSON(output)
}
