package usecase

import (
	"errors"
	"fmt"
	goslices "slices"

	"skadi/backend/config"
	"skadi/backend/internal/app/entity"
	"skadi/backend/internal/app/solution"
	"skadi/backend/internal/app/status"
)

// Ensure UCStudent implements interfaces.
var _ solution.UsecaseStudent = (*UCStudent)(nil)

// UCStudent represents a solution usecase for student.
// It implements the [solution.UsecaseStudent] interface.
type UCStudent struct {
	cfg          *config.Config
	solRepoDB    solution.RepositoryDB
	statusRepoDB status.RepositoryDB
}

// NewUCStudent returns a new instance of [UCStudent].
func NewUCStudent(cfg *config.Config, solRepoDB solution.RepositoryDB,
	statusRepoDB status.RepositoryDB) *UCStudent {

	return &UCStudent{
		cfg:          cfg,
		solRepoDB:    solRepoDB,
		statusRepoDB: statusRepoDB,
	}
}

// Update updates the given solution by given ID with the new data.
// It returns the updated solution object.
// Allows to update the status (apart of archived), answer and solution files.
func (u *UCStudent) Update(studID, solutionID int,
	newData *entity.SolutionUpdate) (*entity.Solution, error) {

	// get solution with files
	solObj, err := u.solRepoDB.GetByIDFull(solutionID)
	if err != nil {
		return nil, fmt.Errorf("get full solution: %w", err)
	}
	if studID != solObj.StudentID {
		return nil, fmt.Errorf("%w: user (student) is not a solution owner",
			solution.ErrForbidden)
	}

	newData.DelFiles = make(entity.Files, 0, len(newData.DelFilesIDs))
	solFilesRemains := make(entity.Files, 0, len(solObj.Files))
	// collect files' metadata to delete them from solution
	for _, oldFile := range solObj.Files {
		if goslices.Contains(newData.DelFilesIDs, oldFile.ID) {
			newData.DelFiles = append(newData.DelFiles, oldFile)
		} else {
			solFilesRemains = append(solFilesRemains, oldFile)
		}
	}
	solObj.Files = solFilesRemains

	newData.Grade = nil
	if newData.StatusID != nil {
		if err := u.getStatusToUpdate(solObj, *newData.StatusID); err != nil {
			return nil, err
		}
	}
	if newData.Answer != nil {
		solObj.Answer = newData.Answer
	}

	// check that solution has either an answer or at least one file to change status to "ready"
	if err := statusRestrictions(solObj); err != nil {
		return nil, err
	}

	if err := u.solRepoDB.Update(solutionID, newData); err != nil {
		return nil, fmt.Errorf("update: %w", err)
	}
	solObj.UpdatedAt = newData.UpdatedAt

	// append new files to solution files
	solObj.Files = append(solObj.Files, newData.AddFiles...)
	// remove deleted files from file system
	newData.DelFiles.Cleanup()

	return solObj, nil
}

// GetManyForStudent returns all student solutions.
// Search param appends condition to filter solutions by task title (substring).
// StatusID param appends condition to filter solutions by status.
func (u *UCStudent) GetManyForStudent(studID int, search string, statusID []int,
	page *entity.Pagination) ([]entity.Solution, error) {

	return u.solRepoDB.GetManyForStudent(studID, search, statusID, page)
}

// getStatusToUpdate sets the new status object to updated solution.
func (u *UCStudent) getStatusToUpdate(solObj *entity.Solution, newStatusID int) error {
	solObj.StatusID = newStatusID
	// apart of archived statuses
	if newStatusID == _archivedStatusID {
		return fmt.Errorf("%w: student cannot set the given status", solution.ErrInvalidData)
	}
	// get status object
	var err error
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

// statusRestrictions checks data-status compliance when updating solution.
func statusRestrictions(solObj *entity.Solution) error {
	// check that solution has either an answer or at least one file to change status to "ready"
	zeroFiles := len(solObj.Files) == 0
	nilAnswer := solObj.Answer == nil || solObj.Answer != nil && *solObj.Answer == ""

	if solObj.StatusID == _readyStatusID && zeroFiles && nilAnswer {
		return fmt.Errorf("%w: set ready status without any answer", solution.ErrUnsupportedData)
	}
	return nil
}
