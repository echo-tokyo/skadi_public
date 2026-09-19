package entity

// Class represents a student class data.
type Class struct {
	// class id
	ID int `gorm:"primaryKey" json:"id" validate:"required"`
	// class name
	Name string `json:"name" validate:"required"`
	// teacher id
	TeacherID *int `json:"-"`
	// class schedule
	Schedule *string `json:"schedule,omitempty" validate:"omitempty"`

	// teacher object
	Teacher     *Profile `gorm:"-" json:"teacher,omitempty" validate:"omitempty"`
	TeacherUser *User    `gorm:"foreignKey:TeacherID;references:ID" json:"-"`
	// student objects list
	Students []Profile `gorm:"-" json:"students,omitempty" validate:"omitempty"`
}

// TableName determines DB table name for the class object.
func (*Class) TableName() string {
	return "class"
}

// ClassUpdate represents a data to update class.
type ClassUpdate struct {
	// new class name
	Name *string
	// new teacher id
	TeacherID *int
	// new class schedule
	Schedule *string
	// IDs of new students to completely replace old students
	NewFullStudents []int
	// IDs of students to add them to the class
	AddStudents []int
	// IDs of students to delete them from the class
	DelStudents []int
}

// ToUpdatesMap generates map to update class object.
func (c *ClassUpdate) ToUpdatesMap() map[string]any {
	updates := make(map[string]any)
	// set new name
	if c.Name != nil {
		updates["name"] = *c.Name
	}
	// set new schedule
	if c.Schedule != nil {
		updates["schedule"] = c.Schedule
		if *c.Schedule == "" {
			updates["schedule"] = nil
		}
	}
	// set new teacher ID
	if c.TeacherID != nil {
		updates["teacher_id"] = c.TeacherID
		if *c.TeacherID == 0 {
			updates["teacher_id"] = nil
		}
	}
	return updates
}
