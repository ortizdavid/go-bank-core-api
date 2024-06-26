package controllers

import (
	"net/http"
	"gorm.io/gorm"
	reportRepo "github.com/ortizdavid/go-bank-core-api/pkg/repositories/reports"
)

type TransactionReportController struct {
	repositoryRepository reportRepo.TransactionReportRepository
}

func NewTransactionReportController(db *gorm.DB) *TransactionReportController {
	return &TransactionReportController{
		repositoryRepository: *reportRepo.NewTransactionReportRepository(db),
	}
}

func (*TransactionReportController) Routes(router *http.ServeMux, db *gorm.DB) {

}