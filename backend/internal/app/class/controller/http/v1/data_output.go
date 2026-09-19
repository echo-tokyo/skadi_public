package v1

import "skadi/backend/internal/app/entity"

// @description listClassOut represents a classes list and pagination params.
type listClassOut struct {
	// classes list
	Data []entity.Class `json:"data" validate:"required"`
	// pagination params
	Pagination *entity.Pagination `json:"pagination,omitempty" validate:"omitempty"`
}
