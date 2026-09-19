package middleware

import (
	fiber "github.com/gofiber/fiber/v2"
	fibercors "github.com/gofiber/fiber/v2/middleware/cors"
)

// CORS represents a middleware with CORS setting up.
func CORS(origins, methods string, credentials bool) fiber.Handler {
	return fibercors.New(fibercors.Config{
		AllowOrigins:     origins,
		AllowMethods:     methods,
		AllowCredentials: credentials,
	})
}
