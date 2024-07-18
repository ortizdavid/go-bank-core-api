package entities

import "time"

type Customer struct {
	CustomerId           int64 `gorm:"primaryKey;autoIncrement"`
	CustomerType         CustomerType  `gorm:"column:customer_type"`
	CustomerStatus       CustomerStatus  `gorm:"column:customer_status"`
	CustomerName         string `gorm:"column:customer_name"`
	IdentificationNumber string `gorm:"column:identification_number"`
	Gender               string `gorm:"column:gender"`
	BirtDate             time.Time `gorm:"column:birth_date"`
	Email                string `gorm:"column:email"`
	Phone                string `gorm:"column:phone"`
	Address              string `gorm:"column:address"`
	UniqueId             string `gorm:"column:unique_id"`
	CreatedAt            time.Time `gorm:"column:created_at"`
	UpdatedAt            time.Time `gorm:"column:updated_at"`
}

func (Customer) TableName() string {
	return "customers"
}
