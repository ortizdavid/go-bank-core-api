package controllers

import (
	"net/http"

	"github.com/ortizdavid/go-bank-core-api/config"
	customerRepo "github.com/ortizdavid/go-bank-core-api/pkg/repositories/customers"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CustomerController struct {
	customerRepository customerRepo.CustomerRepository
	logger *zap.Logger
}

func NewCustomerController(db *gorm.DB) *CustomerController {
	return &CustomerController{
		customerRepository: *customerRepo.NewCustomerRepository(db),
		logger:             config.NewLogger("customners.log"),
	}
}

func (*CustomerController) Routes(router *http.ServeMux, db *gorm.DB) {
	
}