package task

import "skadi/backend/internal/app/entity"

// RepositoryDB describes all DB methods for task and solution objects.
type RepositoryDB interface {
	// CreateForStudents create a new task and empty solutions for every student.
	CreateForStudents(taskObj *entity.Task, students []entity.Profile) ([]entity.Solution, error)
	// GetByID returns task info by the given ID.
	GetByID(id int) (*entity.Task, error)
	// Update updates the given task by given ID with the new data.
	Update(taskID int, newData *entity.TaskUpdate) error
	// Delete deletes task and task files.
	Delete(taskObj *entity.Task) error

	// GetMany returns all teacher tasks.
	// Search param appends condition to filter tasks by title (substring).
	GetMany(teacherID int, search string, page *entity.Pagination) ([]entity.Task, error)
	// GetTaskStudents returns all students linked to the given task.
	GetTaskStudents(taskID int) ([]entity.Profile, error)
}
