// Package db provides *gorm.DB for interaction with the database through GORM methods.
package db

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	_debugLevel  = logger.Info   // debug log level
	_warnLevel   = logger.Warn   // warn log level
	_errorLevel  = logger.Error  // error log level
	_silentLevel = logger.Silent // silent log level
)

// _logLevels is a matched general log levels and db log levels.
var _logLevels = map[int]logger.LogLevel{
	1: _debugLevel,  // debug
	2: _warnLevel,   // info
	3: _warnLevel,   // warn
	4: _errorLevel,  // error
	5: _silentLevel, // silent
}

// Logger is an internal interface compatible with a gorm logger.Writer.
// It is used to configure a custom *gorm.DB logger.
type Logger interface {
	Printf(format string, args ...any)
}

// dbSettings represents custom options to create *gorm.DB.
type dbSettings struct {
	customLogger    Logger
	logLevel        logger.LogLevel
	translateError  bool
	ignoreNotFound  bool
	disableColorful bool
}

// Option represents an option for DB initializing.
type Option func(*dbSettings)

// New returns a new instance of DB with connection to given DSN.
// Options can be set with "With..." funcs.
func New(dsn string, options ...Option) (*gorm.DB, error) {
	dbStorage := &dbSettings{
		customLogger:    log.Default(),
		logLevel:        _warnLevel,
		translateError:  false,
		ignoreNotFound:  false,
		disableColorful: false,
	}

	// apply all options to customize DB struct
	for _, opt := range options {
		opt(dbStorage)
	}

	gormDB, err := gorm.Open(
		withConn(dsn),
		&gorm.Config{
			// set UTC time zone
			NowFunc: func() time.Time {
				return time.Now().UTC()
			},
			Logger: logger.New(
				dbStorage.customLogger,
				logger.Config{
					LogLevel:                  dbStorage.logLevel,
					IgnoreRecordNotFoundError: dbStorage.ignoreNotFound,
					Colorful:                  !dbStorage.disableColorful,
				},
			),
			TranslateError: dbStorage.translateError,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("open db connection: %w", err)
	}

	dbStorage.customLogger.Printf("successfully connected to db")
	return gormDB, nil
}

// withConn sets connection for DB. Required.
// There is the MySQL is used as DB.
func withConn(dsn string) gorm.Dialector {
	return mysql.Open(dsn)
}

// WithLogger sets custom logger for DB.
func WithLogger(customLogger Logger) Option {
	return func(d *dbSettings) {
		d.customLogger = customLogger
	}
}

// WithLogLevel sets log level for DB.
// Accepted values: 1 - debug, 2 - info (like warn), 3 - warn, 4 - error, 5 - silent.
func WithLogLevel(logLevel int) Option {
	var level logger.LogLevel
	var ok bool
	if level, ok = _logLevels[logLevel]; !ok {
		// if log level not found in logLevels map (default level is warn)
		level = _warnLevel
	}
	return func(d *dbSettings) {
		d.logLevel = level
	}
}

// WithDebugLogLevel sets debug log level for DB.
func WithDebugLogLevel() Option {
	return func(d *dbSettings) {
		d.logLevel = _debugLevel
	}
}

// WithInfoLogLevel sets info log level for DB.
func WithInfoLogLevel() Option {
	return func(d *dbSettings) {
		d.logLevel = _warnLevel
	}
}

// WithWarnLogLevel sets warn log level for DB.
func WithWarnLogLevel() Option {
	return func(d *dbSettings) {
		d.logLevel = _warnLevel
	}
}

// WithErrorLogLevel sets error log level for DB.
func WithErrorLogLevel() Option {
	return func(d *dbSettings) {
		d.logLevel = _errorLevel
	}
}

// WithSilentLogLevel sets silent log level for DB.
func WithSilentLogLevel() Option {
	return func(d *dbSettings) {
		d.logLevel = _silentLevel
	}
}

// WithTranslateError sets translate error parameter true.
func WithTranslateError() Option {
	return func(d *dbSettings) {
		d.translateError = true
	}
}

// WithIgnoreNotFound sets ignore record not found error parameter true.
func WithIgnoreNotFound() Option {
	return func(d *dbSettings) {
		d.ignoreNotFound = true
	}
}

// WithDisableColorful sets colorful log output false.
func WithDisableColorful() Option {
	return func(d *dbSettings) {
		d.disableColorful = true
	}
}
