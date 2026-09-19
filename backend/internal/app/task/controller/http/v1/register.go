// Package http/v1 is a first version of task HTTP-controller.
// It provides registers for task HTTP-routes and controller with handlers for them.
package v1

import (
	fiber "github.com/gofiber/fiber/v2"

	"skadi/backend/internal/app/entity"
	"skadi/backend/internal/app/service/server/middleware"
)

// RegisterEndpoints registers all task endpoints.
func RegisterEndpoints(router fiber.Router, controllerTeacher *TaskControllerTeacher,
	mwJWTAccess fiber.Handler, mwAllow middleware.AllowFunc) {

	mwTeacherOnly := mwAllow(entity.Teacher)

	authGroup := router.Group("/task", mwJWTAccess)
	authGroup.Post("/", mwTeacherOnly, controllerTeacher.Create)
	authGroup.Get("/", mwTeacherOnly, controllerTeacher.List)
	authGroup.Get("/:id", mwTeacherOnly, controllerTeacher.Read)
	authGroup.Patch("/:id", mwTeacherOnly, controllerTeacher.Update)
	authGroup.Delete("/:id", mwTeacherOnly, controllerTeacher.Delete)
}
