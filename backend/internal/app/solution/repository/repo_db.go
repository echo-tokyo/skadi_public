// Package repository contains task.RepositoryDB implementation.
package repository

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"skadi/backend/internal/app/entity"
	"skadi/backend/internal/app/solution"
)

const (
	_preloadTask           = "Task"                     // object field name
	_preloadTaskFiles      = "Task.Files"               // object field name
	_preloadTeacher        = "Task.TeacherUser"         // object field name
	_preloadTeacherProfile = "Task.TeacherUser.Profile" // object field name
	_preloadStudent        = "StudentUser"              // object field name
	_preloadStudentProfile = "StudentUser.Profile"      // object field name
	_preloadStatus         = "Status"                   // object field name
	_preloadFiles          = "Files"                    // object field name

	_fieldID        = "id"          // table field name
	_fieldFullname  = "fullname"    // table field name
	_fieldTitle     = "title"       // table field name
	_fieldDesc      = "description" // table field name
	_fieldTeacherID = "teacher_id"  // table field name
	_fieldUpdatedAt = "updated_at"  // table field name

	_orderByIDDESC = "id DESC" // condition to order data by id DESC
)

// Ensure RepoDB implements interface.
var _ solution.RepositoryDB = (*RepoDB)(nil)

// RepoDB is a task DB repo.
// It implements the [solution.RepositoryDB] interface.
type RepoDB struct {
	dbStorage *gorm.DB
}

// NewRepoDB returns a new instance of [RepoDB].
func NewRepoDB(dbStorage *gorm.DB) *RepoDB {
	return &RepoDB{
		dbStorage: dbStorage,
	}
}

// GetByID returns solution info (with task only) by the given ID.
func (r *RepoDB) GetByID(id int) (*entity.Solution, error) {
	var solObj entity.Solution
	err := r.dbStorage.
		Preload(_preloadTask).
		Where(id).First(&solObj).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// solution object with such id not found
		return nil, fmt.Errorf("solution with id: %w", solution.ErrNotFound)
	}
	return &solObj, err // err OR nil
}

// GetByIDFull returns a full solution info by the given ID.
func (r *RepoDB) GetByIDFull(id int) (*entity.Solution, error) {
	var solObj entity.Solution
	err := r.dbStorage.
		Preload(_preloadTask).
		Preload(_preloadTaskFiles).
		Preload(_preloadTeacher).
		Preload(_preloadTeacherProfile, func(db *gorm.DB) *gorm.DB {
			return db.Select(_fieldID, _fieldFullname) // preload only ID and fullname
		}).
		Preload(_preloadStudent).
		Preload(_preloadStudentProfile, func(db *gorm.DB) *gorm.DB {
			return db.Select(_fieldID, _fieldFullname) // preload only ID and fullname
		}).
		Preload(_preloadStatus).
		Preload(_preloadFiles).
		Where(id).First(&solObj).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// solution object with such id not found
		return nil, fmt.Errorf("solution with id: %w", solution.ErrNotFound)
	}
	if err != nil {
		return nil, err
	}

	// set student and teacher profiles
	solObj.Student = solObj.StudentUser.Profile
	solObj.Task.Teacher = solObj.Task.TeacherUser.Profile
	return &solObj, nil
}

// Update updates the given solution by given ID with the new data.
func (r *RepoDB) Update(solutionID int, newData *entity.SolutionUpdate) error {
	return r.dbStorage.Transaction(func(tx *gorm.DB) error {
		updates := newData.ToUpdatesMap()

		var updatedSol *entity.Solution
		// update solution
		err := tx.Model(&entity.Solution{}).
			Where(_fieldID+" = ?", solutionID).
			Updates(updates).
			Scan(&updatedSol).Error
		if err != nil {
			return err
		}
		// update solution files
		if err := r.updateSolutionFiles(tx, solutionID, newData); err != nil {
			return err
		}

		newData.UpdatedAt = updatedSol.UpdatedAt
		return nil
	})
}

// Delete deletes solution and solution files.
func (r *RepoDB) Delete(solObj *entity.Solution) error {
	return r.dbStorage.Transaction(func(tx *gorm.DB) error {
		// delete solution and solution_file records
		if err := tx.Delete(&entity.Solution{}, solObj.ID).Error; err != nil {
			return err
		}
		// delete files
		for _, file := range solObj.Files {
			if err := tx.Delete(&entity.File{}, file.ID).Error; err != nil {
				return fmt.Errorf("delete file %d: %w", file.ID, err)
			}
		}
		return nil
	})
}

