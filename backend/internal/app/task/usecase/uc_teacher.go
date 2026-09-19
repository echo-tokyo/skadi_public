// Package usecase contains task.UsecaseTeacher implementation.
package usecase

import (
	"errors"
	"fmt"
	goslices "slices"
	"sync"

	"skadi/backend/config"
	"skadi/backend/internal/app/entity"
	"skadi/backend/internal/app/task"
	"skadi/backend/internal/app/user"
	"skadi/backend/internal/pkg/utils/slices"
)

// Ensure UCAdmin implements interfaces.
var _ task.UsecaseTeacher = (*UCTeacher)(nil)

// UCTeacher represents a task usecase for teacher.
// It implements the [task.UsecaseTeacher] interface.
type UCTeacher struct {
	cfg        *config.Config
	taskRepoDB task.RepositoryDB
	userRepoDB user.RepositoryDB
}

// NewUCTeacher returns a new instance of [UCTeacher].
func NewUCTeacher(cfg *config.Config, taskRepoDB task.RepositoryDB,
	userRepoDB user.RepositoryDB) *UCTeacher {

	return &UCTeacher{
		cfg:        cfg,
		taskRepoDB: taskRepoDB,
		userRepoDB: userRepoDB,
	}
}

// CreateWithSolutions creates a new task and solutions
// for all given students and for all students linked to the given classes.
func (u *UCTeacher) CreateWithSolutions(taskObj *entity.Task,
	studentIDs []int, classIDs []int) ([]entity.Solution, error) {

	teacher, err := u.userRepoDB.GetByIDWithProfileShort(taskObj.TeacherID)
	if err != nil {
		return nil, fmt.Errorf("get teacher: %w", err)
	}
	if !teacher.IsTeacher() {
		return nil, fmt.Errorf("%w: user %d: not teacher", task.ErrInvalidTeacher, teacher.ID)
	}

	// get all users by student IDs
	studentUsers, err := u.userRepoDB.GetManyWithProfilesShort(studentIDs)
	if err != nil {
		return nil, fmt.Errorf("get students by ids: %w", err)
	}
	studentProfiles := make([]entity.Profile, 0, len(studentUsers))
	// check user roles
	for idx := range studentUsers {
		if !studentUsers[idx].IsStudent() {
			return nil, fmt.Errorf("%w: user %d: not student",
				task.ErrInvalidStudent, studentUsers[idx].ID)
		}
		studentProfiles = append(studentProfiles, *studentUsers[idx].Profile)
	}

	var classStudents []entity.Profile
	// get all students by class IDs
	for _, classID := range classIDs {
		classStudents, err = u.userRepoDB.GetProfilesShortByClass(classID)
		if err != nil {
			return nil, fmt.Errorf("get students: class %d: %w", classID, err)
		}
		studentProfiles = append(studentProfiles, classStudents...)
	}
	// delete duplicated students
	studentProfiles = slices.DelDuplsFunc(studentProfiles, func(p entity.Profile) int {
		return *p.ID
	})

	// set teacher profile
	taskObj.TeacherUser = teacher
	taskObj.Teacher = teacher.Profile
	// create task and solutions for all collected students
	solutions, err := u.taskRepoDB.CreateForStudents(taskObj, studentProfiles)
	if err != nil {
		return nil, fmt.Errorf("create task for students: %w", err)
	}
	return solutions, nil
}

// GetByID returns a task object by the given id and
// a list of students linked to the task solutions.
func (u *UCTeacher) GetByID(teacherID, taskID int) (*entity.Task, []entity.Profile, error) {
	// get task
	taskObj, err := u.taskRepoDB.GetByID(taskID)
	if err != nil {
		return nil, nil, fmt.Errorf("get task: %w", err)
	}
	if teacherID != taskObj.TeacherID {
		return nil, nil, fmt.Errorf("%w: user is not a task owner", task.ErrForbidden)
	}

	// get teacher
	teacher, err := u.userRepoDB.GetByIDWithProfileShort(taskObj.TeacherID)
	if err != nil {
		return nil, nil, fmt.Errorf("get teacher: %w", err)
	}
	taskObj.Teacher = teacher.Profile

	// get students
	studProfiles, err := u.taskRepoDB.GetTaskStudents(taskID)
	if err != nil {
		return nil, nil, fmt.Errorf("get task students: %w", err)
	}
	return taskObj, studProfiles, nil
}

