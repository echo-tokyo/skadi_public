// Package errhandler contains custom error handler func for fiber server.
package errhandler

import (
	"errors"
	"log/slog"

	fiber "github.com/gofiber/fiber/v2"

	"skadi/backend/internal/pkg/httperror"
	"skadi/backend/internal/pkg/validator"
)

// errorResponse represents a response with user-friendly error message.
type errorResponse struct {
	source     string `json:"-"`
	StatusCode int    `json:"-"`
	Message    string `json:"message"`
}

// CustomErrorHandler is a custom error handler for server.
func CustomErrorHandler(ctx *fiber.Ctx, err error) error {
	var errResp errorResponse

	var fiberErr *fiber.Error
	var httpErr *httperror.HTTPError

	switch {
	// *fiber.Error error
	case errors.As(err, &fiberErr):
		errResp.source = "fiber"
		errResp.StatusCode = fiberErr.Code
		errResp.Message = fiberErr.Message
		// change default message for 403 error
		if errResp.StatusCode == fiber.StatusForbidden {
			errResp.Message = "запрещено"
		}
		// change default message for 404 error
		if errResp.StatusCode == fiber.StatusNotFound {
			errResp.Message = "ресурс не найден"
		}
	// *httperror.HTTPError error
	case errors.As(err, &httpErr):
		errResp.source = "httperror"
		errResp.StatusCode = httpErr.StatusCode
		errResp.Message = httpErr.Message
	// validate error
	case errors.Is(err, validator.ErrValidate):
		errResp.StatusCode = fiber.StatusBadRequest
		errResp.Message = err.Error()
	// unknown error
	default:
		errResp.source = err.Error()
		errResp.StatusCode = fiber.StatusInternalServerError
		errResp.Message = "internal server error"
	}

	// skip error log for validation errors
	if errResp.source != "" {
		slog.Error(errResp.source,
			"code", errResp.StatusCode,
			"cause", err)
	}
	// send error response
	return ctx.Status(errResp.StatusCode).JSON(errResp)
}
