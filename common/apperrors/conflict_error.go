package apperrors

import (
	"net/http"
)

func NewConflictError(message string) *AppError {
	return &AppError{
		Message: message,
		StatusCode: http.StatusConflict,
	}
}
