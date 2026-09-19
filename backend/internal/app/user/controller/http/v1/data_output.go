package v1

import "skadi/backend/internal/app/entity"

// @description listUserOut represents a users list and pagination params.
type listUserOut struct {
	// users list
	Data []entity.User `json:"data" validate:"required"`
	// pagination params
	Pagination *entity.Pagination `json:"pagination,omitempty" validate:"omitempty"`
}