// Update updates the given task by given ID with the new data.
// It returns the updated task object and updated students linked to the task.
// Allows to update the title, desc, linked students and task files.
func (u *UCTeacher) Update(teacherID, taskID int,
	newData *entity.TaskUpdate) (*entity.Task, []entity.Profile, error) {

	// get task with files
	taskObj, err := u.taskRepoDB.GetByID(taskID)
	if err != nil {
		return nil, nil, fmt.Errorf("get task: %w", err)
	}
	if teacherID != taskObj.TeacherID {
		return nil, nil, fmt.Errorf("%w: user is not a task owner", task.ErrForbidden)
	}

	newData.DelFiles = make(entity.Files, 0, len(newData.DelFilesIDs))
	taskFilesRemains := make(entity.Files, 0, len(taskObj.Files))
	// collect files' metadata to delete them from task
	for _, oldFile := range taskObj.Files {
		if goslices.Contains(newData.DelFilesIDs, oldFile.ID) {
			newData.DelFiles = append(newData.DelFiles, oldFile)
		} else {
			taskFilesRemains = append(taskFilesRemains, oldFile)
		}
	}
	taskObj.Files = taskFilesRemains

	var students []entity.Profile
	if newData.NewFullStudents != nil {
		// set add/del student lists to newData object and
		// get slice of new student profiles
		newData.AddStudents, newData.DelStudents,
			students, err = u.sepNewStudents(taskID, newData.NewFullStudents)
	} else {
		// get actual students list
		students, err = u.taskRepoDB.GetTaskStudents(taskID)
	}
	if err != nil {
		return nil, nil, fmt.Errorf("get students: %w", err)
	}

	if err := u.taskRepoDB.Update(taskID, newData); err != nil {
		return nil, nil, fmt.Errorf("update: %w", err)
	}

	if newData.Title != nil {
		taskObj.Title = *newData.Title
	}
	if newData.Desc != nil {
		taskObj.Desc = *newData.Desc
	}
	// append new files to task files
	taskObj.Files = append(taskObj.Files, newData.AddFiles...)
	// remove deleted files from file system
	newData.DelFiles.Cleanup()

	return taskObj, students, nil
}

// DeleteByID deletes task object by given ID.
func (u *UCTeacher) DeleteByID(userID, taskID int) error {
	// get task info
	taskObj, err := u.taskRepoDB.GetByID(taskID)
	// return nil error if task was not found
	if errors.Is(err, task.ErrNotFound) {
		return nil
	}
	if err != nil {
		return err
	}
	// check that given user (teacher) is a task owner
	if userID != taskObj.TeacherID {
		return fmt.Errorf("%w: user (teacher) is not a task owner", task.ErrForbidden)
	}
	// delete task
	if err := u.taskRepoDB.Delete(taskObj); err != nil {
		return err
	}
	// delete files from the file system
	taskObj.Files.Cleanup()
	return nil
}

// GetMany returns all teacher tasks.
// Search param appends condition to filter tasks by title (substring).
func (u *UCTeacher) GetMany(teacherID int, search string,
	page *entity.Pagination) ([]entity.TaskWithStudents, error) {

	// get tasks
	taskList, err := u.taskRepoDB.GetMany(teacherID, search, page)
	if err != nil {
		return nil, fmt.Errorf("get tasks: %w", err)
	}
	if len(taskList) == 0 {
		return []entity.TaskWithStudents{}, nil
	}

	errChan := make(chan error, len(taskList))
	var wg sync.WaitGroup

	// fan out pattern to get students for each task
	res := make([]entity.TaskWithStudents, len(taskList))
	for idx := range taskList {
		wg.Add(1)
		go func(taskIdx int) {
			defer wg.Done()
			res[taskIdx].Task = &taskList[taskIdx]
			// get students for task
			students, loopErr := u.taskRepoDB.GetTaskStudents(taskList[taskIdx].ID)
			if loopErr != nil {
				errChan <- fmt.Errorf("task %d: %w", taskList[taskIdx].ID, loopErr)
				return
			}
			res[taskIdx].Students = students
		}(idx)
	}

	go func() {
		defer close(errChan)
		wg.Wait()
	}()

	for err := range errChan {
		if err != nil {
			return nil, fmt.Errorf("get students: %w", err)
		}
	}
	return res, nil
}

// sepNewStudents separates students from new students list into add/delete (linked to task) lists.
// It returns both result lists and new student profiles list.
func (u *UCTeacher) sepNewStudents(taskID int, newStudIDs []int) (add []int, del []int,
	newStuds []entity.Profile, err error) {

	// get actual students list before updating task solutions
	oldStuds, err := u.taskRepoDB.GetTaskStudents(taskID)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("old list: %w", err)
	}
	// get user objects of new students (and check them)
	newStudUsers, err := u.getStudentsByIDs(newStudIDs)
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
func (u *UCTeacher) getStudentsByIDs(studentIDs []int) ([]entity.User, error) {
	// get user objects by IDs
	students, err := u.userRepoDB.GetManyWithProfilesShort(studentIDs)
	if err != nil {
		return nil, err
	}
	// check students
	for idx := range students {
		if !students[idx].IsStudent() {
			return nil, fmt.Errorf("%w: user %d: not student",
				task.ErrInvalidData, students[idx].ID)
		}
	}
	return students, nil
}
