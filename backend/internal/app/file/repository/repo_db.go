// Package repository contains file.RepositoryDB implementation.
package repository

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"skadi/backend/internal/app/entity"
	"skadi/backend/internal/app/file"
)

// Ensure RepoDB implements interface.
var _ file.RepositoryDB = (*RepoDB)(nil)

// RepoDB is a file DB repo.
// It implements the [file.RepositoryDB] interface.
type RepoDB struct {
	dbStorage *gorm.DB
}

// NewRepoDB returns a new instance of [RepoDB].
func NewRepoDB(dbStorage *gorm.DB) *RepoDB {
	return &RepoDB{
		dbStorage: dbStorage,
	}
}

// GetByID returns file by the given ID.
func (r *RepoDB) GetByID(id int) (*entity.File, error) {
	var fileObj entity.File
	err := r.dbStorage.
		Where(id).First(&fileObj).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// file object with such id not found
		return nil, fmt.Errorf("file with id: %w: %s", file.ErrNotFound, err.Error())
	}
	return &fileObj, err // err OR nil
}

// TeacherPermit returns nil error if teacher has rights to the given file.
func (r *RepoDB) TeacherPermit(teacherID, fileID int) error {
	// check tasks
	/*
		task:
		    teacher_id
		    files
	*/
	var taskIDs []int
	err := r.dbStorage.
		Model(&entity.Task{}).
		Select("task.id").
		Joins("INNER JOIN task_file ON task_file.task_id = task.id").
		Where("teacher_id = ?", teacherID).
		Where("task_file.file_id = ?", fileID).
		Scan(&taskIDs).Error
	if err != nil {
		return fmt.Errorf("get teacher tasks with file: %w", err)
	}
	// if the file was found in one of the teacher tasks
	if len(taskIDs) > 0 {
		return nil
	}

	// check solutions
	/*
		solution:
		    task:
		        teacher_id
		    files
	*/
	var solIDs []int
	err = r.dbStorage.
		Model(&entity.Solution{}).
		Select("solution.id").
		Joins("INNER JOIN task ON task.id = solution.task_id").
		Joins("INNER JOIN solution_file ON solution_file.solution_id = solution.id").
		Where("task.teacher_id = ?", teacherID).
		Where("solution_file.file_id = ?", fileID).
		Scan(&solIDs).Error
	if err != nil {
		return fmt.Errorf("get solutions with file with teacher tasks: %w", err)
	}
	// if the file was found in a solution linked to one of the teacher tasks
	if len(solIDs) > 0 {
		return nil
	}

	// forbidden error
	return fmt.Errorf("%w: user %d (teacher) has no one relationship with file %d",
		file.ErrForbidden, teacherID, fileID)
}

// StudentPermit returns nil error if student has rights to the given file.
func (r *RepoDB) StudentPermit(studentID, fileID int) error {
	// check solutions
	/*
		solution:
		    student_id
		    files
	*/
	var solIDs []int
	err := r.dbStorage.
		Model(&entity.Solution{}).
		Select("solution.id").
		Joins("INNER JOIN solution_file ON solution_file.solution_id = solution.id").
		Where("solution.student_id = ?", studentID).
		Where("solution_file.file_id = ?", fileID).
		Scan(&solIDs).Error
	if err != nil {
		return fmt.Errorf("get student solutions with file: %w", err)
	}
	// if the file was found in one of the student solutions
	if len(solIDs) > 0 {
		return nil
	}

	// check tasks
	/*
		solution:
		    student_id
		    task:
		        files
	*/
	var taskIDs []int
	err = r.dbStorage.
		Model(&entity.Task{}).
		Select("task.id").
		Joins("INNER JOIN solution ON solution.task_id = task.id").
		Joins("INNER JOIN task_file ON task_file.task_id = task.id").
		Where("solution.student_id = ?", studentID).
		Where("task_file.file_id = ?", fileID).
		Scan(&taskIDs).Error
	if err != nil {
		return fmt.Errorf("get tasks with file for student solutions: %w", err)
	}
	// if the file was found in a task for one of the student solutions
	if len(taskIDs) > 0 {
		return nil
	}

	// forbidden error
	return fmt.Errorf("%w: user %d (stud) has no one relationship with file %d",
		file.ErrForbidden, studentID, fileID)
}
