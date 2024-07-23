package controllers

import (
	"net/http"
	"github.com/ortizdavid/go-bank-core-api/common/config"
	accountRepo "github.com/ortizdavid/go-bank-core-api/core/repositories/accounts"
	customerRepo "github.com/ortizdavid/go-bank-core-api/core/repositories/customers"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AccountController struct {
	repository *accountRepo.AccountRepository
	customerRepository *customerRepo.CustomerRepository
	logger *zap.Logger
}

func NewAccountController(db *gorm.DB) *AccountController {
	return &AccountController{
		repository:  accountRepo.NewAccountRepository(db),
		customerRepository: customerRepo.NewCustomerRepository(db),
		logger:             config.NewLogger("accounts.log"),
	}
}

func (controller *AccountController) RegisterRoutes(router *http.ServeMux) {
	
}