package commands

import (
	"context"
	"fmt"

	cli "github.com/urfave/cli/v3"

	"skadi/backend/internal/pkg/migrate"
)

// Force command instance.
func NewForce(migrateManager migrate.Migrate) *cli.Command {
	return &cli.Command{
		Name:   "force",
		Usage:  "Set a specific migration version",
		Action: newForceAction(migrateManager),
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:      "version",
				Aliases:   []string{"v", "n"},
				Value:     1,
				Usage:     "Version of migration to set",
				Validator: positiveFlagValidator,
			},
		},
	}
}

// Handler for force command.
func newForceAction(migrateManager migrate.Migrate) cli.ActionFunc {
	return func(_ context.Context, cmd *cli.Command) error {
		version := cmd.Int("n")

		fmt.Printf("Set %d migration version... \n", version)
		if err := migrateManager.Force(version); err != nil {
			return err
		}
		fmt.Println("Successfully!")
		return nil
	}
}
