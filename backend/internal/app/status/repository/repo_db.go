// Package repository contains status.RepositoryDB implementation.
package repository

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"skadi/backend/internal/app/entity"
	"skadi/backend/internal/app/status"
)

// Ensure RepoDB implements interface.
var _ status.RepositoryDB = (*RepoDB)(nil)

// RepoDB is a status DB repo.
// It implements the [status.RepositoryDB] interface.
type RepoDB struct {
	dbStorage *gorm.DB
}

// NewRepoDB returns a new instance of [RepoDB].
func NewRepoDB(dbStorage *gorm.DB) *RepoDB {
	return &RepoDB{
		dbStorage: dbStorage,
	}
}

// GetByID returns solution status by the given ID.
func (r *RepoDB) GetByID(id int) (*entity.Status, error) {
	var statusObj entity.Status
	err := r.dbStorage.
		Where(id).First(&statusObj).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// status object with such id not found
		return nil, fmt.Errorf("status with id: %w: %s", status.ErrNotFound, err.Error())
	}
	return &statusObj, err // err OR nil
}
