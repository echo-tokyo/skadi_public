package v1

import (
	"errors"
	"fmt"

	fiber "github.com/gofiber/fiber/v2"

	"skadi/backend/internal/app/solution"
	"skadi/backend/internal/pkg/httperror"
	"skadi/backend/internal/pkg/serialize"
	utilsjwt "skadi/backend/internal/pkg/utils/jwt"
	"skadi/backend/internal/pkg/validator"
)

// SolController represents a controller for solution routes accepted for all clients.
type SolController struct {
	valid       validator.Validator
	solUCClient solution.UsecaseClient
}

// NewController returns a new instance of [SolController].
func NewController(solUCClient solution.UsecaseClient,
	valid validator.Validator) *SolController {

	return &SolController{
		valid:       valid,
		solUCClient: solUCClient,
	}
}

// @summary		Получение решения задания по id. [Преподаватель и ученик]
// @description	Получение всех данных о решении задания (с файлами решения) с полной инфой о задании (с файлами задания) и преподе (ID и полное имя), а также со списком учеников, которые тоже выполняют это задание.
// @router			/solution/{id} [get]
// @id				solution-read
// @tags			solution
// @accept			json
// @produce		json
// @security		JWTAccess
// @param			id	path		int	true	"ID решения задания"
// @success		200	{object}	solutionOut
// @failure		401	"неверный токен (пустой, истекший или неверный формат)"
// @failure		403	"доступ запрещён"
// @failure		404	"решение задания не найдено"
func (c *SolController) Read(ctx *fiber.Ctx) error {
	// parse user claims
	userClaims := utilsjwt.ParseUserClaimsFromRequest(ctx)

	inputPath := &solutionIDPath{}
	if err := serialize.Deserialize(inputPath, ctx.ParamsParser, c.valid.Validate); err != nil {
		return err
	}

	sol, students, err := c.solUCClient.GetByIDFull(inputPath.ID, userClaims)
	if errors.Is(err, solution.ErrNotFound) {
		return &httperror.HTTPError{
			CauseErr:   err,
			StatusCode: fiber.StatusNotFound,
			Message:    "решение задания не найдено",
		}
	}
	if errors.Is(err, solution.ErrForbidden) {
		return &httperror.HTTPError{
			CauseErr:   err,
			StatusCode: fiber.StatusForbidden,
			Message:    "доступ запрещён",
		}
	}
	if err != nil {
		return fmt.Errorf("read: %w", err)
	}

	res := &solutionOut{
		Solution:      sol,
		OtherStudents: students,
	}
	return ctx.Status(fiber.StatusOK).JSON(res)
}
