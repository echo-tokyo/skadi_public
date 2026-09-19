// Package http/v1 is a first version of example HTTP-controller.
// It provides registers for example HTTP-routes and controller with handlers for them.
package v1

import (
	fiber "github.com/gofiber/fiber/v2"

	"skadi/backend/internal/app/entity"
	"skadi/backend/internal/app/service/server/middleware"
)

// RegisterEndpoints registers all example endpoints.
func RegisterEndpoints(router fiber.Router, controller *ExampleController,
	mwJWTAccess fiber.Handler, mwAllow middleware.AllowFunc) {

	group := router.Group("/example")
	// public
	group.Get("/free", controller.Free)
	// authenticated only
	authGroup := group.Use(mwJWTAccess)
	authGroup.Get("/private", controller.Private)
	authGroup.Get("/admin", mwAllow(entity.Admin), controller.Admin)       // admin only
	authGroup.Get("/teacher", mwAllow(entity.Teacher), controller.Teacher) // teacher only
	authGroup.Get("/student", mwAllow(entity.Student), controller.Student) // student only
}
