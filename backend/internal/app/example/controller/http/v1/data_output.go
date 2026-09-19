package v1

import "skadi/backend/internal/app/entity"

// @description exampleData represents an output data for example endpoints.
type exampleData struct {
	// handler name
	Handler string `json:"handler" example:"student"`
	// endpoint access description
	Access string `json:"access" example:"student only"`
	// user claims (for endpoints with auth restriction)
	UserClaims *entity.UserClaims `json:"user_claims,omitempty"`
}
