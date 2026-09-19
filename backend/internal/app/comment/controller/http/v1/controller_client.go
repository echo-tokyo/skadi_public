package v1

import (
	"errors"

	fiber "github.com/gofiber/fiber/v2"

	"skadi/backend/internal/app/comment"
	"skadi/backend/internal/app/entity"
	"skadi/backend/internal/app/solution"
	"skadi/backend/internal/pkg/httperror"
	"skadi/backend/internal/pkg/serialize"
	utilsjwt "skadi/backend/internal/pkg/utils/jwt"
	"skadi/backend/internal/pkg/validator"
)

// CommentController represents a controller for the comment routes.
type CommentController struct {
	valid           validator.Validator
	commentUCClient comment.UsecaseClient
}

// NewController returns a new instance of [CommentController].
func NewController(commentUCClient comment.UsecaseClient,
	valid validator.Validator) *CommentController {

	return &CommentController{
		valid:           valid,
		commentUCClient: commentUCClient,
	}
}

// @summary		Создание комментария под решением задания. [Преподаватель и ученик]
// @description	Создание комментария от лица преподавателя или ученика для данного решения задания.
// @router			/solution/{id}/comment [post]
// @id				comment-create
// @tags			comment
// @accept			json
// @produce		json
// @security		JWTAccess
// @param			id			path		int			true	"ID решения задания"
// @param			commentBody	body		commentBody	true	"commentBody"
// @success		201			{object}	entity.Comment
// @failure		400			"решение задания не найдено"
// @failure		401			"неверный токен (пустой, истекший или неверный формат)"
// @failure		403			"доступ запрещён"
func (c *CommentController) Create(ctx *fiber.Ctx) error {
	// parse user claims
	userClaims := utilsjwt.ParseUserClaimsFromRequest(ctx)

	inputPath := &solutionIDPath{}
	if err := serialize.Deserialize(inputPath, ctx.ParamsParser, c.valid.Validate); err != nil {
		return err
	}
	inputBody := &commentBody{}
	if err := serialize.Deserialize(inputBody, ctx.BodyParser, c.valid.Validate); err != nil {
		return err
	}

	// data reshaping
	commentObj := &entity.Comment{
		SolutionID: inputPath.ID,
		Role:       userClaims.Role,
		Message:    inputBody.Message,
	}
	// create a new comment
	err := c.commentUCClient.Create(userClaims, commentObj)
	if errors.Is(err, solution.ErrNotFound) {
		return &httperror.HTTPError{
			CauseErr:   err,
			StatusCode: fiber.StatusBadRequest,
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
		return err
	}
	return ctx.Status(fiber.StatusCreated).JSON(commentObj)
}

// @summary		Получение комментариев под решением задания. [Преподаватель и ученик]
// @description	Получение списка комментариев для данного решения задания.
// @router			/solution/{id}/comment [get]
// @id				comment-list
// @tags			comment
// @accept			json
// @produce		json
// @security		JWTAccess
// @param			id					path		int					true	"ID решения задания"
// @param			listCommentQuery	query		listCommentQuery	false	"listCommentQuery"
// @success		200					{object}	listCommentOut
// @failure		400					"решение задания не найдено"
// @failure		401					"неверный токен (пустой, истекший или неверный формат)"
// @failure		403					"доступ запрещён"
func (c *CommentController) List(ctx *fiber.Ctx) error {
	// parse user claims
	userClaims := utilsjwt.ParseUserClaimsFromRequest(ctx)

	inputPath := &solutionIDPath{}
	if err := serialize.Deserialize(inputPath, ctx.ParamsParser, c.valid.Validate); err != nil {
		return err
	}
	inputQuery := &listCommentQuery{}
	if err := serialize.Deserialize(inputQuery, ctx.QueryParser, c.valid.Validate); err != nil {
		return err
	}
	// get pagination object OR nil
	pageParams := inputQuery.PaginationQuery.ToPagination()

	// get comments
	commentListResp, err := c.commentUCClient.List(inputPath.ID, userClaims, pageParams)
	if errors.Is(err, solution.ErrNotFound) {
		return &httperror.HTTPError{
			CauseErr:   err,
			StatusCode: fiber.StatusBadRequest,
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
		return err
	}
	output := &listCommentOut{
		Data:       commentListResp,
		Pagination: pageParams,
	}
	return ctx.Status(fiber.StatusOK).JSON(output)
}
