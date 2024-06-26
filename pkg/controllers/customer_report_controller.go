package controllers

import (
	"net/http"
	"gorm.io/gorm"
	reportRepo "github.com/ortizdavid/go-bank-core-api/pkg/repositories/reports"
)

type CustomerReportController struct {
	repositoryRepository reportRepo.CustomerReportRepository
}

func NewCustomerReportController(db *gorm.DB) *CustomerReportController {
	return &CustomerReportController{
		repositoryRepository: *reportRepo.NewCustomerReportRepository(db),
	}
}

func (*CustomerReportController) Routes(router *http.ServeMux, db *gorm.DB) {

}