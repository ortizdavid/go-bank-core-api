package controllers

import (
	"net/http"
	"gorm.io/gorm"
	"github.com/ortizdavid/go-bank-core-api/core/repositories/reports"
)

type CustomerReportController struct {
	repositoryRepository repositories.CustomerReportRepository
}

func NewCustomerReportController(db *gorm.DB) *CustomerReportController {
	return &CustomerReportController{
		repositoryRepository: *repositories.NewCustomerReportRepository(db),
	}
}

func (*CustomerReportController) RegisterRoutes(router *http.ServeMux) {

}