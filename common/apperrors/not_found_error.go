package apperrors

import (
	"fmt"
	"net/http"
)

type NotFoundError struct {
	Message string
	StatusCode int
}

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{
		Message: message,
		StatusCode: http.StatusNotFound,
	}
}

func (err *NotFoundError) Error() string {
	return fmt.Sprintf("NotFound: %s", err.Message)
}