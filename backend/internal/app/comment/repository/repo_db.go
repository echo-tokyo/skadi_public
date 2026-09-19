// Package repository contains comment.RepositoryDB implementation.
package repository

import (
	"gorm.io/gorm"

	"skadi/backend/internal/app/comment"
	"skadi/backend/internal/app/entity"
)

// Ensure RepoDB implements interface.
var _ comment.RepositoryDB = (*RepoDB)(nil)

// RepoDB is a comment DB repo.
// It implements the [comment.RepositoryDB] interface.
type RepoDB struct {
	dbStorage *gorm.DB
}

// NewRepoDB returns a new instance of [RepoDB].
func NewRepoDB(dbStorage *gorm.DB) *RepoDB {
	return &RepoDB{
		dbStorage: dbStorage,
	}
}

// Create creates a new comment and fills given struct.
func (r *RepoDB) Create(commentObj *entity.Comment) error {
	return r.dbStorage.Create(commentObj).Error // nil OR error
}

// List returns slice of solution comments.
func (r *RepoDB) List(solutionID int, page *entity.Pagination) ([]entity.Comment, error) {
	comments := []entity.Comment{}
	// create query
	query := r.dbStorage.
		Model(&entity.Comment{}).
		Where("solution_id = ?", solutionID).
		Order("id DESC")
	// apply pagination if it's not nil
	if page != nil {
		page.CountTotal(query)
		query = page.Query(query)
	}
	// exec query
	if err := query.Find(&comments).Error; err != nil {
		return nil, err
	}
	return comments, nil
}
