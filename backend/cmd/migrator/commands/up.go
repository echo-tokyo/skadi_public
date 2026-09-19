package commands

import (
	"context"
	"fmt"

	cli "github.com/urfave/cli/v3"

	"skadi/backend/internal/pkg/migrate"
)

// Up command instance.
func NewUp(migrateManager migrate.Migrate) *cli.Command {
	return &cli.Command{
		Name:   "up",
		Usage:  "Apply migrations (apply all migrations if the flag -n is not specified)",
		Action: newUpAction(migrateManager),
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:        "number",
				Aliases:     []string{"n"},
				Value:       0,
				Usage:       "Set the number of migrations to apply",
				Validator:   positiveFlagValidator,
				HideDefault: true,
			},
		},
	}
}

// Handler for up command.
func newUpAction(migrateManager migrate.Migrate) cli.ActionFunc {
	return func(_ context.Context, cmd *cli.Command) error {
		var err error
		step := cmd.Int("n")

		if step == 0 {
			fmt.Println("Apply all migrations...")
			err = migrateManager.Up()
		} else {
			fmt.Printf("Apply %d migrations... \n", step)
			err = migrateManager.Step(step)
		}

		if err != nil {
			return err
		}
		fmt.Println("Successfully!")
		return nil
	}
}
