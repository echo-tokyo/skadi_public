// Package solution contains all repos, usecases and controllers for solution.
// Sub-package repo contains RepoDB implementation.
// Sub-package usecase contains UsecaseTeacher, UsecaseStudent and UsecaseClient implementations.
package solution

import "skadi/backend/internal/app/entity"

// UsecaseTeacher describes all class usecases for teacher.
type UsecaseTeacher interface {
	// Update updates the given solution by given ID with the new data.
	// It returns the updated solution object.
	// Allows to update the grade and status (only ready and archived).
	Update(teacherID, solutionID int,
		newData *entity.SolutionUpdate) (*entity.Solution, error)
	// DeleteByID deletes solution object by given ID.
	DeleteByID(userID, solutionID int) error
	// GetManyForTeacher returns all solutions for the teacher tasks.
	// Search param appends condition to filter solutions
	// by task title or student fullname (substring).
	// StatusIDs param appends condition to filter solutions by statuses.
	GetManyForTeacher(teacherID int, search string, statusIDs []int,
		page *entity.Pagination) ([]entity.Solution, error)
}

// UsecaseStudent describes all class usecases for student.
type UsecaseStudent interface {
	// Update updates the given solution by given ID with the new data.
	// It returns the updated solution object.
	// Allows to update the status (apart of archived), answer and solution files.
	Update(studID, solutionID int, newData *entity.SolutionUpdate) (*entity.Solution, error)
	// GetManyForStudent returns all student solutions.
	// Search param appends condition to filter solutions by task title (substring).
	// StatusIDs param appends condition to filter solutions by statuses.
	GetManyForStudent(studID int, search string, statusIDs []int,
		page *entity.Pagination) ([]entity.Solution, error)
}

// UsecaseClient describes all class usecases for teacher and student.
type UsecaseClient interface {
	// GetByIDFull returns a full solution info and all students linked to the solution task.
	GetByIDFull(solutionID int,
		userClaims *entity.UserClaims) (*entity.Solution, []entity.Profile, error)
}
