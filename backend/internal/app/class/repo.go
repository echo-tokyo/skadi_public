package class

import "skadi/backend/internal/app/entity"

// RepositoryDB describes all DB methods for class.
type RepositoryDB interface {
	// Create creates a new class and fills given struct.
	Create(classObj *entity.Class, studentIDs []int) error
	// GetByIDShort returns class (ID and name only) by given id.
	GetByIDShort(id int) (*entity.Class, error)
	// GetByID returns class by given id.
	GetByID(id int) (*entity.Class, error)
	// Update updates class by given ID with the new data.
	// It returns the updated class object.
	Update(id int, newData *entity.ClassUpdate) error
	// DeleteByID deletes class object by given id.
	DeleteByID(id int) error

	// ListShort returns slice of class objects (IDs and names only).
	// Search param appends condition to filter classes by name (substring).
	ListShort(search string, page *entity.Pagination) ([]entity.Class, error)
	// ListFull returns slice of class objects with full data.
	// Search param appends condition to filter classes by name (substring).
	ListFull(search string, page *entity.Pagination) ([]entity.Class, error)
}
