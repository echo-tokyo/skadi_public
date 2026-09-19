// Package status contains DB repo for status.
// Sub-package repo contains RepoDB implementation.
package status

import "skadi/backend/internal/app/entity"

// RepositoryDB describes all DB methods for solution status object.
type RepositoryDB interface {
	// GetByID returns solution status by the given ID.
	GetByID(id int) (*entity.Status, error)
}
