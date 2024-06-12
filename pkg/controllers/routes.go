package controllers

import "net/http"

func RegisterRoutes(router *http.ServeMux) {
	CustomerController{}.Routes(router)
	AccountController{}.Routes(router)
	TransactionController{}.Routes(router)
}