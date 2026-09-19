package solution

import "skadi/backend/internal/app/entity"

// RepositoryDB describes all DB methods for task and solution objects.
type RepositoryDB interface {
	// GetByID returns solution info (with task only) by the given ID.
	GetByID(id int) (*entity.Solution, error)
	// GetByIDFull returns a full solution info by the given ID.
	GetByIDFull(id int) (*entity.Solution, error)
	// Update updates the given solution by given ID with the new data.
	Update(solutionID int, newData *entity.SolutionUpdate) error
	// Delete deletes solution and solution files.
	Delete(solObj *entity.Solution) error

	// GetManyForTeacher returns all solutions for the teacher tasks.
	// Search param appends condition to filter solutions
	// by task title or student fullname (substring).
	// StatusIDs param appends condition to filter solutions by statuses.
	GetManyForTeacher(teacherID int, search string, statusIDs []int,
		page *entity.Pagination) ([]entity.Solution, error)
	// GetManyForStudent returns all student solutions.
	// Search param appends condition to filter solutions by task title (substring).
	// StatusIDs param appends condition to filter solutions by statuses.
	GetManyForStudent(studID int, search string, statusIDs []int,
		page *entity.Pagination) ([]entity.Solution, error)

	// UserPermit returns nil error if user has rights to the given solution.
	UserPermit(solutionID int, userClaims *entity.UserClaims) error
}
