package controllers

import (
	"net/http"
	"gorm.io/gorm"
	"github.com/ortizdavid/go-bank-core-api/core/repositories/reports"
)

type AccountReportController struct {
	repositoryRepository repositories.AccountReportRepository
}

func NewAccountReportController(db *gorm.DB) *AccountReportController {
	return &AccountReportController{
		repositoryRepository: *repositories.NewAccountReportRepository(db),
	}
}

func (*AccountReportController) RegisterRoutes(router *http.ServeMux) {

}