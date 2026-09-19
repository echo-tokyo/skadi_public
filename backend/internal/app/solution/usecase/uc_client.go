// Package usecase contains solution.UsecaseTeacher, solution.UsecaseStudent and
// solution.UsecaseClient implementations.
package usecase

import (
	"fmt"

	"skadi/backend/config"
	"skadi/backend/internal/app/entity"
	"skadi/backend/internal/app/solution"
	"skadi/backend/internal/app/task"
)

// Ensure UCClient implements interfaces.
var _ solution.UsecaseClient = (*UCClient)(nil)

// UCClient represents a task usecase for client.
// It implements the [task.UsecaseClient] interface.
type UCClient struct {
	cfg        *config.Config
	solRepoDB  solution.RepositoryDB
	taskRepoDB task.RepositoryDB
}

// NewUCClient returns a new instance of [UCClient].
func NewUCClient(cfg *config.Config, solRepoDB solution.RepositoryDB,
	taskRepoDB task.RepositoryDB) *UCClient {

	return &UCClient{
		cfg:        cfg,
		solRepoDB:  solRepoDB,
		taskRepoDB: taskRepoDB,
	}
}

// GetByID returns a full solution info and all students linked to the solution task.
func (u *UCClient) GetByIDFull(solutionID int,
	userClaims *entity.UserClaims) (*entity.Solution, []entity.Profile, error) {

	sol, err := u.solRepoDB.GetByIDFull(solutionID)
	if err != nil {
		return nil, nil, fmt.Errorf("get solution: %w", err)
	}

	// role checks
	if userClaims.IsTeacher() && sol.Task.TeacherID != userClaims.ID {
		return nil, nil, fmt.Errorf("%w: user (teacher) is not a task owner", solution.ErrForbidden)
	}
	if userClaims.IsStudent() && sol.StudentID != userClaims.ID {
		return nil, nil, fmt.Errorf("%w: user (stud) is not a solution owner", solution.ErrForbidden)
	}

	studProfiles, err := u.taskRepoDB.GetTaskStudents(sol.TaskID)
	if err != nil {
		return nil, nil, fmt.Errorf("get task students: %w", err)
	}
	return sol, studProfiles, nil
}
