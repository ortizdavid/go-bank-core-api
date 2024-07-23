package apperrors

import (
	"fmt"
	"net/http"
)

type ConflictError struct {
	Message string
	StatusCode int
}

func NewConflictError(message string) *ConflictError {
	return &ConflictError{
		Message: message,
		StatusCode: http.StatusConflict,
	}
}

func (err *ConflictError) Error() string {
	return fmt.Sprintf("Conflict: %s", err.Message)
}