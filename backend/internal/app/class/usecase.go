// Package class contains all repos, usecases and controllers for class.
// Sub-package repo contains RepoDB and RepoCache implementations.
// Sub-package usecase contains UsecaseClient and UsecaseMiddleware implementations.
package class

import "skadi/backend/internal/app/entity"

// UsecaseAdmin describes all class usecases for admin panel.
type UsecaseAdmin interface {
	// Create creates a new class and fills given struct.
	Create(classObj *entity.Class, studentIDs []int) error
	// Update updates class by ID (in new class object) with the new data.
	// It returns the updated class object.
	Update(id int, newData *entity.ClassUpdate) (*entity.Class, error)
	// DeleteByID deletes class object by given ID.
	DeleteByID(id int) error
}

// UsecaseClient describes all class usecases for client.
type UsecaseClient interface {
	// GetByID returns a class object by given ID.
	GetByID(id int) (*entity.Class, error)
	// ListShort returns slice of class objects (IDs and names only).
	// Search param used to filter classes by name (substring).
	ListShort(search string, page *entity.Pagination) ([]entity.Class, error)
	// ListFull returns slice of class objects with full data.
	// Search param used to filter classes by name (substring).
	ListFull(search string, page *entity.Pagination) ([]entity.Class, error)
}
