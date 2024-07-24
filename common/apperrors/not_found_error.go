package apperrors

import (
	"net/http"
)

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Message: message,
		StatusCode: http.StatusNotFound,
	}
}

