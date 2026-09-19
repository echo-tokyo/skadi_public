// Package http/v1 is a first version of file HTTP-controller.
// It provides registers for file HTTP-routes and controller with handlers for them.
package v1

import (
	fiber "github.com/gofiber/fiber/v2"

	"skadi/backend/internal/app/entity"
	"skadi/backend/internal/app/service/server/middleware"
)

// RegisterEndpoints registers all file endpoints.
func RegisterEndpoints(router fiber.Router, controller *FileController,
	mwJWTAccess fiber.Handler, mwAllow middleware.AllowFunc) {

	mwTeacherStudent := mwAllow(entity.Teacher, entity.Student)

	authGroup := router.Group("/file", mwJWTAccess)
	authGroup.Get("/:id", mwTeacherStudent, controller.Download)
}
