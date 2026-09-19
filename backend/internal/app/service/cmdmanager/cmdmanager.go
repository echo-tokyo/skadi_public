// Package cmdmanager provides command line manager service.
package cmdmanager

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"gorm.io/gorm"

	"skadi/backend/config"
)

// Handler represents a command handler.
type Handler func() error

// CmdManager represents a command line manager.
type CmdManager struct {
	// will be closed if the service was completely started and is ready-to-use now
	ready    chan struct{}
	cfg      *config.Config
	commands map[string]Handler
}

// New returns a new instance of [CmdManager].
func New(cfg *config.Config, dbStorage *gorm.DB) (*CmdManager, error) {
	manager := &CmdManager{
		ready: make(chan struct{}),
		cfg:   cfg,
	}
	// register all commands
	manager.registerCommands(cfg, dbStorage)
	return manager, nil
}

// StartWithShutdown starts cmd manager and waits for
// context is done for gracefully shutdown it.
// This method is blocking.
func (c *CmdManager) StartWithShutdown(ctx context.Context) error {
	slog.Info("start cmd manager...")
	defer slog.Info("stop cmd manager: ok")

	successChan := make(chan struct{}, 1)
	defer close(successChan)
	errChan := make(chan error, 1)
	defer close(errChan)
	// start cmd manager
	go func() {
		if err := c.run(); err != nil {
			errChan <- fmt.Errorf("cmd manager: %w", err)
		}
		successChan <- struct{}{}
	}()

	// notify that service is ready-to-use
	close(c.ready)
	// wait for context, cmd manager success or error
	select {
	case <-ctx.Done():
		return nil
	case <-successChan:
		return nil
	case err := <-errChan:
		return err
	}
}

// Ready signals that the service is ready-to-use.
func (c *CmdManager) Ready() <-chan struct{} {
	return c.ready
}

// run runs cmd manager to serve registered commands.
func (c *CmdManager) run() error {
	// collect available commands slice to string
	var builder strings.Builder
	for command := range c.commands {
		builder.WriteString("\t - ")
		builder.WriteString(command)
		builder.WriteString("\n")
	}
	cmdAvailable := builder.String()

	var (
		cmd     string
		handler Handler
		ok      bool
		err     error
	)
	for handler == nil {
		// ask for command
		fmt.Println("Available commands:")
		fmt.Print(cmdAvailable)
		fmt.Print("Enter command (or use Ctrl+C to exit): ")
		fmt.Scan(&cmd)

		// get command handler
		if handler, ok = c.commands[cmd]; !ok {
			fmt.Println("Invalid command!")
			continue
		}
	}
	// exec command
	if err = handler(); err != nil {
		return fmt.Errorf("command %s: %w", cmd, err)
	}
	return nil
}
