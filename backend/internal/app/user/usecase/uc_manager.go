package usecase

import (
	"errors"
	"fmt"

	"skadi/backend/config"
	"skadi/backend/internal/app/entity"
	"skadi/backend/internal/app/user"
	"skadi/backend/internal/pkg/password"
)

// Ensure UCManager implements interface.
var _ user.UsecaseManager = (*UCManager)(nil)

// UCManager represents a user usecase for CLI-manager.
// It implements the [user.UsecaseManager] interface.
type UCManager struct {
	cfg        *config.Config
	userRepoDB user.RepositoryDB
}

// NewUCManager returns a new instance of [UCManager].
func NewUCManager(cfg *config.Config, userRepoDB user.RepositoryDB) *UCManager {
	return &UCManager{
		cfg:        cfg,
		userRepoDB: userRepoDB,
	}
}

// CreateAdmin creates a new admin user in the DB and returns them.
// Password is a raw (not hashed) password.
func (u *UCManager) CreateAdmin(username string, passwd []byte) (*entity.User, error) {
	// hash password
	hashPasswd, err := password.Encode(passwd)
	if err != nil {
		return nil, fmt.Errorf("encode password: %w", err)
	}

	userObj := &entity.User{
		Username: username,
		Password: hashPasswd,
		Role:     entity.Admin,
	}
	// create user with default profile
	err = u.userRepoDB.CreateUserWithDefaultProfile(userObj)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}
	return userObj, nil
}

// DeleteAdminByID deletes admin user with given id.
func (u *UCManager) DeleteAdminByID(id int) error {
	userObj, err := u.userRepoDB.GetByID(id)
	if err != nil {
		return fmt.Errorf("get by id: %w", err)
	}
	// deny deletion of non-admins
	if !userObj.IsAdmin() {
		return errors.New("cannot delete non-admin")
	}
	if err := u.userRepoDB.Delete(userObj); err != nil {
		return fmt.Errorf("delete by id: %w", err)
	}
	return nil
}

// GetAdmins returns all admins.
func (u *UCManager) GetAdmins() ([]entity.User, error) {
	userList, err := u.userRepoDB.GetByRoles([]entity.Role{entity.Admin}, false, "", nil)
	if err != nil {
		return nil, fmt.Errorf("get many: %w", err)
	}
	return userList, nil
}
