// Package comment contains all repos, usecases and controllers for comment.
// Sub-package repo contains RepoDB implementation.
// Sub-package usecase contains UsecaseClient implementation.
package comment

import "skadi/backend/internal/app/entity"

// UsecaseClient describes all comment usecases for teacher and student.
type UsecaseClient interface {
	// Create creates a new comment and fills the given struct.
	Create(userClaims *entity.UserClaims, commentObj *entity.Comment) error
	// List returns slice of solution comments.
	List(solutionID int, userClaims *entity.UserClaims,
		page *entity.Pagination) ([]entity.Comment, error)
}
