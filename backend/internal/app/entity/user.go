// Package entity contains all app entities used in the app's core logic.
package entity

import "time"

// User represents an app user with main user data.
type User struct {
	// user id
	ID int `gorm:"primaryKey;autoIncrement" json:"id" validate:"required"`
	// user username
	Username string `json:"username" validate:"required"`
	// user password hash
	Password []byte `json:"-"`
	// admin, teacher or student
	Role Role `json:"role" validate:"required"`
	// user creating datetime
	CreatedAt time.Time `json:"-"`
	// class id (for students)
	ClassID *int `json:"-"`

	// user profile
	Profile *Profile `gorm:"foreignKey:ID" json:"profile" validate:"required"`
	// student's class
	Class *Class `gorm:"foreignKey:ClassID" json:"class,omitempty" validate:"omitempty"`
}

// TableName determines DB table name for the user object.
func (*User) TableName() string {
	return "user"
}

// IsAdmin returns true if the user role is Admin.
func (u *User) IsAdmin() bool {
	return u.Role == Admin
}

// IsTeacher returns true if the user role is Teacher.
func (u *User) IsTeacher() bool {
	return u.Role == Teacher
}

// IsStudent returns true if the user role is Student.
func (u *User) IsStudent() bool {
	return u.Role == Student
}

// UserClaims represents a claims with user data for JWT-tokens.
type UserClaims struct {
	// user ID
	ID int `json:"id"`
	// admin, teacher or student
	Role Role `json:"role"`
}

// IsAdmin returns true if the user role is Admin.
func (u UserClaims) IsAdmin() bool {
	return u.Role == Admin
}

// IsTeacher returns true if the user role is Teacher.
func (u UserClaims) IsTeacher() bool {
	return u.Role == Teacher
}

// IsStudent returns true if the user role is Student.
func (u UserClaims) IsStudent() bool {
	return u.Role == Student
}

// Token represents a user token pair (access and refresh).
type Token struct {
	// access token
	Access string
	// refresh token
	Refresh string
}

// UserWithToken is a user object and a token object.
type UserWithToken struct {
	User  *User
	Token *Token `json:"-"`
}
