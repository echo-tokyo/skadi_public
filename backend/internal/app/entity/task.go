package entity

import (
	"time"
)

// Task represents a task data.
type Task struct {
	// task id
	ID int `gorm:"primaryKey;autoIncrement" json:"id" validate:"required"`
	// task title
	Title string `json:"title" validate:"required"`
	// task description
	Desc string `gorm:"column:description" json:"description,omitempty" validate:"omitempty"`
	// task teacher id
	TeacherID int `json:"-"`
	// task creating datetime
	CreatedAt time.Time `json:"-"`

	// teacher object
	Teacher     *Profile `gorm:"-" json:"teacher,omitempty" validate:"omitempty"`
	TeacherUser *User    `gorm:"foreignKey:TeacherID;references:ID" json:"-"`
	// task files
	Files Files `gorm:"many2many:task_file;foreignKey:ID;References:ID" json:"files,omitempty" validate:"omitempty"`
}

// TableName determines DB table name for the task object.
func (*Task) TableName() string {
	return "task"
}

// TaskWithStudents represents a task data with students linked to it.
type TaskWithStudents struct {
	// task object
	Task *Task `json:"task" validate:"required"`
	// students solving this task
	Students []Profile `json:"students,omitempty" validate:"omitempty"`
}

// TaskUpdate represents a data to update task.
type TaskUpdate struct {
	// new task title
	Title *string
	// new task description
	Desc *string

	// IDs of new students to completely replace old students
	NewFullStudents []int
	// IDs of students to issue the task solutions for them
	AddStudents []int
	// IDs of students to delete their task solutions
	DelStudents []int

	// List of files to append to the task
	AddFiles Files
	// IDs of files to delete from the task
	DelFilesIDs []int
	// List of files to delete from the task
	DelFiles Files
}

// ToUpdatesMap generates map to update task object.
func (t *TaskUpdate) ToUpdatesMap() map[string]any {
	updates := make(map[string]any)
	// set new title
	if t.Title != nil {
		updates["title"] = *t.Title
	}
	// set new desctiption
	if t.Desc != nil {
		updates["description"] = t.Desc
	}
	return updates
}
