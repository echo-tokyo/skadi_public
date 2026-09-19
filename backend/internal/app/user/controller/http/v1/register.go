// Package http/v1 is a first version of user HTTP-controller.
// It provides registers for user HTTP-routes and controller with handlers for them.
package v1

import (
	fiber "github.com/gofiber/fiber/v2"

	"skadi/backend/internal/app/entity"
	"skadi/backend/internal/app/service/server/middleware"
)

// RegisterEndpoints registers all user endpoints.
func RegisterEndpoints(router fiber.Router,
	controller *UserController, controllerAdmin *UserControllerAdmin,
	mwJWTAccess fiber.Handler, mwAllow middleware.AllowFunc) {

	mwAdminOnly := mwAllow(entity.Admin)
	mwAdminTeacher := mwAllow(entity.Admin, entity.Teacher)

	authGroup := router.Group("/user", mwJWTAccess)
	authGroup.Post("/", mwAdminOnly, controllerAdmin.Create)
	authGroup.Get("/me", controller.GetMe)
	authGroup.Put("/me/profile", controller.UpdateMeProfile)
	authGroup.Put("/me/password", controller.ChangePassword)
	authGroup.Get("/", mwAdminTeacher, controllerAdmin.List)
	authGroup.Get("/:id", mwAdminOnly, controllerAdmin.Read)
	authGroup.Put("/:id", mwAdminOnly, controllerAdmin.Update)
	authGroup.Put("/:id/password", mwAdminOnly, controllerAdmin.ChangePassword)
	authGroup.Delete("/:id", mwAdminOnly, controllerAdmin.Delete)
}
