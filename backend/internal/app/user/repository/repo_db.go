// Package repository contains auth.RepositoryDB implementation.
package repository

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"skadi/backend/internal/app/entity"
	"skadi/backend/internal/app/user"
)

const (
	_defaultFullname = "Your name" // default fullname for user profile

	_preloadClass         = "Class"                 // object field name
	_preloadProfile       = "Profile"               // object field name
	_preloadContact       = "Profile.Contact"       // object field name
	_preloadParentContact = "Profile.ParentContact" // object field name
	_tableUser            = "User"                  // table name

	_fieldID       = "id"       // table field name
	_fieldName     = "name"     // table field name
	_fieldFullname = "fullname" // table field name
)

// Ensure RepoDB implements interface.
var _ user.RepositoryDB = (*RepoDB)(nil)

// RepoDB is a user DB repo.
// It implements the [user.RepositoryDB] interface.
type RepoDB struct {
	dbStorage *gorm.DB
}

// NewRepoDB returns a new instance of [RepoDB].
func NewRepoDB(dbStorage *gorm.DB) *RepoDB {
	return &RepoDB{
		dbStorage: dbStorage,
	}
}

// CreateUserWithDefaultProfile creates new user with default profile.
func (r *RepoDB) CreateUserWithDefaultProfile(userObj *entity.User) error {
	return r.dbStorage.Transaction(func(tx *gorm.DB) error {
		// create user
		err := r.dbStorage.Create(userObj).Error
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			// user with such username already exists
			return fmt.Errorf("user with username: %w: %s", user.ErrAlreadyExists, err.Error())
		}

		userObj.Profile = &entity.Profile{
			ID:       &userObj.ID,
			Fullname: _defaultFullname,
		}
		// create user profile record
		if err = tx.Create(userObj.Profile).Error; err != nil {
			return fmt.Errorf("profile: %w", err)
		}
		return err // err OR nil
	})
}

// CreateUserFull creates a new user with class (if set) and
// profile for them and fills given structs.
func (r *RepoDB) CreateUserFull(userObj *entity.User) error {
	return r.dbStorage.Transaction(func(tx *gorm.DB) error {
		// create user record
		err := tx.Omit(_preloadProfile).Create(userObj).Error
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			// user with such username already exists
			return fmt.Errorf("user with username: %w: %s", user.ErrAlreadyExists, err.Error())
		}
		if errors.Is(err, gorm.ErrForeignKeyViolated) {
			// class with given ID is not found
			return fmt.Errorf("user: %w: class is not found", user.ErrInvalidData)
		}
		if err != nil {
			return err
		}
		// set profile id
		userObj.Profile.ID = &userObj.ID

		// find or create contact and parent contact
		if err := findOrCreateContacts(tx, userObj.Profile); err != nil {
			return fmt.Errorf("profile: %w", err)
		}
		// create user profile record
		if err = tx.Create(userObj.Profile).Error; err != nil {
			return fmt.Errorf("profile: %w", err)
		}
		// set nil to skip ID field serialization
		userObj.Profile.ID = nil

		// get class info (ID and name only)
		if userObj.ClassID != nil {
			err = tx.Select(_fieldID, _fieldName).
				Where(*userObj.ClassID).
				First(&userObj.Class).Error
			if err != nil {
				return fmt.Errorf("class: %w", err)
			}
		}
		return nil
	})
}

// GetByID returns user by given id.
func (r *RepoDB) GetByID(id int) (*entity.User, error) {
	var userObj entity.User
	err := r.dbStorage.Where(id).
		First(&userObj).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// user object with such id not found
		return nil, fmt.Errorf("user with id: %w: %s", user.ErrNotFound, err.Error())
	}
	return &userObj, err // err OR nil
}

// GetOneFull returns user with class (if set) and profile with given cond.
func (r *RepoDB) GetOneFull(field string, value any) (*entity.User, error) {
	var userObj entity.User
	err := r.dbStorage.
		Preload(_preloadClass, func(db *gorm.DB) *gorm.DB {
			return db.Select(_fieldID, _fieldName) // preload only ID and name
		}).
		Preload(_preloadProfile).
		Preload(_preloadContact).
		Preload(_preloadParentContact).
		Where(field+" = ?", value).First(&userObj).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// user object not found
		return nil, fmt.Errorf("user by %s: %w: %s", field, user.ErrNotFound, err.Error())
	}
	// set nil to skip ID field serialization
	userObj.Profile.ID = nil
	return &userObj, err // err OR nil
}

