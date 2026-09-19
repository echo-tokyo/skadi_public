package usecase

import (
	"errors"
	"fmt"

	"skadi/backend/config"
	"skadi/backend/internal/app/entity"
	"skadi/backend/internal/app/solution"
	"skadi/backend/internal/app/status"
)

const (
	_defaultStatusID  = 1 // ID of solution status "backlog"
	_inWorkStatusID   = 2 // ID of solution status "in-work"
	_readyStatusID    = 3 // ID of solution status "ready"
	_archivedStatusID = 4 // ID of solution status "checked"
)

// Ensure UCAdmin implements interfaces.
var _ solution.UsecaseTeacher = (*UCTeacher)(nil)

// UCTeacher represents a solution usecase for teacher.
// It implements the [solution.UsecaseTeacher] interface.
type UCTeacher struct {
	cfg          *config.Config
	solRepoDB    solution.RepositoryDB
	statusRepoDB status.RepositoryDB
}

// NewUCTeacher returns a new instance of [UCTeacher].
func NewUCTeacher(cfg *config.Config, solRepoDB solution.RepositoryDB,
	statusRepoDB status.RepositoryDB) *UCTeacher {

	return &UCTeacher{
		cfg:          cfg,
		solRepoDB:    solRepoDB,
		statusRepoDB: statusRepoDB,
	}
}

// Update updates the given solution by given ID with the new data.
// It returns the updated solution object.
// Allows to update the grade and status (only ready and archived).
func (u *UCTeacher) Update(teacherID, solutionID int,
	newData *entity.SolutionUpdate) (*entity.Solution, error) {

	solObj, err := u.solRepoDB.GetByIDFull(solutionID)
	if err != nil {
		return nil, fmt.Errorf("get full solution: %w", err)
	}
	if teacherID != solObj.Task.TeacherID {
		return nil, fmt.Errorf("%w: user is not a solution task owner", solution.ErrForbidden)
	}

	newData.Answer = nil
	if newData.StatusID != nil {
		if err := u.getStatusToUpdate(solObj, *newData.StatusID); err != nil {
			return nil, err
		}
	}
	if newData.Grade != nil {
		solObj.Grade = newData.Grade
	}

	if err := u.solRepoDB.Update(solutionID, newData); err != nil {
		return nil, fmt.Errorf("update: %w", err)
	}
	solObj.UpdatedAt = newData.UpdatedAt
	return solObj, nil
}

// DeleteByID deletes solution object by given ID.
func (u *UCTeacher) DeleteByID(userID, solutionID int) error {
	// get short solution info
	solObj, err := u.solRepoDB.GetByID(solutionID)
	// return nil error if solution was not found
	if errors.Is(err, solution.ErrNotFound) {
		return nil
	}
	if err != nil {
		return err
	}
	// check that given user (teacher) is a solution task owner
	if userID != solObj.Task.TeacherID {
		return fmt.Errorf("%w: user (teacher) is not a solution task owner", solution.ErrForbidden)
	}
	// delete solution
	if err := u.solRepoDB.Delete(solObj); err != nil {
		return err
	}
	// delete files from the file system
	solObj.Files.Cleanup()
	return nil
}

// GetManyForTeacher returns all solutions for the teacher tasks.
// Search param appends condition to filter solutions
// by task title or student fullname (substring).
// // StatusID param appends condition to filter solutions by status.
func (u *UCTeacher) GetManyForTeacher(teacherID int, search string, statusID []int,
	page *entity.Pagination) ([]entity.Solution, error) {

	solList, err := u.solRepoDB.GetManyForTeacher(teacherID, search, statusID, page)
	if err != nil {
		return nil, err
	}
	// set student field for each solution to serialize it
	for idx := range solList {
		solList[idx].Student = solList[idx].StudentUser.Profile
	}
	return solList, nil
}

// getStatusToUpdate sets the new status object to updated solution.
func (u *UCTeacher) getStatusToUpdate(solObj *entity.Solution, newStatusID int) error {
	var err error
	// get status object
	solObj.StatusID = newStatusID
	solObj.Status, err = u.statusRepoDB.GetByID(newStatusID)
	// if status object with such id not found
	if errors.Is(err, status.ErrNotFound) {
		return fmt.Errorf("%w: status: %s", solution.ErrInvalidData, err.Error())
	}
	if err != nil {
		return fmt.Errorf("get status: %w", err)
	}
	return nil
}
