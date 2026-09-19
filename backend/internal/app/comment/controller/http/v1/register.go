// Package http/v1 is a first version of comment HTTP-controller.
// It provides registers for comment HTTP-routes and controller with handlers for them.
package v1

import (
	fiber "github.com/gofiber/fiber/v2"

	"skadi/backend/internal/app/entity"
	"skadi/backend/internal/app/service/server/middleware"
)

// RegisterEndpoints registers all comment endpoints.
func RegisterEndpoints(router fiber.Router, controller *CommentController,
	mwJWTAccess fiber.Handler, mwAllow middleware.AllowFunc) {

	mwTeacherStudent := mwAllow(entity.Teacher, entity.Student)

	authGroup := router.Group("/solution/:id/comment", mwJWTAccess, mwTeacherStudent)
	authGroup.Post("/", controller.Create)
	authGroup.Get("/", mwTeacherStudent, controller.List)
}
