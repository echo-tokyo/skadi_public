// Package file contains all repos, usecases and controllers for file.
// Sub-package repo contains RepoDB implementation.
// Sub-package usecase contains UsecaseClient implementation.
package file

import "skadi/backend/internal/app/entity"

// UsecaseClient describes all file usecases for teacher and student.
type UsecaseClient interface {
	// GetByID returns file metadata by the given ID.
	GetByID(fileID int, userClaims *entity.UserClaims) (*entity.File, error)
}
