package v1

import (
	"skadi/backend/internal/app/entity"
)

// @description solutionOut represents a solution data with students (solving the same task).
type solutionOut struct {
	// solution object
	Solution *entity.Solution `json:"solution" validate:"required"`
	// other students solving the same task
	OtherStudents []entity.Profile `json:"other_students,omitempty" validate:"omitempty"`
}

// @description listSolutionOut represents a solution list data.
type listSolutionOut struct {
	// solutions list
	Data []entity.Solution `json:"data" validate:"required"`
	// pagination params
	Pagination *entity.Pagination `json:"pagination,omitempty" validate:"omitempty"`
}
