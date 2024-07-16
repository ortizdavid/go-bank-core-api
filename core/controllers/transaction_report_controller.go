package controllers

import (
	"net/http"
	"gorm.io/gorm"
	reportRepo "github.com/ortizdavid/go-bank-core-api/core/repositories/reports"
)

type TransactionReportController struct {
	repositoryRepository reportRepo.TransactionReportRepository
}

func NewTransactionReportController(db *gorm.DB) *TransactionReportController {
	return &TransactionReportController{
		repositoryRepository: *reportRepo.NewTransactionReportRepository(db),
	}
}

func (*TransactionReportController) RegisterRoutes(router *http.ServeMux, db *gorm.DB) {

}