// Package http/v1 is a first version of class HTTP-controller.
// It provides registers for class HTTP-routes and controller with handlers for them.
package v1

import (
	fiber "github.com/gofiber/fiber/v2"

	"skadi/backend/internal/app/entity"
	"skadi/backend/internal/app/service/server/middleware"
)

// RegisterEndpoints registers all class endpoints.
func RegisterEndpoints(router fiber.Router,
	controller *ClassController, controllerAdmin *ClassControllerAdmin,
	mwJWTAccess fiber.Handler, mwAllow middleware.AllowFunc) {

	mwAdminOnly := mwAllow(entity.Admin)

	authGroup := router.Group("/class", mwJWTAccess)
	authGroup.Post("/", mwAdminOnly, controllerAdmin.Create)
	authGroup.Get("/short", controller.ListShort)
	authGroup.Get("/", controller.ListFull)
	authGroup.Get("/:id", controller.Read)
	authGroup.Patch("/:id", mwAdminOnly, controllerAdmin.Update)
	authGroup.Delete("/:id", mwAdminOnly, controllerAdmin.Delete)
}
