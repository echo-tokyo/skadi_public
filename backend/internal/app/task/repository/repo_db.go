// Package repository contains task.RepositoryDB implementation.
package repository

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"skadi/backend/internal/app/entity"
	"skadi/backend/internal/app/task"
)

const (
	_preloadTask           = "Task"                     // object field name
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
	_fieldUpdatedAt = "updated_at"  // table field name

	_orderByIDDESC = "id DESC" // condition to order data by id DESC

	_defaultStatusID  = 1 // ID of default solution status "backlog"
	_archivedStatusID = 4 // ID of archived solution status "checked"
)

// Ensure RepoDB implements interface.
var _ task.RepositoryDB = (*RepoDB)(nil)

// RepoDB is a task DB repo.
// It implements the [task.RepositoryDB] interface.
type RepoDB struct {
	dbStorage *gorm.DB
}

// NewRepoDB returns a new instance of [RepoDB].
func NewRepoDB(dbStorage *gorm.DB) *RepoDB {
	return &RepoDB{
		dbStorage: dbStorage,
	}
}

// CreateForStudents create a new task and empty solutions for every student.
func (r *RepoDB) CreateForStudents(taskObj *entity.Task,
	students []entity.Profile) ([]entity.Solution, error) {

	var (
		solutions = make([]entity.Solution, len(students))
	)

	// transaction to create the task and solutions
	err := r.dbStorage.Transaction(func(tx *gorm.DB) error {
		// create task and files
		err := tx.Create(taskObj).Error
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			// teacher with given ID is not found
			return fmt.Errorf("task: teacher: %w", task.ErrNotFoundUser)
		}
		if err != nil {
			return err
		}

		// get status object
		var statusObj *entity.Status
		err = tx.Where(_defaultStatusID).First(&statusObj).Error
		if err != nil {
			return fmt.Errorf("get default solution status: %w", err)
		}

		// finish transaction if no one student was given
		if len(students) == 0 {
			return nil
		}

		for idx := range students {
			solutions[idx] = entity.Solution{
				TaskID:    taskObj.ID,
				StudentID: *students[idx].ID,
				StatusID:  _defaultStatusID,
				Student:   &students[idx],
				Status:    statusObj,
			}
		}
		if err = tx.Omit(_preloadStudent, _preloadStatus).Create(solutions).Error; err != nil {
			return fmt.Errorf("solutions for students: %w", err)
		}
		return nil
	})
	return solutions, err // err OR nil
}

// GetByID returns task info by the given ID.
func (r *RepoDB) GetByID(id int) (*entity.Task, error) {
	var taskObj entity.Task
	err := r.dbStorage.
		Preload(_preloadFiles).
		Where(id).First(&taskObj).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// task object with such id not found
		return nil, fmt.Errorf("task with id: %w", task.ErrNotFound)
	}
	return &taskObj, err // err OR nil
}

// Update updates the given task by given ID with the new data.
func (r *RepoDB) Update(taskID int, newData *entity.TaskUpdate) error {
	return r.dbStorage.Transaction(func(tx *gorm.DB) error {
		updates := newData.ToUpdatesMap()

		// update task
		err := tx.Model(&entity.Task{}).
			Where(_fieldID+" = ?", taskID).
			Updates(updates).Error
		if err != nil {
			return fmt.Errorf("task: %w", err)
		}
		// update files
		if err := r.updateTaskFiles(tx, taskID, newData); err != nil {
			return err
		}
		// update solutions
		if err := r.updateTaskSolutions(tx, taskID, newData); err != nil {
			return err
		}
		return err
	})
}

// Delete deletes task and task files.
func (r *RepoDB) Delete(taskObj *entity.Task) error {
	return r.dbStorage.Transaction(func(tx *gorm.DB) error {
		// delete task and task_file records
		if err := tx.Delete(&entity.Task{}, taskObj.ID).Error; err != nil {
			return err
		}
		// delete files
		for _, file := range taskObj.Files {
			if err := tx.Delete(&entity.File{}, file.ID).Error; err != nil {
				return fmt.Errorf("delete file %d: %w", file.ID, err)
			}
		}
		return nil
	})
}

// GetMany returns all teacher tasks.
// Search param appends condition to filter tasks by title (substring).
func (r *RepoDB) GetMany(teacherID int, search string,
	page *entity.Pagination) ([]entity.Task, error) {

	var taskList []entity.Task
	query := r.dbStorage.Model(&entity.Task{}).
		Preload(_preloadFiles).
		Where("teacher_id = ?", teacherID)
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
	if err := query.Find(&taskList).Error; err != nil {
		return nil, err
	}
	return taskList, nil
}

// GetTaskStudents returns all students linked to the given task.
func (r *RepoDB) GetTaskStudents(taskID int) ([]entity.Profile, error) {
	profiles := make([]entity.Profile, 0)
	err := r.dbStorage.Model(entity.Profile{}).
		Select("profile.id", "profile.fullname").
		Joins("INNER JOIN user ON user.id = profile.id").
		Joins("INNER JOIN solution ON solution.student_id = profile.id").
		Joins("INNER JOIN task ON task.id = solution.task_id").
		Where("task.id = ?", taskID).Find(&profiles).Error
	return profiles, err // err OR nil
}

// updateTaskFiles deletes old task files and creates new ones.
func (r *RepoDB) updateTaskFiles(tx *gorm.DB, taskID int, newData *entity.TaskUpdate) error {
	// delete old files
	if len(newData.DelFilesIDs) > 0 {
		err := tx.Where("id IN ?", newData.DelFilesIDs).
			Delete(&entity.File{}).Error
		if err != nil {
			return fmt.Errorf("delete files: %w", err)
		}
	}

	// create new files and link them to the task object
	if len(newData.AddFiles) == 0 {
		return nil
	}
	for _, newFile := range newData.AddFiles {
		if err := tx.Create(newFile).Error; err != nil {
			return fmt.Errorf("create file: %w", err)
		}
		// create many2many relationship
		err := tx.Model(&entity.Task{ID: taskID}).
			Association("Files").Append(newFile)
		if err != nil {
			return fmt.Errorf("associate the file: %w", err)
		}
	}
	return nil
}

// updateTaskSolutions deletes old task solutions and creates new ones.
func (r *RepoDB) updateTaskSolutions(tx *gorm.DB, taskID int, newData *entity.TaskUpdate) error {
	// delete old solutions
	if len(newData.DelStudents) > 0 {
		err := tx.Where("student_id IN ? AND task_id = ?", newData.DelStudents, taskID).
			Delete(&entity.Solution{}).Error
		if err != nil {
			return fmt.Errorf("delete solutions for students: %w", err)
		}
	}

	// skip creating new solutions if add list is empty
	if len(newData.AddStudents) == 0 {
		return nil
	}
	// create new solutions
	solutions := make([]entity.Solution, len(newData.AddStudents))
	for idx, studID := range newData.AddStudents {
		solutions[idx] = entity.Solution{
			TaskID:    taskID,
			StudentID: studID,
			StatusID:  _defaultStatusID,
		}
	}
	if err := tx.Omit(_preloadStudent, _preloadStatus).Create(solutions).Error; err != nil {
		return fmt.Errorf("create solutions for students: %w", err)
	}
	return nil
}
