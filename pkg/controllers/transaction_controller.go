package controllers

import (
	"net/http"

	accountRepo "github.com/ortizdavid/go-bank-core-api/pkg/repositories/accounts"
	transactionRepo "github.com/ortizdavid/go-bank-core-api/pkg/repositories/transactions"
)

type TransactionController struct {
	transactionRepository transactionRepo.TransactionRepository
	accountRepository accountRepo.AccountRepository
}

func (TransactionController) Routes(router *http.ServeMux)  {

}