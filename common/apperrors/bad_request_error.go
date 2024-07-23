package apperrors

import (
	"fmt"
	"net/http"
)

type BadRequestError struct {
	Message string
	StatusCode int
}

func NewBadRequestError(message string) *BadRequestError {
	return &BadRequestError{
		Message: message,
		StatusCode: http.StatusBadRequest,
	}
}

func (err *BadRequestError) Error() string {
	return fmt.Sprintf("BadRequest: %s", err.Message)
}