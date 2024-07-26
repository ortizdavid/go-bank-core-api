package helpers

import (
	"net/http"
	"github.com/ortizdavid/go-bank-core-api/common/apperrors"
	"github.com/ortizdavid/go-nopain/httputils"
)

// HandleHttpErrors centralizes error handling for http-related operations
func HandleHttpErrors(w http.ResponseWriter, err error) {
    if err != nil {
        if e, ok := err.(*apperrors.HttpError); ok {
            httputils.WriteJsonError(w, e.Error(), e.StatusCode)
            return
        } else {
            httputils.WriteJsonError(w, e.Error(), http.StatusInternalServerError)
            return
        }
    }
}