// GetByIDWithProfileShort returns user with short profile (id and fullname only) by given ID.
func (r *RepoDB) GetByIDWithProfileShort(id int) (*entity.User, error) {
	var userObj entity.User
	err := r.dbStorage.
		Preload(_preloadProfile, func(db *gorm.DB) *gorm.DB {
			return db.Select(_fieldID, _fieldFullname) // preload only ID and fullname
		}).
		Where(id).First(&userObj).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// user object with such id not found
		return nil, fmt.Errorf("user with id: %w: %s", user.ErrNotFound, err.Error())
	}
	return &userObj, err // err OR nil
}

// UpdateUser updates old user data to new one (by data ID).
func (r *RepoDB) UpdateUser(data *entity.User) error {
	// collect updates to map
	updates := make(map[string]any, 0)
	updates["class_id"] = data.ClassID
	if len(data.Password) > 0 {
		updates["password"] = data.Password
	}

	if len(updates) == 0 {
		return nil
	}
	// update user
	err := r.dbStorage.Model(&entity.User{}).
		Where("id = ?", data.ID).
		Updates(updates).Error
	if errors.Is(err, gorm.ErrForeignKeyViolated) {
		// class with given ID is not found
		return fmt.Errorf("user: %w: class is not found", user.ErrInvalidData)
	}
	return err
}

// UpdateProfile updates old user profile to new one (by profile ID).
// Old user profile contacts (contact and parent contact) will be deleted
// if they are changed and not used in other profiles.
func (r *RepoDB) UpdateProfile(oldData, newData *entity.Profile) error {
	return r.dbStorage.Transaction(func(tx *gorm.DB) error {
		// find or create contact and parent contact
		if err := findOrCreateContacts(tx, newData); err != nil {
			return err
		}
		// update user profile record
		if err := tx.Save(newData).Error; err != nil {
			return err
		}

		// delete unused contacts
		if err := deleteUnusedContact(tx, oldData.Contact, newData.Contact); err != nil {
			return fmt.Errorf("delete contact: %w", err)
		}
		err := deleteUnusedContact(tx, oldData.ParentContact, newData.ParentContact)
		if err != nil {
			return fmt.Errorf("delete parent contact: %w", err)
		}
		return nil
	})
}

// Delete deletes user object and user profile (by data ID).
// Also user profile contacts (contact and parent contact) will be deleted
// if they are not used in other profiles.
func (r *RepoDB) Delete(data *entity.User) error {
	return r.dbStorage.Transaction(func(tx *gorm.DB) error {
		// delete user (cascade with profile)
		if err := tx.Delete(&entity.User{}, data.ID).Error; err != nil {
			return err
		}
		// delete unused contacts (if user has a profile)
		if data.Profile != nil {
			if err := deleteUnusedContact(tx, data.Profile.Contact, nil); err != nil {
				return fmt.Errorf("delete contact: %w", err)
			}
			if err := deleteUnusedContact(tx, data.Profile.ParentContact, nil); err != nil {
				return fmt.Errorf("delete parent contact: %w", err)
			}
		}
		return nil
	})
}

// GetByRoles returns user (with class if set and profile) list with given roles.
// Free param appends condition (if only student role was given) to get class-free students.
// Search params appends condition to filter users by username and fullname (substring).
func (r *RepoDB) GetByRoles(roles []entity.Role, free bool, search string,
	page *entity.Pagination) ([]entity.User, error) {

	if len(roles) == 0 {
		return nil, errors.New("no one role specified")
	}
	var userList []entity.User
	// create query
	query := r.dbStorage.
		Model(&entity.User{}).
		Preload(_preloadClass, func(db *gorm.DB) *gorm.DB {
			return db.Select(_fieldID, _fieldName) // preload only ID and name
		}).
		Preload(_preloadProfile).
		Preload(_preloadContact).
		Preload(_preloadParentContact).
		Joins("INNER JOIN profile ON user.id = profile.id").
		Where("role IN ?", roles)
	if free {
		query = query.Where("class_id IS NULL")
	}
	if search != "" {
		query = query.Where("username REGEXP ? OR fullname REGEXP ?", search, search)
	}
	query = query.Order("id DESC")
	// apply pagination if it's not nil
	if page != nil {
		page.CountTotal(query)
		query = page.Query(query)
	}
	// execute query
	if err := query.Find(&userList).Error; err != nil {
		return nil, err
	}

	// set nil to skip ID field serialization
	for idx := range userList {
		userList[idx].Profile.ID = nil
	}
	return userList, nil
}

