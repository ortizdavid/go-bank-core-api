package controllers

import (
	"net/http"
	"github.com/ortizdavid/go-bank-core-api/apperrors"
	"github.com/ortizdavid/go-nopain/conversion"
	"github.com/ortizdavid/go-nopain/httputils"
)

// getCurrentPageAndLimit extracts current page and limit from URL query parameters.
func getCurrentPageAndLimit(r *http.Request) (int, int) {
	currentStr := r.URL.Query().Get("current_page")
	limitStr := r.URL.Query().Get("limit")
	if currentStr == "" { 
		currentStr = "0" // Default current page
	}
	if limitStr == "" { 
		limitStr = "10" // Default limit per page
 	}
	currentPage := conversion.StringToInt(currentStr)
	limit := conversion.StringToInt(limitStr)
	return currentPage, limit
}

// handleCustomerError centralizes error handling for customer-related operations
func handleCustomErrors(w http.ResponseWriter, err error) {
    switch e := err.(type) {
    case *apperrors.BadRequestError:
        // Handle bad request errors
        httputils.WriteJsonError(w, e.Error(), e.StatusCode)
    case *apperrors.NotFoundError:
        // Handle not found errors
        httputils.WriteJsonError(w, e.Error(), e.StatusCode)
    case *apperrors.ConflictError:
        // Handle conflict errors
        httputils.WriteJsonError(w, e.Error(), e.StatusCode)
    case *apperrors.UnauthorizedError:
        // Handle unauthorized errors
        httputils.WriteJsonError(w, e.Error(), e.StatusCode)
    case *apperrors.ForbiddenError:
        // Handle forbidden errors
        httputils.WriteJsonError(w, e.Error(), e.StatusCode)
    case *apperrors.UnprocessableEntityError:
        // Handle unprocessable entity errors
        httputils.WriteJsonError(w, e.Error(), e.StatusCode)
    default:
        // Handle internal server errors for unhandled errors
        httputils.WriteJsonError(w, err.Error(), http.StatusInternalServerError)
    }
}