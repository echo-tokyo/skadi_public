// Package usecase contains comment.UsecaseClient implementation.
package usecase

import (
	"fmt"

	"skadi/backend/config"
	"skadi/backend/internal/app/comment"
	"skadi/backend/internal/app/entity"
	"skadi/backend/internal/app/solution"
)

// Ensure UCClient implements interfaces.
var _ comment.UsecaseClient = (*UCClient)(nil)

// UCClient represents a comment usecase for teacher and student.
// It implements the [comment.UsecaseClient] interface.
type UCClient struct {
	cfg           *config.Config
	commentRepoDB comment.RepositoryDB
	solRepoDB     solution.RepositoryDB
}

// NewUCClient returns a new instance of [UCClient].
func NewUCClient(cfg *config.Config, commentRepoDB comment.RepositoryDB,
	solRepoDB solution.RepositoryDB) *UCClient {

	return &UCClient{
		cfg:           cfg,
		commentRepoDB: commentRepoDB,
		solRepoDB:     solRepoDB,
	}
}

// Create creates a new comment and fills the given struct.
func (u *UCClient) Create(userClaims *entity.UserClaims, commentObj *entity.Comment) error {
	// check user rights for this solution
	if err := u.solRepoDB.UserPermit(commentObj.SolutionID, userClaims); err != nil {
		return fmt.Errorf("check solution permissions: %w", err)
	}
	// create comment
	return u.commentRepoDB.Create(commentObj)
}

// List returns slice of solution comments.
func (u *UCClient) List(solutionID int, userClaims *entity.UserClaims,
	page *entity.Pagination) ([]entity.Comment, error) {

	// check user rights for this solution
	if err := u.solRepoDB.UserPermit(solutionID, userClaims); err != nil {
		return nil, fmt.Errorf("check solution permissions: %w", err)
	}
	// get list of comments
	return u.commentRepoDB.List(solutionID, page)
}