// GetManyWithProfilesShort returns users with by IDs with short profiles (ID and fullname).
func (r *RepoDB) GetManyWithProfilesShort(ids []int) ([]entity.User, error) {
	var users []entity.User
	err := r.dbStorage.
		Preload(_preloadProfile, func(db *gorm.DB) *gorm.DB {
			return db.Select(_fieldID, _fieldFullname) // preload only ID and fullname
		}).
		Where("id IN ?", ids).
		Find(&users).Error

	return users, err // err OR nil
}

// GetProfilesByClass returns short profiles (ID and fullname) linked to the class with given ID.
func (r *RepoDB) GetProfilesShortByClass(classID int) ([]entity.Profile, error) {
	var profiles []entity.Profile
	err := r.dbStorage.
		Select("profile.id", _fieldFullname).
		Joins("INNER JOIN user ON user.id=profile.id").
		Where("user.class_id = ?", classID).
		Find(&profiles).Error
	return profiles, err // err OR nil
}

// AddDelChanges separates students from new students list
// into add/delete lists basing on comparation with old students list.
func (r *RepoDB) AddDelChanges(oldStuds, newStuds []entity.Profile) (add, del []int) {
	// collect old students' IDs
	oldStudIDs := make(map[int]struct{}, len(oldStuds))
	for _, oldStud := range oldStuds {
		oldStudIDs[*oldStud.ID] = struct{}{}
	}

	var ok bool
	for _, newStud := range newStuds {
		// del record from map if ID is in the both student slices (old and new)
		if _, ok = oldStudIDs[*newStud.ID]; ok {
			delete(oldStudIDs, *newStud.ID)
		} else {
			// append ID to add-list
			add = append(add, *newStud.ID)
		}
	}
	// del IDs that is in the old students slice only
	for oldStudID := range oldStudIDs {
		del = append(del, oldStudID)
	}
	return add, del
}

// findOrCreateContacts finds or creates contact and parent contact info.
func findOrCreateContacts(tx *gorm.DB, profile *entity.Profile) error {
	// find or create contact record
	if profile.Contact != nil {
		if err := findOrCreateContact(tx, profile.Contact); err != nil {
			return fmt.Errorf("contact: %w", err)
		}
		profile.ContactID = &profile.Contact.ID
	}

	// find or create parent contact record
	if profile.ParentContact != nil {
		if err := findOrCreateContact(tx, profile.ParentContact); err != nil {
			return fmt.Errorf("parent contact: %w", err)
		}
		profile.ParentContactID = &profile.ParentContact.ID
	}
	return nil
}

// findOrCreateContacts finds or creates contact object.
func findOrCreateContact(tx *gorm.DB, contact *entity.Contact) error {
	err := tx.Where("email = ?", contact.Email).
		Where("phone = ?", contact.Phone).
		FirstOrCreate(&contact).Error
	if err != nil {
		return fmt.Errorf("find or create: %w", err)
	}
	return nil
}

// deleteUnusedContact deletes old contact record if it's not in use.
func deleteUnusedContact(tx *gorm.DB, oldCont, newCont *entity.Contact) error {
	if oldCont == nil || newCont == oldCont {
		return nil
	}

	var usageAmount int64
	// check both contact_id and parent_contact_id fields
	err := tx.Model(&entity.Profile{}).
		Where("contact_id = ?", oldCont.ID).
		Or("parent_contact_id = ?", oldCont.ID).
		Count(&usageAmount).Error
	if err != nil {
		return fmt.Errorf("count usage: id %v: %w", oldCont.ID, err)
	}

	// skip deletion if contact is still in use
	if usageAmount != 0 {
		return nil
	}
	// delete contact record
	err = tx.Delete(&entity.Contact{}, oldCont.ID).Error
	if err != nil {
		return fmt.Errorf("delete: id %v: %w", oldCont.ID, err)
	}
	return nil
}
