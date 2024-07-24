package apperrors

import (
	"net/http"
)

func NewForbiddenError(message string) *AppError{
	return &AppError{
		Message:    message,
		StatusCode: http.StatusForbidden,
	}
}
