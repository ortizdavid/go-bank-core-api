package helpers

import (
	"net/http"
	"github.com/ortizdavid/go-nopain/conversion"
)

type PaginationParams struct {
    CurrentPage int
    Limit       int
}

// getCurrentPageAndLimit extracts current page and limit from URL query parameters.
func GetPaginationParams(r *http.Request) PaginationParams {
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
	return PaginationParams {
		CurrentPage: currentPage,
		Limit:       limit,
	}
}
