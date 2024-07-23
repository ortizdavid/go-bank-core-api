package apperrors

import (
	"fmt"
	"net/http"
)

type ForbiddenError struct {
	Message    string
	StatusCode int
}

func (err *ForbiddenError) NewForbiddenError(message string) *ForbiddenError {
	return &ForbiddenError{
		Message:    message,
		StatusCode: http.StatusForbidden,
	}
}

func (err *ForbiddenError) Error() string {
	return fmt.Sprintf("Forbiden: %s", err.Message)
}