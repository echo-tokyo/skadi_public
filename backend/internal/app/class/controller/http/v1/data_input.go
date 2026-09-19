package v1

import "skadi/backend/internal/app/entity"

// @description classBody represents a data with class.
type classBody struct {
	// class name
	Name string `json:"name" validate:"required,max=50" example:"B26-1" maxLength:"50"`
	// teacher id
	TeacherID *int `json:"teacher_id,omitempty" validate:"omitempty,numeric" example:"2"`
	// class schedule
	Schedule *string `json:"schedule,omitempty" validate:"omitempty,max=100" example:"сб 18:00-19:00" maxLength:"100"`
	// student ID list
	StudentIDs []int `json:"students,omitempty" validate:"omitempty"`
}

func (u *classBody) ToEntityClass() *entity.Class {
	if u.TeacherID != nil && *u.TeacherID == 0 {
		u.TeacherID = nil
	}
	if u.Schedule != nil && *u.Schedule == "" {
		u.Schedule = nil
	}
	return &entity.Class{
		Name:      u.Name,
		TeacherID: u.TeacherID,
		Schedule:  u.Schedule,
	}
}

// @description classIDPath represents a data with class ID in path params.
type classIDPath struct {
	// class id
	ID int `params:"id" validate:"required,numeric" example:"2"`
}

// @description updateBody represents a data to update class.
type updateBody struct {
	// new class name
	Name *string `json:"name,omitempty" validate:"omitempty,max=50" example:"F26-2" maxLength:"50"`
	// new teacher id
	TeacherID *int `json:"teacher_id,omitempty" validate:"omitempty,numeric" example:"5"`
	// new class schedule
	Schedule *string `json:"schedule,omitempty" validate:"omitempty,max=100" example:"пн 15:00-16:00" maxLength:"100"`
	// IDs of students (updated list) in the class
	Students []int `json:"students,omitempty" validate:"omitempty"`
}

// @description listClassQuery represents a data with optional query-params to get classes list.
type listClassQuery struct {
	// substring to filter classes by name (case-insensitive)
	Search string `query:"search,omitempty" json:"search" example:"F26"`
	// pagination params
	entity.PaginationQuery
}
