package entity

// Role represents a user role.
type Role string

var (
	Admin   Role = "admin"   // admin role for user
	Teacher Role = "teacher" // teacher role for user
	Student Role = "student" // student role for user
)

// NewRoleFromString returns role suitable to the given string role.
// By default returns Student role.
func NewRoleFromString(raw string) Role {
	switch raw {
	case string(Admin):
		return Admin
	case string(Teacher):
		return Teacher
	default:
		return Student
	}
}
