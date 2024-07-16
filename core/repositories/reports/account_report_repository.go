package repositories

import "gorm.io/gorm"

type AccountReportRepository struct {
	db *gorm.DB
}

func NewAccountReportRepository(db *gorm.DB) *AccountReportRepository {
	return &AccountReportRepository{
		db: db,
	}
}