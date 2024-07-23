package helpers

import (
	"net/http"
	"github.com/ortizdavid/go-bank-core-api/common/apperrors"
	"github.com/ortizdavid/go-nopain/httputils"
)

// handleCustomerError centralizes error handling for customer-related operations
func HandleCustomErrors(w http.ResponseWriter, err error) {
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