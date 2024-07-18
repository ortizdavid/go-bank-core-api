package apperrors

import "fmt"

type ConflictError struct {
	Message string
	StatusCode int
}

func NewConflictError(message string) *ConflictError {
	return &ConflictError{
		Message: message,
		StatusCode: 409,
	}
}

func (err *ConflictError) Error() string {
	return fmt.Sprintf("Conflict: %s", err.Message)
}