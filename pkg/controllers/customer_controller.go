package controllers

import (
	"net/http"

	customerRepo "github.com/ortizdavid/go-bank-core-api/pkg/repositories/customers"
)

type CustomerController struct {
	customerRepository customerRepo.CustomerRepository
}

func (CustomerController) Routes(router *http.ServeMux)  {

}