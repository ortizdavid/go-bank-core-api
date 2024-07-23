package controllers

import (
	"net/http"
	"github.com/ortizdavid/go-bank-core-api/common/config"
	accountRepo "github.com/ortizdavid/go-bank-core-api/core/repositories/accounts"
	transactionRepo "github.com/ortizdavid/go-bank-core-api/core/repositories/transactions"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TransactionController struct {
	repository *transactionRepo.TransactionRepository
	accountRepository *accountRepo.AccountRepository
	logger *zap.Logger
}

func NewTransactionController(db *gorm.DB) *TransactionController {
	return &TransactionController{
		repository: transactionRepo.NewTransactionRepository(db),
		accountRepository:  accountRepo.NewAccountRepository(db),
		logger:             config.NewLogger("transactions.log"),
	}
}

func (*TransactionController) RegisterRoutes(router *http.ServeMux) {
	
}