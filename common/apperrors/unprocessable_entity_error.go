package apperrors

import (
	"fmt"
	"net/http"
)

type UnprocessableEntityError struct {
	Message    string
	StatusCode int
}

func NewUnprocessableEntityError(message string) *UnprocessableEntityError {
	return &UnprocessableEntityError{
		Message:    message,
		StatusCode: http.StatusUnprocessableEntity,
	}
}

func (err *UnprocessableEntityError) Error() string {
	return fmt.Sprintf("Unprocessable: %s", err.Message)
}
