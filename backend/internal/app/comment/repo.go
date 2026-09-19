package comment

import "skadi/backend/internal/app/entity"

// RepositoryDB describes all DB methods for the comment object.
type RepositoryDB interface {
	// Create creates a new comment and fills given struct.
	Create(commentObj *entity.Comment) error
	// List returns slice of solution comments.
	List(solutionID int, page *entity.Pagination) ([]entity.Comment, error)
}
