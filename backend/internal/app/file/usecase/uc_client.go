// Package usecase contains file.UsecaseClient implementation.
package usecase

import (
	"skadi/backend/config"
	"skadi/backend/internal/app/entity"
	"skadi/backend/internal/app/file"
)

// Ensure UCClient implements interfaces.
var _ file.UsecaseClient = (*UCClient)(nil)

// UCClient represents a file usecase for teacher and student.
// It implements the [file.UsecaseClient] interface.
type UCClient struct {
	cfg        *config.Config
	fileRepoDB file.RepositoryDB
}

// NewUCClient returns a new instance of [UCClient].
func NewUCClient(cfg *config.Config, fileRepoDB file.RepositoryDB) *UCClient {
	return &UCClient{
		cfg:        cfg,
		fileRepoDB: fileRepoDB,
	}
}

// GetByID returns file metadata by the given ID.
func (u *UCClient) GetByID(fileID int, userClaims *entity.UserClaims) (*entity.File, error) {
	// get file
	fileObj, err := u.fileRepoDB.GetByID(fileID)
	if err != nil {
		return nil, err
	}

	// choose func to check file permission
	var checkPermit func(teacherID int, fileID int) error
	if userClaims.IsTeacher() {
		checkPermit = u.fileRepoDB.TeacherPermit
	} else if userClaims.IsStudent() {
		checkPermit = u.fileRepoDB.StudentPermit
	}
	// check
	if err := checkPermit(userClaims.ID, fileID); err != nil {
		return nil, err
	}
	return fileObj, nil
}
