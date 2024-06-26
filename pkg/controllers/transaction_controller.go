package controllers

import (
	"net/http"

	"github.com/ortizdavid/go-bank-core-api/config"
	accountRepo "github.com/ortizdavid/go-bank-core-api/pkg/repositories/accounts"
	transactionRepo "github.com/ortizdavid/go-bank-core-api/pkg/repositories/transactions"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TransactionController struct {
	transactionRepository transactionRepo.TransactionRepository
	accountRepository accountRepo.AccountRepository
	logger *zap.Logger
}

func NewTransactionController(db *gorm.DB) *TransactionController {
	return &TransactionController{
		transactionRepository: *transactionRepo.NewTransactionRepository(db),
		accountRepository:  *accountRepo.NewAccountRepository(db),
		logger:             config.NewLogger("transactions.log"),
	}
}

func (*TransactionController) Routes(router *http.ServeMux, db *gorm.DB) {
	
}