package entities

import "time"

type AccountData struct {
	AccountId            int       `json:"account_id"`
	UniqueId             string    `json:"unique_id"`
	AccountNumber        string    `json:"account_number"`
	Iban                 string    `json:"iban"`
	Balance              float64   `json:"balance"`
	Currency             string    `json:"currency"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
	CustomerId           int       `json:"customer_id"`
	CustomerName         string    `json:"customer_name"`
	IdentificationNumber string    `json:"identification_number"`
	TypeId               int       `json:"type_id"`
	TypeName             string    `json:"type_name"`
	StatusId             int       `json:"status_id"`
	StatusName           string    `json:"status_name"`
}
