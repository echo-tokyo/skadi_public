// Package server provides HTTP-server service.
package server

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	fiber "github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"skadi/backend/config"
	"skadi/backend/internal/app/service/server/errhandler"
	"skadi/backend/internal/app/service/server/middleware"
	"skadi/backend/internal/pkg/cache"
	"skadi/backend/internal/pkg/validator"
)

const (
	_bodyLimit    = 30 * 1024 * 1024  // control on Nginx
	_readTimeout  = 300 * time.Second // control on Nginx
	_writeTimeout = 300 * time.Second // control on Nginx
)

// Server represents a HTTP-server.
type Server struct {
	// will be closed if the service was completely started and is ready-to-use now
	ready    chan struct{}
	cfg      *config.Config
	fiberApp *fiber.App
}

//	@title						skadi API
//	@version					1.0.0
//	@description				HTTP API онлайн-сервиса для IT-школы.
//
//	@host						127.0.0.1:8000
//	@basePath					/api/v1
//	@schemes					http
//
//	@accept						json
//	@produce					json
//
//	@securityDefinitions.apiKey	JWTRefresh
//	@in							cookie
//	@name						refresh
//	@description				Refresh JWT-token. Cookie will automatic add after auth is done.
//
//	@securitydefinitions.apikey	JWTAccess
//	@in							cookie
//	@name						access
//	@description				Access JWT-token. Cookie will automatic add after auth is done.
//
// New returns a new instance of [Server].
func New(cfg *config.Config, dbStorage *gorm.DB, cacheStorage cache.Storage,
	valid validator.Validator) (*Server, error) {

	// fiber init
	server := &Server{
		ready: make(chan struct{}),
		cfg:   cfg,
		fiberApp: fiber.New(fiber.Config{
			ServerHeader:  "Skadi",
			StrictRouting: false,
			ErrorHandler:  errhandler.CustomErrorHandler,
			BodyLimit:     _bodyLimit,
			ReadTimeout:   _readTimeout,
			WriteTimeout:  _writeTimeout,
		}),
	}

	// set up base middlewares
	httpLogger := middleware.Logger(cfg.Logging.LogLevel, cfg.Logging.JSONFormat)
	if httpLogger != nil {
		server.fiberApp.Use(httpLogger)
	}
	server.fiberApp.Use(middleware.Recover())
	server.fiberApp.Use(middleware.CORS(cfg.CORS.AllowedOrigins, cfg.CORS.AllowedMethods,
		cfg.CORS.AllowCredentials))
	// register swagger docs if server is in debug mode
	if server.Debug() {
		server.fiberApp.Use(middleware.Swagger())
	}
	// register all endpoints
	server.registerEndpointsV1(cfg, dbStorage, cacheStorage, valid)

	return server, nil
}

// StartWithShutdown starts server and waits for
// context is done for gracefully shutdown server.
// This method is blocking.
func (s *Server) StartWithShutdown(ctx context.Context) error {
	slog.Info("start server...")
	defer slog.Info("stop server: ok")

	errChan := make(chan error, 1)
	defer close(errChan)
	// start server
	go func() {
		if err := s.fiberApp.Listen(":" + s.cfg.Server.Port); err != nil {
			errChan <- fmt.Errorf("server: listen: %w", err)
		}
	}()

	// notify that service is ready-to-use
	close(s.ready)
	// wait for context or server listen error
	select {
	case <-ctx.Done():
		if err := s.fiberApp.ShutdownWithTimeout(s.cfg.Server.ShutdownTimeout); err != nil {
			return fmt.Errorf("server: shutdown: %w", err)
		}
		return nil
	case err := <-errChan:
		return err
	}
}

// Ready signals that the service is ready-to-use.
func (s *Server) Ready() <-chan struct{} {
	return s.ready
}

// Debug returns true if server was started in debug mode.
func (s *Server) Debug() bool {
	return s.cfg.Logging.LogLevel == 1
}
