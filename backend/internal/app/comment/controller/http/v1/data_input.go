package v1

import "skadi/backend/internal/app/entity"

// @description commentBody represents a data with comment.
type commentBody struct {
	// text message
	Message string `json:"message" validate:"required" example:"Добрый день. Подскажите, пожалуйста..."`
}

// @description solutionIDPath represents a data with solution ID in path params.
type solutionIDPath struct {
	// solution id
	ID int `params:"id" validate:"required" example:"2"`
}

// @description listCommentQuery represents a data with optional query-params to get comment list.
type listCommentQuery struct {
	// pagination params
	entity.PaginationQuery
}
