// Package config provides loading config data from external sources
// like env variables, yaml-files etc.
package config

import (
	"fmt"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

const (
	// server
	_defServerPort      = "8080"      // default server port
	_defShutdownTimeout = time.Minute // default server shutdown timeout

	// cors
	_defCorsAllowedOrigins   = "*"   // default allowed origins for cors
	_defCorsAllowedMethods   = "GET" // default allowed methods for cors
	_defCorsAllowCredentials = false // default allow credentials value for cors

	// auth access token
	_defAccessTTL            = 5 * time.Minute // default access token ttl (5 minutes)
	_defAccessCookiePath     = ""
	_defAccessCookieSecure   = false
	_defAccessCookieHTTPOnly = false
	_defAccessCookieSameSite = ""

	// auth refresh token
	_defRefreshTTL            = 24 * 10 * time.Hour // default refresh token ttl (10 days)
	_defRefreshCookiePath     = ""
	_defRefreshCookieSecure   = false
	_defRefreshCookieHTTPOnly = false
	_defRefreshCookieSameSite = ""

	// logging
	_defLogLevel   = 2     // default log level (info)
	_defJSONFormat = false // default log JSON-format

	// cache
	_defCacheHost = "127.0.0.1" // default cache host
	_defCachePort = "3306"      // default cache port

	// db
	_defDBUser       = "test_user"         // default db user
	_defDBName       = "test_db"           // default database name
	_defDBHost       = "127.0.0.1"         // default db host
	_defDBPort       = "3306"              // default db port
	_defMigrationSrc = "file://migrations" // default migrations source URL (dir ./migrations)

	// media
	_defTaskFileDir     = "./media/task_files"     // default dir for task files
	_defSolutionFileDir = "./media/solution_files" // default dir for solution files
)

var _dirPerms os.FileMode = 0o755 // permissions for the file dirs

type (
	Config struct {
		Server  `yaml:"server"`
		Logging `yaml:"logging"`
		Cache   `yaml:"cache"`
		DB      `yaml:"db"`
		Media   `yaml:"media"`
	}

	Server struct {
		Port            string        `yaml:"port"`
		ShutdownTimeout time.Duration `yaml:"shutdown_timeout"`
		CORS            `yaml:"cors"`
		Auth            `yaml:"auth"`
	}

	CORS struct {
		AllowedOrigins   string `yaml:"allowed_origins"`
		AllowedMethods   string `yaml:"allowed_methods"`
		AllowCredentials bool   `yaml:"allow_credentials"`
	}

	Auth struct {
		AccessToken  AccessToken  `yaml:"access_token"`
		RefreshToken RefreshToken `yaml:"refresh_token"`
	}

	AccessToken struct {
		Secret []byte        `env-required:"true" env:"ACCESS_SECRET"`
		TTL    time.Duration `yaml:"ttl"`
		Cookie Cookie        `yaml:"cookie"`
	}

	RefreshToken struct {
		Secret []byte        `env-required:"true" env:"REFRESH_SECRET"`
		TTL    time.Duration `yaml:"ttl"`
		Cookie Cookie        `yaml:"cookie"`
	}

	Cookie struct {
		Path     string `yaml:"path"`
		Secure   bool   `yaml:"secure"`
		HTTPOnly bool   `yaml:"http_only"`
		SameSite string `yaml:"same_site"`
	}

	Logging struct {
		LogLevel int `yaml:"log_level"`
		// use JSON format if true
		JSONFormat bool `yaml:"json_format"`
	}

	Cache struct {
		Host       string `yaml:"host"`
		Port       string `yaml:"port"`
		ConnString string
	}

	DB struct {
		User      string `yaml:"user"`
		Password  string `env-required:"true" env:"DB_PASSWORD"`
		Name      string `yaml:"name"`
		Host      string `yaml:"host"`
		Port      string `yaml:"port"`
		DSN       string
		Migration Migration `yaml:"migration"`
	}

	Migration struct {
		// URL to migrations source
		Src string `yaml:"src"`
		// URL to connect to DB to migrates
		DB string
	}

	Media struct {
		// dir for task files
		TaskFileDir string `yaml:"task_file_dir"`
		// dir for solution files
		SolutionFileDir string `yaml:"solution_file_dir"`
	}
)

// NewDefault returns a new instance of [Config] with default data.
func NewDefault() *Config {
	return &Config{
		Server: Server{
			Port:            _defServerPort,
			ShutdownTimeout: _defShutdownTimeout,
			CORS: CORS{
				AllowedOrigins:   _defCorsAllowedOrigins,
				AllowedMethods:   _defCorsAllowedMethods,
				AllowCredentials: _defCorsAllowCredentials,
			},
			Auth: Auth{
				AccessToken: AccessToken{
					TTL: _defAccessTTL,
					Cookie: Cookie{
						Path:     _defAccessCookiePath,
						Secure:   _defAccessCookieSecure,
						HTTPOnly: _defAccessCookieHTTPOnly,
						SameSite: _defAccessCookieSameSite,
					},
				},
				RefreshToken: RefreshToken{
					TTL: _defRefreshTTL,
					Cookie: Cookie{
						Path:     _defRefreshCookiePath,
						Secure:   _defRefreshCookieSecure,
						HTTPOnly: _defRefreshCookieHTTPOnly,
						SameSite: _defRefreshCookieSameSite,
					},
				},
			},
		},
		Logging: Logging{
			LogLevel:   _defLogLevel,
			JSONFormat: _defJSONFormat,
		},
		Cache: Cache{
			Host: _defCacheHost,
			Port: _defCachePort,
		},
		DB: DB{
			User: _defDBUser,
			Name: _defDBName,
			Host: _defDBHost,
			Port: _defDBPort,
			Migration: Migration{
				Src: _defMigrationSrc,
			},
		},
		Media: Media{
			TaskFileDir:     _defTaskFileDir,
			SolutionFileDir: _defSolutionFileDir,
		},
	}
}

// New returns a new instance of [Config] loaded from YAML-file.
func New() (*Config, error) {
	// fill with default values (for settings in yaml-file)
	cfg := NewDefault()

	// read YAML config file
	if err := cleanenv.ReadConfig("./config.yml", cfg); err != nil {
		return nil, fmt.Errorf("read yaml config file: %w", err)
	}
	// read ENV variables
	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, fmt.Errorf("read env-variables: %w", err)
	}

	// collect cache connection string
	cfg.Cache.ConnString = fmt.Sprintf("%s:%s", cfg.Cache.Host, cfg.Cache.Port)
	// collect DSN string
	dbAddr := fmt.Sprintf("tcp(%s:%s)", cfg.DB.Host, cfg.DB.Port)
	cfg.DB.DSN = fmt.Sprintf(
		"%s:%s@%s/%s?parseTime=true&timeout=5s",
		cfg.DB.User,
		cfg.DB.Password,
		dbAddr,
		cfg.DB.Name,
	)
	// collect DB connection URL string for migrate manager
	cfg.DB.Migration.DB = "mysql://" + cfg.DB.DSN

	// create task and solution files dirs
	if err := mkdirP(cfg.Media.TaskFileDir); err != nil {
		return nil, fmt.Errorf("create task file dir: %w", err)
	}
	if err := mkdirP(cfg.Media.SolutionFileDir); err != nil {
		return nil, fmt.Errorf("create solution file dir: %w", err)
	}
	return cfg, nil
}

// mkdirP creates all dirs in the given path (like mkdir -p).
func mkdirP(path string) error {
	return os.MkdirAll(path, _dirPerms)
}
