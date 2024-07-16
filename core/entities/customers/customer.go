package entities

import "time"

type Customer struct {
	CustomerId           int64 `gorm:"primaryKey;autoIncrement"`
	UniqueId             string `gorm:"column:unique_id"`
	CustomerType         int  `gorm:"column:type_id"`
	StatusId             int  `gorm:"column:status_id"`
	IdentificationNumber string `gorm:"column:identification_number"`
	Gender               string `gorm:"column:gender"`
	BirtDate             time.Time `gorm:"column:birth_date"`
	Email                string `gorm:"column:email"`
	Phone                string `gorm:"column:phone"`
	Address              string `gorm:"column:address"`
	CreatedAt            time.Time `gorm:"column:created_at"`
	UpdatedAt            time.Time `gorm:"column:updated_at"`
}

func (Customer) TableName() string {
	return "customers"
}
