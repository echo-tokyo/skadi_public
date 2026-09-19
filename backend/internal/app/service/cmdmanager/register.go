package cmdmanager

import (
	"gorm.io/gorm"

	"skadi/backend/config"
	usercli "skadi/backend/internal/app/user/controller/cli"
	userrepo "skadi/backend/internal/app/user/repository"
	useruc "skadi/backend/internal/app/user/usecase"
)

// registerCommands register all cmd manager commands.
func (c *CmdManager) registerCommands(cfg *config.Config, dbStorage *gorm.DB) {
	// create repos
	userRepoDB := userrepo.NewRepoDB(dbStorage)
	// create usecases
	userUCManager := useruc.NewUCManager(cfg, userRepoDB)
	// create controllers
	userController := usercli.NewUserController(userUCManager)

	c.commands = map[string]Handler{
		"create-admin": userController.CreateAdmin,
		"delete-admin": userController.DeleteAdmin,
		"get-admins":   userController.GetAdmins,
	}
}
