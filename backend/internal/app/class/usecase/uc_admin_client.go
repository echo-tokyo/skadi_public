// Package usecase contains class.UsecaseAdmin and class.UsecaseClient implementations.
package usecase

import (
	"errors"
	"fmt"

	"skadi/backend/config"
	"skadi/backend/internal/app/class"
	"skadi/backend/internal/app/entity"
	"skadi/backend/internal/app/user"
)

// Ensure UCAdminClient implements interfaces.
var _ class.UsecaseAdmin = (*UCAdminClient)(nil)
var _ class.UsecaseClient = (*UCAdminClient)(nil)

// UCAdminClient represents a class usecase for admin and client.
// It implements the [class.UsecaseAdmin] and the [class.UsecaseClient] interfaces.
type UCAdminClient struct {
	cfg         *config.Config
	classRepoDB class.RepositoryDB
	userRepoDB  user.RepositoryDB
}

// NewUCAdminClient returns a new instance of [UCAdminClient].
func NewUCAdminClient(cfg *config.Config, classRepoDB class.RepositoryDB,
	userRepoDB user.RepositoryDB) *UCAdminClient {

	return &UCAdminClient{
		cfg:         cfg,
		classRepoDB: classRepoDB,
		userRepoDB:  userRepoDB,
	}
}

// CreateClass creates a new class and fills given struct.
func (u *UCAdminClient) Create(classObj *entity.Class, studentIDs []int) (err error) {
	// get teacher object by ID
	if err := u.setTeacherProfile(classObj); err != nil {
		return err
	}
	// get user objects by IDs (and check them)
	students, err := u.getStudentsByIDs(0, studentIDs)
	if err != nil {
		return fmt.Errorf("get students: %w", err)
	}

	// create class with students
	if err = u.classRepoDB.Create(classObj, studentIDs); err != nil {
		return fmt.Errorf("create class: %w", err)
	}
	// collect student profiles
	classObj.Students = make([]entity.Profile, len(students))
	for idx := range students {
		classObj.Students[idx] = *students[idx].Profile
	}
	return nil
}

// GetByID returns a class object by given ID.
func (u *UCAdminClient) GetByID(id int) (*entity.Class, error) {
	// get class
	classObj, err := u.classRepoDB.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("get class: %w", err)
	}
	// get teacher object by ID
	if err := u.setTeacherProfile(classObj); err != nil {
		return nil, err
	}
	// get students
	classObj.Students, err = u.userRepoDB.GetProfilesShortByClass(classObj.ID)
	return classObj, err // err OR nil
}

// Update updates class by ID (in new class object) with the new data.
// It returns the updated class object.
func (u *UCAdminClient) Update(id int, newData *entity.ClassUpdate) (*entity.Class, error) {
	// get class (without teacher and students)
	classObj, err := u.classRepoDB.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("get class: %w", err)
	}
	// set new name
	if newData.Name != nil {
		classObj.Name = *newData.Name
	}
	// set new schedule
	if newData.Schedule != nil {
		classObj.Schedule = newData.Schedule
	}
	// set new teacher ID
	if newData.TeacherID != nil {
		classObj.TeacherID = newData.TeacherID
	}
	if err := u.setTeacherProfile(classObj); err != nil {
		return nil, err
	}

	if newData.NewFullStudents != nil {
		// set add/del student lists to newData object and
		// get slice of new student profiles
		newData.AddStudents, newData.DelStudents,
			classObj.Students, err = u.sepNewStudents(id, newData.NewFullStudents)
	} else {
		// get actual students list
		classObj.Students, err = u.userRepoDB.GetProfilesShortByClass(id)
	}
	if err != nil {
		return nil, fmt.Errorf("get students: %w", err)
	}

	if err := u.classRepoDB.Update(id, newData); err != nil {
		return nil, fmt.Errorf("update: %w", err)
	}
	return classObj, nil
}

// DeleteByID deletes class object by given id.
func (u *UCAdminClient) DeleteByID(id int) error {
	return u.classRepoDB.DeleteByID(id)
}

// ListShort returns slice of class objects (IDs and names only).
func (u *UCAdminClient) ListShort(search string, page *entity.Pagination) ([]entity.Class, error) {
	return u.classRepoDB.ListShort(search, page)
}

// ListFull returns slice of class objects with full data.
// Search param used to filter classes by name (substring).
func (u *UCAdminClient) ListFull(search string, page *entity.Pagination) ([]entity.Class, error) {
	return u.classRepoDB.ListFull(search, page)
}

// setTeacherProfile sets teacher profile for given class object if teacher ID is presented.
// It returns an error if teacher with given ID not found or teacher role is not "teacher".
func (u *UCAdminClient) setTeacherProfile(classObj *entity.Class) error {
	if classObj.TeacherID == nil || *classObj.TeacherID == 0 {
		return nil
	}
	// get teacher object by ID
	teacherUser, err := u.userRepoDB.GetByIDWithProfileShort(*classObj.TeacherID)
	if errors.Is(err, user.ErrNotFound) {
		// teacher object with such id not found
		return fmt.Errorf("teacher: %w: %s", class.ErrNotFoundUser, err.Error())
	}
	if err != nil {
		return fmt.Errorf("get teacher: %w", err)
	}
	if !teacherUser.IsTeacher() {
		return fmt.Errorf("%w: user %d: not teacher", class.ErrInvalidTeacher, teacherUser.ID)
	}
	classObj.Teacher = teacherUser.Profile
	return nil
}

// sepNewStudents separates students from new students list into add/delete (from class) lists.
// It returns both result lists and new student profiles list.
func (u *UCAdminClient) sepNewStudents(classID int, newStudIDs []int) (add []int, del []int,
	newStuds []entity.Profile, err error) {

	// get actual students list before updating class data
	oldStuds, err := u.userRepoDB.GetProfilesShortByClass(classID)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("old list: %w", err)
	}
	// get user objects of new students (and check them)
	newStudUsers, err := u.getStudentsByIDs(classID, newStudIDs)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("new list: %w", err)
	}
	// collect new student profiles
	newStuds = make([]entity.Profile, len(newStudUsers))
	for idx := range newStudUsers {
		newStuds[idx] = *newStudUsers[idx].Profile
	}
	add, del = u.userRepoDB.AddDelChanges(oldStuds, newStuds)
	return add, del, newStuds, nil
}

// getStudentsByIDs gets student (user) objects by given IDs and validates gotten student list.
func (u *UCAdminClient) getStudentsByIDs(classID int, studentIDs []int) ([]entity.User, error) {
	// get user objects by IDs
	students, err := u.userRepoDB.GetManyWithProfilesShort(studentIDs)
	if err != nil {
		return nil, err
	}
	// check students
	for idx := range students {
		if !students[idx].IsStudent() {
			return nil, fmt.Errorf("%w: user %d: not student",
				class.ErrInvalidStud, students[idx].ID)
		}
		if students[idx].ClassID != nil && *students[idx].ClassID != classID {
			return nil, fmt.Errorf("%w: student %d: already linked to class",
				class.ErrInvalidStud, students[idx].ID)
		}
	}
	return students, nil
}
