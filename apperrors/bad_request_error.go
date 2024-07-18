package apperrors

import "fmt"

type BadRequestError struct {
	Message string
	StatusCode int
}

func NewBadRequestError(message string) *BadRequestError {
	return &BadRequestError{
		Message: message,
		StatusCode: 400,
	}
}

func (err *BadRequestError) Error() string {
	return fmt.Sprintf("BadRequest: %s", err.Message)
}