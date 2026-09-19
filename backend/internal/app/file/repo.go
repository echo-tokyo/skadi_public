package file

import "skadi/backend/internal/app/entity"

// RepositoryDB describes all DB methods for file object.
type RepositoryDB interface {
	// GetByID returns file by the given ID.
	GetByID(id int) (*entity.File, error)
	// TeacherPermit returns nil error if teacher has rights to the given file.
	TeacherPermit(teacherID, fileID int) error
	// StudentPermit returns nil error if student has rights to the given file.
	StudentPermit(studentID, fileID int) error
}
