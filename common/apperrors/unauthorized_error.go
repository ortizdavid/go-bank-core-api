package apperrors

import (
	"fmt"
	"net/http"
)

type UnauthorizedError struct {
	Message    string
	StatusCode int
}

func NewUnauthorizedError(message string) *UnauthorizedError {
	return &UnauthorizedError{
		Message:    message,
		StatusCode: http.StatusUnauthorized,
	}
}

func (err *UnauthorizedError) Error() string {
	return fmt.Sprintf("Unathorized: %s", err.Message)
}
