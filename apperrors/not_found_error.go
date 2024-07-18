package apperrors

import "fmt"

type NotFoundError struct {
	Message string
	StatusCode int
}

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{
		Message: message,
		StatusCode: 404,
	}
}

func (err *NotFoundError) Error() string {
	return fmt.Sprintf("NotFound: %s", err.Message)
}