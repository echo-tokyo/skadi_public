// Package migrate contains migrate interface to control migrations
// and its implementations for different databases.
package migrate

// Migrate describes migration controller.
type Migrate interface {
	// Status returns version and isDirty value.
	Status() (version uint, isDirty bool, err error)
	// Up applies all migrations.
	Up() error
	// Down rollbacks all migrations.
	Down() error
	// Step migrates up if n > 0, and down if n < 0.
	Step(n int) error
	// Force sets a specific migration version.
	Force(n int) error
	// Close closes connection with database and migrations source.
	Close() error
}
