package apperrors

import "fmt"

type UnauthorizedError struct {
	Message    string
	StatusCode int
}

func NewUnauthorizedError(message string) *UnauthorizedError {
	return &UnauthorizedError{
		Message:    message,
		StatusCode: 401,
	}
}

func (err *UnauthorizedError) Error() string {
	return fmt.Sprintf("Unathorized: %s", err.Message)
}