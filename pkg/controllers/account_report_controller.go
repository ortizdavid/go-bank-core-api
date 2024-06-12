package controllers

import (
	"net/http"
	"gorm.io/gorm"
	reportRepo "github.com/ortizdavid/go-bank-core-api/pkg/repositories/reports"
)

type AccountReportController struct {
	repositoryRepository reportRepo.AccountReportRepository
}

func NewAccountReportController(db *gorm.DB) *AccountReportController {
	return &AccountReportController{
		repositoryRepository: *reportRepo.NewAccountReportRepository(db),
	}
}

func (*AccountReportController) Routes(router *http.ServeMux, db *gorm.DB) {

}