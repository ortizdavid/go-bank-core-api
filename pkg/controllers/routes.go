package controllers

import (
	"net/http"

	"gorm.io/gorm"
)

func RegisterRoutes(router *http.ServeMux, db *gorm.DB) {
	
	customer := NewCustomerController(db)
	account := NewAccountController(db)
	transaction := NewTransactionController(db)
	customerReport := NewCustomerReportController(db)
	accountReport := NewAccountReportController(db)
	transactionReport := NewTransactionReportController(db)

	//-- Routes regristration
	customer.Routes(router, db)
	account.Routes(router, db)
	transaction.Routes(router, db)
	customerReport.Routes(router, db)
	accountReport.Routes(router, db)
	transactionReport.Routes(router, db)
}