// Package task contains all repos, usecases and controllers for task.
// Sub-package repo contains RepoDB implementation.
// Sub-package usecase contains UsecaseTeacher implementation.
package task

import "skadi/backend/internal/app/entity"

// UsecaseAdmin describes all class usecases for teacher.
type UsecaseTeacher interface {
	// CreateWithSolutions creates a new task and solutions
	// for all given students and for all students linked to the given classes.
	CreateWithSolutions(taskObj *entity.Task, studentIDs []int,
		classIDs []int) ([]entity.Solution, error)
	// GetByID returns a task object by the given id and
	// a list of students linked to the task solutions.
	GetByID(teacherID, taskID int) (*entity.Task, []entity.Profile, error)
	// Update updates the given task by given ID with the new data.
	// It returns the updated task object and updated students linked to the task.
	// Allows to update the title, desc, linked students and task files.
	Update(teacherID, taskID int,
		newData *entity.TaskUpdate) (*entity.Task, []entity.Profile, error)
	// DeleteByID deletes task object by given ID.
	DeleteByID(userID, taskID int) error

	// GetMany returns all teacher tasks.
	// Search param appends condition to filter tasks by title (substring).
	// For each elem returns task object and list of students linked to it.
	GetMany(teacherID int, search string,
		page *entity.Pagination) ([]entity.TaskWithStudents, error)
}
