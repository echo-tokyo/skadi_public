// Package repository contains class.RepositoryDB implementation.
package repository

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"skadi/backend/internal/app/class"
	"skadi/backend/internal/app/entity"
)

const (
	_tableUser             = "User"                // table name
	_preloadTeacher        = "TeacherUser"         // object field name
	_preloadTeacherProfile = "TeacherUser.Profile" // object field name

	_fieldID       = "id"       // table field name
	_fieldName     = "name"     // table field name
	_fieldFullname = "fullname" // table field name
	_fieldClassID  = "class_id" // table field name
)

// Ensure RepoDB implements interface.
var _ class.RepositoryDB = (*RepoDB)(nil)

// RepoDB is a class DB repo.
// It implements the [class.RepositoryDB] interface.
type RepoDB struct {
	dbStorage *gorm.DB
}

// NewRepoDB returns a new instance of [RepoDB].
func NewRepoDB(dbStorage *gorm.DB) *RepoDB {
	return &RepoDB{
		dbStorage: dbStorage,
	}
}

// Create creates a new class and fills given struct.
func (r *RepoDB) Create(classObj *entity.Class, studentIDs []int) error {
	return r.dbStorage.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(classObj).Error
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			// class with such name already exists
			return fmt.Errorf("class with name: %w: %s", class.ErrAlreadyExists, err.Error())
		}
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			// teacher with given ID is not found
			return fmt.Errorf("class: %w: not found", class.ErrInvalidTeacher)
		}
		if err != nil {
			return err
		}

		// update class ID for given students
		err = tx.Model(&entity.User{}).
			Where(_fieldID+" IN ?", studentIDs).
			UpdateColumn(_fieldClassID, classObj.ID).Error
		if err != nil {
			return fmt.Errorf("update students class: %w", err)
		}
		return nil
	})
}

// GetByIDShort returns class (ID and name only) by given id.
func (r *RepoDB) GetByIDShort(id int) (*entity.Class, error) {
	var classObj entity.Class
	err := r.dbStorage.Select(_fieldID, _fieldName).
		Where(id).First(&classObj).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// class object with such id not found
		return nil, fmt.Errorf("class with id: %w: %s", class.ErrNotFound, err.Error())
	}
	return &classObj, err // err OR nil
}

// GetByID returns class by given id.
func (r *RepoDB) GetByID(id int) (*entity.Class, error) {
	var classObj entity.Class
	err := r.dbStorage.Where(id).First(&classObj).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// class object with such id not found
		return nil, fmt.Errorf("class with id: %w: %s", class.ErrNotFound, err.Error())
	}
	return &classObj, err // err OR nil
}

// Update updates class by given ID with the new data.
// It returns the updated class object.
func (r *RepoDB) Update(classID int, newData *entity.ClassUpdate) error {
	return r.dbStorage.Transaction(func(tx *gorm.DB) error {
		updates := newData.ToUpdatesMap()

		// update class
		err := tx.Model(&entity.Class{}).
			Where(_fieldID+" = ?", classID).
			Updates(updates).Error
		if err != nil {
			return fmt.Errorf("class: %w", err)
		}

		// add students to the class
		if len(newData.AddStudents) > 0 {
			err := tx.Model(&entity.User{}).
				Where(_fieldID+" IN ?", newData.AddStudents).
				Update(_fieldClassID, classID).Error
			if err != nil {
				return fmt.Errorf("add students: %w", err)
			}
		}
		// skip deleting students if del list is empty
		if len(newData.DelStudents) == 0 {
			return nil
		}
		// delete students from the class
		err = tx.Model(&entity.User{}).
			Where(_fieldID+" IN ?", newData.DelStudents).
			Update(_fieldClassID, nil).Error
		if err != nil {
			return fmt.Errorf("delete students: %w", err)
		}
		return nil
	})
}

// DeleteByID deletes class object by given id.
func (r *RepoDB) DeleteByID(id int) error {
	return r.dbStorage.Delete(&entity.Class{}, id).Error
}

// ListShort returns slice of class objects (IDs and names only).
func (r *RepoDB) ListShort(search string, page *entity.Pagination) ([]entity.Class, error) {
	classes := []entity.Class{}
	query := r.dbStorage.Model(&entity.Class{}).Select(_fieldID, _fieldName)
	if search != "" {
		query = query.Where("name REGEXP ?", search)
	}
	query = query.Order("id DESC")
	// apply pagination if it's not nil
	if page != nil {
		page.CountTotal(query)
		query = page.Query(query)
	}
	// exec query
	if err := query.Find(&classes).Error; err != nil {
		return nil, err
	}
	return classes, nil
}

// ListFull returns slice of class objects with full data.
// Search param appends condition to filter classes by name (substring).
func (r *RepoDB) ListFull(search string, page *entity.Pagination) ([]entity.Class, error) {
	classes := []entity.Class{}
	// create query
	query := r.dbStorage.
		Model(&entity.Class{}).
		Preload(_preloadTeacher).
		Preload(_preloadTeacherProfile, func(db *gorm.DB) *gorm.DB {
			return db.Select(_fieldID, _fieldFullname) // preload only ID and fullname
		})
	if search != "" {
		query = query.Where("name REGEXP ?", search)
	}
	query = query.Order("id DESC")
	// apply pagination if it's not nil
	if page != nil {
		page.CountTotal(query)
		query = page.Query(query)
	}
	// exec query
	if err := query.Find(&classes).Error; err != nil {
		return nil, err
	}

	for idx := range classes {
		if classes[idx].TeacherUser != nil {
			classes[idx].Teacher = classes[idx].TeacherUser.Profile
		}
	}

	// collect class IDs to the slice
	classIDs := make([]int, len(classes))
	for idx := range classes {
		classIDs[idx] = classes[idx].ID
	}
	// get students by classes
	studByClasses, err := r.getStudentsByClasses(classIDs)
	if err != nil {
		return nil, fmt.Errorf("students by classes: %w", err)
	}
	for idx := range classes {
		classes[idx].Students = studByClasses[classes[idx].ID]
	}
	return classes, nil
}

// student represents a short student info (ID and fullname) with his class ID.
type student struct {
	*entity.Profile
	ClassID int `gorm:"column:class_id"`
}

// getStudentsByClasses returns map with students grouped by his classes.
func (r *RepoDB) getStudentsByClasses(classIDs []int) (map[int][]entity.Profile, error) {
	// select student in all given classes
	var students []*student
	err := r.dbStorage.
		Model(&entity.Profile{}).
		Select("profile.id", _fieldFullname, "user.class_id").
		Joins("INNER JOIN user ON user.id=profile.id").
		Where("user.class_id IN ?", classIDs).
		Find(&students).Error
	if err != nil {
		return nil, err
	}

	// group students by classes
	byClass := make(map[int][]entity.Profile)
	for _, stud := range students {
		byClass[stud.ClassID] = append(byClass[stud.ClassID], *stud.Profile)
	}
	return byClass, nil
}
