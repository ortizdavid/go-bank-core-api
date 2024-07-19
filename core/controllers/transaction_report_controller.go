package controllers

import (
	"net/http"
	"gorm.io/gorm"
	"github.com/ortizdavid/go-bank-core-api/core/repositories/reports"
)

type TransactionReportController struct {
	repositoryRepository repositories.TransactionReportRepository
}

func NewTransactionReportController(db *gorm.DB) *TransactionReportController {
	return &TransactionReportController{
		repositoryRepository: *repositories.NewTransactionReportRepository(db),
	}
}

func (*TransactionReportController) RegisterRoutes(router *http.ServeMux) {

}