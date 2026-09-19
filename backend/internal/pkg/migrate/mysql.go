package migrate

import (
	"errors"
	"fmt"

	gomigrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql" // mysql engine for migrate
	_ "github.com/golang-migrate/migrate/v4/source/file"    // engine for migration files
	multierror "github.com/hashicorp/go-multierror"
)

// Ensure MySQL implements interface.
var _ Migrate = (*MySQL)(nil)

// MySQL represents a [Migrate] implementation for the MySQL database.
type MySQL struct {
	mgrt *gomigrate.Migrate
}

// NewMySQL returns a new instance of [MySQL].
func NewMySQL(sourceURL, databaseURL string) (*MySQL, error) {
	mgrt, err := gomigrate.New(sourceURL, databaseURL)
	if err != nil {
		return nil, fmt.Errorf("create migrate manager: %w", err)
	}
	return &MySQL{mgrt: mgrt}, nil
}

// Status returns version and isDirty value.
func (c *MySQL) Status() (version uint, isDirty bool, err error) {
	v, d, err := c.mgrt.Version()
	if err != nil && !errors.Is(err, gomigrate.ErrNilVersion) {
		return 0, false, fmt.Errorf("migrate status: %w", err)
	}
	return v, d, nil
}

// Up applies all migrations.
func (c *MySQL) Up() error {
	err := c.mgrt.Up()
	if err != nil && !errors.Is(err, gomigrate.ErrNoChange) {
		return fmt.Errorf("migrate up: %w", err)
	}
	return nil
}

// Down rollbacks all migrations.
func (c *MySQL) Down() error {
	err := c.mgrt.Down()
	if err != nil && !errors.Is(err, gomigrate.ErrNoChange) {
		return fmt.Errorf("migrate down: %w", err)
	}
	return nil
}

// Step migrates up if n > 0, and down if n < 0.
func (c *MySQL) Step(n int) error {
	err := c.mgrt.Steps(n)
	if err != nil {
		return fmt.Errorf("migrate step: %w", err)
	}
	return nil
}

// Force sets a specific migration version.
func (c *MySQL) Force(n int) error {
	err := c.mgrt.Force(n)
	if err != nil {
		return fmt.Errorf("migrate force: %w", err)
	}
	return nil
}

// Close closes connection with database and migrations source.
func (c *MySQL) Close() error {
	var err error
	srcErr, dbErr := c.mgrt.Close()
	if srcErr != nil {
		err = multierror.Append(err, fmt.Errorf("close source: %w", srcErr))
	}
	if dbErr != nil {
		err = multierror.Append(err, fmt.Errorf("close database: %w", dbErr))
	}
	return err // err OR nil
}
