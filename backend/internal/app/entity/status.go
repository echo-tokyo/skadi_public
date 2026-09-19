package entity

// Status represents a solution status.
type Status struct {
	// status id
	ID *int `gorm:"primaryKey;autoIncrement" json:"id,omitempty" validate:"omitempty"`
	// status name
	Name string `json:"name" validate:"required"`
}

// TableName determines DB table name for the status object.
func (*Status) TableName() string {
	return "status"
}
