package apperrors

import "net/http"

type AppError struct {
	Message    string
	StatusCode int
}

func (err *AppError) Error() string {
	return http.StatusText(err.StatusCode) + ": " + err.Message
}