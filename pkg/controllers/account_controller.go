package controllers

import (
	"net/http"

	accountRepo "github.com/ortizdavid/go-bank-core-api/pkg/repositories/accounts"
	customerRepo "github.com/ortizdavid/go-bank-core-api/pkg/repositories/customers"
)

type AccountController struct {
	accountRepository accountRepo.AccountRepository
	customerRepository customerRepo.CustomerRepository
}

func (AccountController) Routes(router *http.ServeMux)  {

}