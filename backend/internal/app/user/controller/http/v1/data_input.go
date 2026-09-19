package v1

import (
	"fmt"
	"slices"

	"skadi/backend/internal/app/entity"
	"skadi/backend/internal/pkg/serialize"
	"skadi/backend/internal/pkg/validator"
)

// _acceptedRoles contains accepted roles for listUserQuery
var _acceptedRoles = []entity.Role{entity.Teacher, entity.Student}

// @description userBody represents a data with user.
type userBody struct {
	// user username
	Username string `json:"username" validate:"required,max=50" example:"user1" maxLength:"50"`
	// user password
	Password string `json:"password" validate:"required,strong-passwd,min=8,max=40" example:"qwerty123" minLength:"8" maxLength:"40"`
	// user role (teacher or student)
	Role entity.Role `json:"role" validate:"required,oneof=teacher student" example:"teacher"`
	// class id (for students)
	ClassID *int `json:"class_id" validate:"omitempty" example:"3"`
	// user profile
	Profile profileBody `json:"profile" validate:"required"`
}

func (u *userBody) ToEntityUser() *entity.User {
	if u.ClassID != nil && *u.ClassID == 0 {
		u.ClassID = nil
	}
	return &entity.User{
		Username: u.Username,
		Password: []byte(u.Password),
		Role:     u.Role,
		ClassID:  u.ClassID,
		Profile:  u.Profile.ToEntityProfile(),
	}
}

// @description profileBody represents a data with user profile.
type profileBody struct {
	// user full name
	FullName string `json:"fullname" validate:"required,max=150" example:"Иванов Василий Петрович" maxLength:"150"`
	// user address
	Address *string `json:"address,omitempty" validate:"omitempty,max=255" example:"3-я ул.Строителей д. 25" maxLength:"255"`
	// extra data (for admin only)
	Extra *string `json:"extra,omitempty" validate:"omitempty" example:"только для админов"`
	// user contact info
	Contact *contactBody `json:"contact,omitempty" validate:"omitempty"`
	// parent contact info (for students)
	ParentContact *contactBody `json:"parent_contact,omitempty" validate:"omitempty"`
}

func (p *profileBody) ToEntityProfile() *entity.Profile {
	profile := &entity.Profile{
		Fullname: p.FullName,
		Address:  p.Address,
		Extra:    p.Extra,
	}
	if p.Contact != nil {
		profile.Contact = p.Contact.ToEntityContact()
	}
	if p.ParentContact != nil {
		profile.ParentContact = p.ParentContact.ToEntityContact()
	}
	return profile
}

// @description contactBody represents a data with profile contact.
type contactBody struct {
	Phone string `json:"phone" validate:"omitempty,max=15" example:"88005553535" maxLength:"15"`
	Email string `json:"email" validate:"omitempty,email,max=50" example:"ivanovvp@gmail.com" maxLength:"50"`
}

func (c *contactBody) ToEntityContact() *entity.Contact {
	if c.Phone == "" && c.Email == "" {
		return nil
	}

	contact := &entity.Contact{}
	if c.Phone != "" {
		contact.Phone = &c.Phone
	}
	if c.Email != "" {
		contact.Email = &c.Email
	}
	return contact
}

// @description userIDPath represents a data with user ID in path params.
type userIDPath struct {
	// user (and profile) id
	ID int `params:"id" validate:"required,numeric" example:"2"`
}

// @description updateBody represents a data to update user and profile.
type updateBody struct {
	// class id (for students)
	ClassID *int `json:"class_id,omitempty" validate:"omitempty,numeric" example:"3"`
	// user password
	Password string `json:"password,omitempty" validate:"omitempty,strong-passwd,min=8,max=40" example:"qwerty123" minLength:"8" maxLength:"40"`
	// user profile
	Profile profileBody `json:"profile" validate:"required"`
}

func (u *updateBody) ToEntityUser() *entity.User {
	if u.ClassID != nil && *u.ClassID == 0 {
		u.ClassID = nil
	}
	// data reshaping
	return &entity.User{
		ClassID:  u.ClassID,
		Password: []byte(u.Password),
		Profile:  u.Profile.ToEntityProfile(),
	}
}

// @description listUserQuery represents a data with optional query-params to get users list.
type listUserQuery struct {
	// user roles (accepted: "teacher", "student")
	Roles []entity.Role `query:"role,omitempty" json:"role" example:"student"`
	// return only class-free students (only if param role=student)
	Free bool `query:"free,omitempty" json:"free" example:"true"`
	// substring to filter users by username and fullname (case-insensitive)
	Search string `query:"search,omitempty" json:"search" example:"иванов"`
	// pagination params
	entity.PaginationQuery
}

// Parse parses [listUserQuery] request data and validates it.
func (u *listUserQuery) Validate(valid validator.Validator) serialize.Validator {
	return func(_ any) error {
		// validate parsed roles list
		for _, role := range u.Roles {
			if !slices.Contains(_acceptedRoles, role) {
				return fmt.Errorf("%w: invalid user role %s", validator.ErrValidate, role)
			}
		}
		// set all roles if no one role specified
		if len(u.Roles) == 0 {
			u.Roles = _acceptedRoles
		}
		// validate parsed data
		if err := valid.Validate(u); err != nil {
			return err
		}
		return nil
	}
}

// @description updatePasswordAdminBody represents a data to update user password by admin.
type updatePasswordAdminBody struct {
	// new password for user
	NewPasswd string `json:"new" validate:"required,strong-passwd,min=8,max=40" example:"ytrewq321" minLength:"8" maxLength:"40"`
}

// @description updatePasswordBody represents a data to update user password by the user himself.
type updatePasswordBody struct {
	// old user password
	OldPasswd string `json:"old" validate:"required,strong-passwd,min=8,max=40" example:"qwerty123" minLength:"8" maxLength:"40"`
	// new user password
	NewPasswd string `json:"new" validate:"required,strong-passwd,min=8,max=40" example:"ytrewq321" minLength:"8" maxLength:"40"`
}
