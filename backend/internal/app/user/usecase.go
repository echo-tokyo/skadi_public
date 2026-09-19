// Package user contains all repos, usecases and controllers for user.
// Sub-package repo contains RepoDB implementation.
// Sub-package usecase contains UsecaseManager, UsecaseAdmin and UsecaseClient implementations.
package user

import "skadi/backend/internal/app/entity"

// UsecaseManager describes all user usecases for CLI-manager.
type UsecaseManager interface {
	// CreateAdmin creates a new admin user in the DB and returns them.
	// Password is a raw (not hashed) password.
	CreateAdmin(username string, passwd []byte) (*entity.User, error)
	// DeleteAdminByID deletes admin user with given id.
	DeleteAdminByID(id int) error
	// GetAdmins returns all admins.
	GetAdmins() ([]entity.User, error)
}

// UsecaseAdmin describes all user usecases for admin panel.
type UsecaseAdmin interface {
	// CreateWithProfile creates a new user with profile in the DB and returns them.
	// User.Password is a raw (not hashed) password.
	CreateWithProfile(userObj *entity.User) error
	// GetByID returns user object with profile by given id.
	GetByID(id int) (*entity.User, error)
	// Update updates user (class, password and profile) by given ID with given data.
	// It returns user object with updated user data.
	Update(id int, data *entity.User) (*entity.User, error)
	// DeleteByID deletes user object and user profile with given id.
	DeleteByID(id int) error
	// GetByRoles returns user list with given roles.
	// Free param (if only student role was given) used to get class-free students.
	// Search param used to filter users by username and fullname (substring).
	GetByRoles(roleList []entity.Role, free bool,
		search string, page *entity.Pagination) ([]entity.User, error)
	// ChangePasswordAsAdmin changes password of any client.
	// New password is a raw (not hashed) password.
	ChangePasswordAsAdmin(id int, newPasswd []byte) error
}

// UsecaseClient describes all user usecases for client.
type UsecaseClient interface {
	// GetByID returns user object with profile by given id.
	GetByID(id int) (*entity.User, error)
	// UpdateProfile updates user profile by given ID with given data.
	// It returns user object with updated profile data.
	UpdateProfile(id int, data *entity.Profile) (*entity.User, error)
	// ChangePasswordAsClient changes client password.
	// New password is a raw (not hashed) password.
	// The new password cannot be the same as the old one.
	ChangePasswordAsClient(id int, oldPasswd, newPasswd []byte) error
}
