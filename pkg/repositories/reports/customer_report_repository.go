package repositories

import "gorm.io/gorm"

type CustomerReportRepository struct {
	db *gorm.DB
}

func NewCustomerReportRepository(db *gorm.DB) *CustomerReportRepository {
	return &CustomerReportRepository{
		db: db,
	}
}