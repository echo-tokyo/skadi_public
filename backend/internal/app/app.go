// Package app provides object with Run method to start full application.
package app

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"skadi/backend/config"
	"skadi/backend/internal/app/service/cmdmanager"
	"skadi/backend/internal/app/service/server"

	"skadi/backend/internal/pkg/cache"
	"skadi/backend/internal/pkg/db"
	"skadi/backend/internal/pkg/logger"
	"skadi/backend/internal/pkg/validator"
)

// Ensure server implements interface.
var _ Service = (*server.Server)(nil)

// Service describes an app service.
type Service interface {
	// StartWithShutdown starts service and wait for context cancellation to shutdown it.
	StartWithShutdown(ctx context.Context) error
	// Ready returns true if service was completely started and is ready-to-use now.
	Ready() <-chan struct{}
}

// App represents full application with services.
type App struct {
	cfg      *config.Config
	services []Service
}

// NewServer returns a new instance of [App] with services to start HTTP-server.
func NewServer() (*App, error) {
	// load config
	cfg, err := config.New()
	if err != nil {
		return nil, fmt.Errorf("create config: %w", err)
	}

	// setup logger
	logOpts := []logger.SlogOption{logger.WithLogLevel(cfg.Logging.LogLevel)}
	if cfg.Logging.JSONFormat {
		logOpts = append(logOpts, logger.WithJSONFormat())
	}
	logger.InitSlog(logOpts...)
	// print out config params in debug log mode
	slog.Debug("current config", "config", cfg)

	// init validator
	valid, err := validator.NewRuTagValidator()
	if err != nil {
		return nil, fmt.Errorf("validator: %w", err)
	}

	// connect to DB
	dbStorage, err := db.New(cfg.DB.DSN,
		db.WithTranslateError(),
		db.WithIgnoreNotFound(),
		db.WithDisableColorful(),
		db.WithLogLevel(cfg.Logging.LogLevel),
		db.WithLogger(log.Default()))
	if err != nil {
		return nil, fmt.Errorf("db: %w", err)
	}

	// init cache storage
	cacheStorage, err := cache.NewRedis(cfg.Cache.ConnString)
	if err != nil {
		return nil, fmt.Errorf("redis cache: %w", err)
	}

	// init server service
	srv, err := server.New(cfg, dbStorage, cacheStorage, valid)
	if err != nil {
		return nil, fmt.Errorf("create server service: %w", err)
	}

	return &App{
		cfg:      cfg,
		services: []Service{srv},
	}, nil
}

// NewCmdManager returns a new instance of [App] with services to start command line manager.
func NewCmdManager() (*App, error) {
	// load config
	cfg, err := config.New()
	if err != nil {
		return nil, fmt.Errorf("create config: %w", err)
	}

	// setup logger (text format and error level)
	logger.InitSlog(logger.WithErrorLevel())

	// connect to DB
	dbStorage, err := db.New(cfg.DB.DSN,
		db.WithTranslateError(),
		db.WithIgnoreNotFound(),
		db.WithDisableColorful(),
		db.WithLogLevel(cfg.Logging.LogLevel),
		db.WithLogger(log.Default()))
	if err != nil {
		return nil, fmt.Errorf("db: %w", err)
	}

	// init cmd manager service
	manager, err := cmdmanager.New(cfg, dbStorage)
	if err != nil {
		return nil, fmt.Errorf("create cmd manager service: %w", err)
	}

	return &App{
		cfg:      cfg,
		services: []Service{manager},
	}, nil
}

// Run starts all services. This function is blocking.
// It waits for os signal to gracefully shutdown all services.
// Or it waits for fall down one of the services and stops other services.
func (a *App) Run() error {
	var appErr error

	// ctx for app services
	appContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	// handle shutdown process signals
	quitSig := make(chan os.Signal, 1)
	signal.Notify(quitSig,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	// start all services
	var wgRunning sync.WaitGroup
	var wgReady sync.WaitGroup
	serviceErr := make(chan error, 1)
	serviceSuccess := make(chan struct{}, 1)
	for _, service := range a.services {
		wgRunning.Add(1)
		wgReady.Add(1)
		// start service
		go func() {
			defer wgRunning.Done()
			if err := service.StartWithShutdown(appContext); err != nil {
				serviceErr <- err
			}
			serviceSuccess <- struct{}{}
		}()
		// wait for service is ready
		go func() {
			defer wgReady.Done()
			<-service.Ready()
		}()
	}
	// wait for all services until they are ready
	wgReady.Wait()
	slog.Warn("all services were started successfully")

	select {
	case handledSignal := <-quitSig:
		cancel()
		slog.Warn("got os signal. Shutdown services...", "signal", handledSignal.String())
	case err := <-serviceErr:
		cancel()
		appErr = fmt.Errorf("service: %w", err)
		slog.Warn("one of the services fell down. Shutdown other services...")
	case <-serviceSuccess:
		cancel()
		slog.Warn("one of the services has finished. Shutdown app...")
	case <-appContext.Done():
		appErr = appContext.Err()
		slog.Warn("context canceled. Shutdown app...")
	}

	// wait for all services
	wgRunning.Wait()
	slog.Warn("all services was stopped. Shutdown app...")
	return appErr
}
