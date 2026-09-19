// Package middleware provides all middlewares for HTTP-server.
package middleware

import (
	"os"
	"time"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// JSON-format for logs
const (
	_infoLevel = 2 // info log level

	_textLogFormat = `time=${time} level=INFO status=${status} method=${method} path=${path}`                       // nolint:lll // output format
	_jsonLogFormat = `{"time":"${time}","level":"INFO","status":"${status}","method":"${method}","path":"${path}"}` // nolint:lll // output format
)

// Logger is a middleware for logging all requests.
func Logger(logLevel int, jsonFormat bool) fiber.Handler {
	if logLevel > _infoLevel {
		return nil
	}

	// set formatter
	format := _textLogFormat
	disableColor := false
	if jsonFormat {
		format = _jsonLogFormat
		disableColor = true
	}

	return logger.New(logger.Config{
		TimeFormat:    time.RFC3339Nano,
		TimeZone:      "UTC",
		Format:        format + "\n",
		Output:        os.Stderr,
		DisableColors: disableColor,
	})
}
