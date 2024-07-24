package helpers

import (
	"net/http"
	"github.com/ortizdavid/go-bank-core-api/common/apperrors"
	"github.com/ortizdavid/go-nopain/httputils"
)

// handleCustomerError centralizes error handling for customer-related operations
func HandleCustomErrors(w http.ResponseWriter, err error) {
    if err != nil {
        if e, ok := err.(*apperrors.AppError); ok {
            httputils.WriteJsonError(w, e.Error(), e.StatusCode)
            return
        } else {
            httputils.WriteJsonError(w, e.Error(), http.StatusInternalServerError)
            return
        }
    }

}