package repositories

import "gorm.io/gorm"

type TransactionReportRepository struct {
	db *gorm.DB
}

func NewTransactionReportRepository(db *gorm.DB) *TransactionReportRepository {
	return &TransactionReportRepository{
		db: db,
	}
}