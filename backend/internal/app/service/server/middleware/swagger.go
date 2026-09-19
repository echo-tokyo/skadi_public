package middleware

import (
	"github.com/gofiber/contrib/swagger"
	fiber "github.com/gofiber/fiber/v2"
)

// Swagger is a middleware for Swagger docs.
func Swagger() fiber.Handler {
	return swagger.New(swagger.Config{
		BasePath: "/api/v1/",
		FilePath: "./docs/swagger.json",
		Path:     "docs",
		Title:    "skadi API",
	})
}