// GetManyForTeacher returns all solutions for the teacher tasks.
// Search param appends condition to filter solutions
// by task title or student fullname (substring).
// StatusIDs param appends condition to filter solutions by statuses.
func (r *RepoDB) GetManyForTeacher(teacherID int, search string, statusIDs []int,
	page *entity.Pagination) ([]entity.Solution, error) {

	solList := make([]entity.Solution, 0)
	query := r.dbStorage.Model(entity.Solution{}).
		Omit("answer", "updated_at").
		Preload(_preloadTask, func(db *gorm.DB) *gorm.DB {
			return db.Select(_fieldID, _fieldTitle) // preload only ID and title
		}).
		Preload(_preloadStudent).
		Preload(_preloadStudentProfile, func(db *gorm.DB) *gorm.DB {
			return db.Select(_fieldID, _fieldFullname) // preload only ID and fullname
		}).
		Preload(_preloadStatus).
		Joins("INNER JOIN task ON task.id = solution.task_id").
		Joins("INNER JOIN user ON solution.student_id = user.id").
		Joins("INNER JOIN profile ON user.id=profile.id").
		Where("task.teacher_id = ?", teacherID)
	// add filters
	if len(statusIDs) != 0 {
		query = query.Where("status_id IN ?", statusIDs)
	}
	if search != "" {
		query = query.Where("title REGEXP ? OR profile.fullname REGEXP ?", search, search)
	}
	query = query.Order(_orderByIDDESC)
	// apply pagination if it's not nil
	if page != nil {
		page.CountTotal(query)
		query = page.Query(query)
	}
	// exec query
	if err := query.Find(&solList).Error; err != nil {
		return nil, err
	}
	return solList, nil
}

// GetManyForStudent returns all student solutions.
// StatusIDs param appends condition to filter solutions by statuses.
func (r *RepoDB) GetManyForStudent(studID int, search string, statusIDs []int,
	page *entity.Pagination) ([]entity.Solution, error) {

	solList := make([]entity.Solution, 0)
	query := r.dbStorage.Model(entity.Solution{}).
		Omit("answer", "student_id").
		Preload(_preloadTask, func(db *gorm.DB) *gorm.DB {
			return db.Select(_fieldID, _fieldTitle, _fieldDesc) // preload only ID, title and desc
		}).
		Preload(_preloadStatus).
		Joins("INNER JOIN task ON task.id = solution.task_id").
		Where("student_id = ?", studID)
	// add filters
	if len(statusIDs) != 0 {
		query = query.Where("status_id IN ?", statusIDs)
	}
	if search != "" {
		query = query.Where("title REGEXP ?", search)
	}
	query = query.Order(_orderByIDDESC)
	// apply pagination if it's not nil
	if page != nil {
		page.CountTotal(query)
		query = page.Query(query)
	}
	// exec query
	if err := query.Find(&solList).Error; err != nil {
		return nil, err
	}
	return solList, nil
}

// UserPermit returns nil error if user has rights to the given solution.
func (r *RepoDB) UserPermit(solutionID int, userClaims *entity.UserClaims) error {
	// get student_id from solution and teacher_id from task
	var solObj entity.Solution
	err := r.dbStorage.
		Model(&entity.Solution{}).
		Select("task_id, student_id").
		Preload(_preloadTask, func(db *gorm.DB) *gorm.DB {
			return db.Select(_fieldID, _fieldTeacherID) // preload only teacher ID
		}).
		Where(solutionID).First(&solObj).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// solution object with such id not found
		return fmt.Errorf("solution with id: %w", solution.ErrNotFound)
	}
	if err != nil {
		return err
	}

	// role checks
	if userClaims.IsTeacher() && solObj.Task.TeacherID != userClaims.ID {
		return fmt.Errorf("%w: user (teacher) is not a task owner", solution.ErrForbidden)
	}
	if userClaims.IsStudent() && solObj.StudentID != userClaims.ID {
		return fmt.Errorf("%w: user (stud) is not a solution owner", solution.ErrForbidden)
	}
	return nil
}

// updateSolutionFiles deletes old solution files and creates new ones.
func (r *RepoDB) updateSolutionFiles(tx *gorm.DB,
	solID int, newData *entity.SolutionUpdate) error {

	// delete old files
	if len(newData.DelFilesIDs) > 0 {
		err := tx.Where("id IN ?", newData.DelFilesIDs).
			Delete(&entity.File{}).Error
		if err != nil {
			return fmt.Errorf("delete files: %w", err)
		}
	}

	// create new files and link them to the solution object
	if len(newData.AddFiles) == 0 {
		return nil
	}
	for _, newFile := range newData.AddFiles {
		if err := tx.Create(newFile).Error; err != nil {
			return fmt.Errorf("create file: %w", err)
		}
		// create many2many relationship
		err := tx.Model(&entity.Solution{ID: solID}).
			Association("Files").Append(newFile)
		if err != nil {
			return fmt.Errorf("associate the file: %w", err)
		}
	}
	return nil
}
