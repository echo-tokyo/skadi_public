// Package http/v1 is a first version of solution HTTP-controller.
// It provides registers for task HTTP-routes and controller with handlers for them.
package v1

import (
	fiber "github.com/gofiber/fiber/v2"

	"skadi/backend/internal/app/entity"
	"skadi/backend/internal/app/service/server/middleware"
)

// RegisterEndpoints registers all solution endpoints.
func RegisterEndpoints(router fiber.Router, controller *SolController,
	controllerStudent *SolControllerStudent, controllerTeacher *SolControllerTeacher,
	mwJWTAccess fiber.Handler, mwAllow middleware.AllowFunc) {

	mwTeacherOnly := mwAllow(entity.Teacher)
	mwStudentOnly := mwAllow(entity.Student)
	mwTeacherStudent := mwAllow(entity.Teacher, entity.Student)

	authGroup := router.Group("/solution", mwJWTAccess)
	authGroup.Get("/for-teacher", mwTeacherOnly, controllerTeacher.List)
	authGroup.Get("/for-student", mwStudentOnly, controllerStudent.List)
	authGroup.Patch("/for-teacher/:id", mwTeacherOnly, controllerTeacher.Update)
	authGroup.Patch("/for-student/:id", mwStudentOnly, controllerStudent.Update)
	authGroup.Get("/:id", mwTeacherStudent, controller.Read)
	authGroup.Delete("/:id", mwTeacherOnly, controllerTeacher.Delete)
}
