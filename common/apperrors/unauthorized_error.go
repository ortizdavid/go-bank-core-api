package apperrors

import (
	"net/http"
)

func NewUnauthorizedError(message string) *AppError {
	return &AppError{
		Message:    message,
		StatusCode: http.StatusUnauthorized,
	}
}
