package apperrors

import (
	"net/http"
)

func NewUnprocessableEntityError(message string) *AppError {
	return &AppError{
		Message:    message,
		StatusCode: http.StatusUnprocessableEntity,
	}
}
