// Package HTTPError provides interface for HTTP errors with
// custom error code and user-friendly message.
package httperror

import "fmt"

// HTTPError represents a HTTP-error with status code and user-friendly message.
type HTTPError struct {
	CauseErr   error
	StatusCode int
	Message    string
}

// Error implements an error interface.
func (h *HTTPError) Error() string {
	return fmt.Sprintf("code %d: %s", h.StatusCode, h.CauseErr.Error())
}
