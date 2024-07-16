package entities

import "time"

type CustomerData struct {
	CustomerId           int64  `json:"customer_id"`
	UniqueId             string `json:"unique_id"`
	CustomerName         string `json:"customer_name"`
	IdentificationNumber string `json:"identification_number"`
	Gender               string `json:"gender"`
	BirtDate             time.Time `json:"birth_date"`
	Email				 string `json:"email"`
	Phone				 string `json:"phone"` 
	Address				 string `json:"address"`
	CreatedAt			 time.Time `json:"created_at"`
	UpdatedAt			 time.Time `json:"updated_at"`
	TypeId				 int `json:"type_id"`
	TypeName			 string `json:"type_name"`
	StatusId			 int `json:"status_id"`
	StatusName			 string `json:"status_name"`
}