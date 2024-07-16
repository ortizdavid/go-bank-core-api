package controllers

import (
	"net/http"
	"gorm.io/gorm"
	reportRepo "github.com/ortizdavid/go-bank-core-api/core/repositories/reports"
)

type AccountReportController struct {
	repositoryRepository reportRepo.AccountReportRepository
}

func NewAccountReportController(db *gorm.DB) *AccountReportController {
	return &AccountReportController{
		repositoryRepository: *reportRepo.NewAccountReportRepository(db),
	}
}

func (*AccountReportController) RegisterRoutes(router *http.ServeMux, db *gorm.DB) {

}