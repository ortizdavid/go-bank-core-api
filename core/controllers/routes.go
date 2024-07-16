package controllers

import (
	"net/http"
	"gorm.io/gorm"
)

func RegisterRoutes(router *http.ServeMux, db *gorm.DB) {
	ApiRootController{}.RegisterRoutes(router)
	NewCustomerController(db).RegisterRoutes(router, db)
	NewAccountController(db).RegisterRoutes(router, db)
	NewTransactionController(db).RegisterRoutes(router, db)
	NewCustomerReportController(db).RegisterRoutes(router, db)
	NewAccountReportController(db).RegisterRoutes(router, db)
	NewTransactionReportController(db).RegisterRoutes(router, db)
}