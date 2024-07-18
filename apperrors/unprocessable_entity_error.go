package apperrors

import "fmt"

type UnprocessableEntityError struct {
	Message    string
	StatusCode int
}

func NewUnprocessableEntityError(message string) *UnprocessableEntityError {
	return &UnprocessableEntityError{
		Message:    message,
		StatusCode: 422,
	}
}

func (err *UnprocessableEntityError) Error() string {
	return fmt.Sprintf("Unprocessable: %s", err.Message)
}