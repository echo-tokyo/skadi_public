package v1

// @description fileIDPath represents a data with file ID in path params.
type fileIDPath struct {
	// file id
	ID int `params:"id" validate:"required" example:"2"`
}
