package user

import "skadi/backend/internal/app/entity"

// RepositoryDB describes all DB methods for user.
type RepositoryDB interface {
	// CreateUserWithDefaultProfile creates a new user with default profile.
	CreateUserWithDefaultProfile(userObj *entity.User) error
	// CreateUserFull creates a new user with class (if set) and
	// profile for them and fills given structs.
	CreateUserFull(userObj *entity.User) error

	// GetByID returns user by given ID.
	GetByID(id int) (*entity.User, error)
	// GetOneFull returns user with class (if set) and profile with given cond.
	GetOneFull(field string, value any) (*entity.User, error)
	// GetByIDWithProfileShort returns user with short profile (id and fullname only) by given ID.
	GetByIDWithProfileShort(id int) (*entity.User, error)

	// UpdateUser updates old user data to new one (by data ID).
	UpdateUser(data *entity.User) error
	// UpdateProfile updates old user profile to new one (by profile ID).
	// Old user profile contacts (contact and parent contact) will be deleted
	// if they are changed and not used in other profiles.
	UpdateProfile(oldData, newData *entity.Profile) error

	// Delete deletes user object and user profile (by data ID).
	// Also user profile contacts (contact and parent contact) will be deleted
	// if they are not used in other profiles.
	Delete(data *entity.User) error

	// GetByRoles returns user (with class if set and profile) list with given roles.
	// Free param appends condition (if only student role was given) to get class-free students.
	// Search params appends condition to filter users by username and fullname (substring).
	GetByRoles(roles []entity.Role, free bool, search string,
		page *entity.Pagination) ([]entity.User, error)

	// GetManyWithProfilesShort returns users by given IDs with short profiles (ID and fullname).
	GetManyWithProfilesShort(ids []int) ([]entity.User, error)
	// GetProfilesByClass returns short profiles (ID and fullname) linked to the class with given ID.
	GetProfilesShortByClass(classID int) ([]entity.Profile, error)
	// AddDelChanges separates students from new students list
	// into add/delete lists basing on comparation with old students list.
	AddDelChanges(oldStuds, newStuds []entity.Profile) (add, del []int)
}
