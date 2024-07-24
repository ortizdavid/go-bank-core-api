package apperrors

import (
	"net/http"
)

func NewBadRequestError(message string) *AppError {
	return &AppError{
		Message: message,
		StatusCode: http.StatusBadRequest,
	}
}