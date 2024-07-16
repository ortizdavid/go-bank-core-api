package entities

import "time"

type Transaction struct {
	TransactionId     int64`gorm:"primaryKey;autoIncrement"`
	SourceId          int64`gorm:"column:source_id"`
	DestinationId     int64`gorm:"column:destination_id"`
	TransactionStatus TransactionStatus`gorm:"column:transaction_status_id"`
	TransactionType   TransactionType`gorm:"column:transaction_type_id"`
	Code              string`gorm:"column:code"`
	Amount            float64`gorm:"column:amount"`
	Currency          string`gorm:"column:currency"`
	BalanceBefore     float64`gorm:"column:balance_before"`
	BalanceAfter      float64`gorm:"column:balance_after"`
	Description       string`gorm:"column:description"`
	TransactionDate   time.Time`gorm:"column:transaction_date"`
	UniqueId          string`gorm:"column:unique_id"`
	CreatedAt         time.Time`gorm:"column:created_at"`
	UpdatedAt         time.Time`gorm:"column:updated_at"`
}

func (*Transaction) TableName() string {
	return "transactions"
}
