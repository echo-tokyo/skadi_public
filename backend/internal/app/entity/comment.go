package entity

import "time"

// Comment represents an info about solution comment.
type Comment struct {
	// comment ID
	ID int `json:"id" validate:"required"`
	// solution ID
	SolutionID int `json:"-"`
	// role (student or teacher)
	Role Role `json:"role" validate:"required"`
	// message text
	Message string `json:"message" validate:"required"`
	// datetime the message was created
	CreatedAt time.Time `json:"created_at" validate:"required"`
}

// TableName determines DB table name for the comment object.
func (*Comment) TableName() string {
	return "comment"
}
