package v1

import (
	"errors"
	"fmt"

	fiber "github.com/gofiber/fiber/v2"

	"skadi/backend/internal/app/file"
	"skadi/backend/internal/pkg/httperror"
	"skadi/backend/internal/pkg/serialize"
	utilsjwt "skadi/backend/internal/pkg/utils/jwt"
	"skadi/backend/internal/pkg/validator"
)

// FileController represents a controller for file routes.
type FileController struct {
	valid        validator.Validator
	fileUCClient file.UsecaseClient
}

// NewController returns a new instance of [FileController].
func NewController(fileUCClient file.UsecaseClient,
	valid validator.Validator) *FileController {

	return &FileController{
		valid:        valid,
		fileUCClient: fileUCClient,
	}
}

// @summary		Загрузка файла по id. [Преподаватель и ученик]
// @description	Загрузка файла по его id.
// @router			/file/{id} [get]
// @id				file-download
// @tags			file
// @accept			json
// @security		JWTAccess
// @param			id	path	int	true	"ID файла"
// @success		200	"файл"
// @failure		401	"неверный токен (пустой, истекший или неверный формат)"
// @failure		403	"доступ запрещён"
// @failure		404	"файл не найден"
func (c *FileController) Download(ctx *fiber.Ctx) error {
	// parse user claims
	userClaims := utilsjwt.ParseUserClaimsFromRequest(ctx)

	inputPath := &fileIDPath{}
	if err := serialize.Deserialize(inputPath, ctx.ParamsParser, c.valid.Validate); err != nil {
		return err
	}

	// get file from DB
	fileObj, err := c.fileUCClient.GetByID(inputPath.ID, userClaims)
	if errors.Is(err, file.ErrForbidden) {
		return &httperror.HTTPError{
			CauseErr:   err,
			StatusCode: fiber.StatusForbidden,
			Message:    "доступ запрещён",
		}
	}
	if errors.Is(err, file.ErrNotFound) {
		return &httperror.HTTPError{
			CauseErr:   err,
			StatusCode: fiber.StatusNotFound,
			Message:    "файл не найден",
		}
	}
	if err != nil {
		return err
	}

	// set filename
	ctx.Set("Content-Disposition", fmt.Sprintf("attachment; filename=%q", fileObj.Name))
	// set MIME-type
	ctx.Set("Content-Type", fileObj.MimeType)
	// send file
	return ctx.SendFile(fileObj.Path)
}
