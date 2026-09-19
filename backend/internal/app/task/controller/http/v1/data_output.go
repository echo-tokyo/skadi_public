package v1

import (
	"skadi/backend/internal/app/entity"
)

// @description createTaskOut represents a task data with solutions.
type createTaskOut struct {
	// task object
	Task *entity.Task `json:"task" validate:"required"`
	// task solutions
	Solutions []entity.Solution `json:"solutions,omitempty" validate:"omitempty"`
}

// @description listTaskOut represents a task list data.
type listTaskOut struct {
	// tasks list
	Data []entity.TaskWithStudents `json:"data" validate:"required"`
	// pagination params
	Pagination *entity.Pagination `json:"pagination,omitempty" validate:"omitempty"`
}
