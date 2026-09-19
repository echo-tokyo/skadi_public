package commands

import (
	"context"
	"fmt"

	cli "github.com/urfave/cli/v3"

	"skadi/backend/internal/pkg/migrate"
)

// Status command instance.
func NewStatus(migrateManager migrate.Migrate) *cli.Command {
	return &cli.Command{
		Name:   "status",
		Usage:  "Returns migrations status (current version and durty value)",
		Action: newStatusAction(migrateManager),
	}
}

// Handler for status command.
func newStatusAction(migrateManager migrate.Migrate) cli.ActionFunc {
	return func(_ context.Context, _ *cli.Command) error {
		version, isDirty, err := migrateManager.Status()
		if err != nil {
			return err
		}
		fmt.Printf("Version: %d | Dirty: %v \n", version, isDirty)
		return nil
	}
}
