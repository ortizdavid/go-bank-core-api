package apperrors

import "fmt"

type ForbiddenError struct {
	Message    string
	StatusCode int
}

func (err *ForbiddenError) NewForbidenError(message string) *ForbiddenError {
	return &ForbiddenError{
		Message:    message,
		StatusCode: 403,
	}
}

func (err *ForbiddenError) Error() string {
	return fmt.Sprintf("Forbiden: %s", err.Message)
}