package entities

import "time"

type TransactionData struct {
	TransactionId        		int64 `json:"transaction_id"`
	UniqueId             		string `json:"unique_id"`
	Code                 		string `json:"code"`
	Amount               		float64 `json:"amount"`
	BalanceBefore        		float64 `json:"balance_before"`
	BalanceAfter         		float64 `json:"balance_after"`
	Currency             		string `json:"currency"`
	Description          		string `json:"description"`
	TransactionDate      		time.Time `json:"transaction_date"`
	CreatedAt            		time.Time `json:"created_at"`
	UpdatedAt            		time.Time `json:"updated_at"`
	SourceAccountId      		int64 `json:"source_account_id"`
	SourceAccountNumber       	string `json:"source_number"`
	SourceIban                 	string `json:"source_iban"`
	DestinationAccountId      	int64 `json:"destination_account_id"`
	DestinationAccountNumber    string `json:"destination_number"`
	DestinationIban             string `json:"destionation_iban"`
	CustomerName         		string `json:"customer_name"`
	IdentificationNumber	 	string `json:"identification_number"`
	TypeId               		int `json:"type_id"`
	TypeName             		string `json:"type_name"`
	StatusId             		int `json:"status_id"`
	StatusName           		string `json:"status_name"`
}