package controllers

import (
	"net/http"
	"gorm.io/gorm"
)

func RegisterRoutes(router *http.ServeMux, db *gorm.DB) {
	ApiRootController{}.RegisterRoutes(router)
	NewCustomerController(db).RegisterRoutes(router)
	NewAccountController(db).RegisterRoutes(router)
	NewTransactionController(db).RegisterRoutes(router)
	NewCustomerReportController(db).RegisterRoutes(router)
	NewAccountReportController(db).RegisterRoutes(router)
	NewTransactionReportController(db).RegisterRoutes(router)
}