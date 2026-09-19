package v1

import "skadi/backend/internal/app/entity"

// @description listCommentOut represents a comment list and pagination params.
type listCommentOut struct {
	// comment list
	Data []entity.Comment `json:"data" validate:"required"`
	// pagination params
	Pagination *entity.Pagination `json:"pagination,omitempty" validate:"omitempty"`
}
