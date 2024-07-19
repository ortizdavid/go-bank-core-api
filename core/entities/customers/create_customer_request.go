package entities

import (
	"github.com/go-playground/validator/v10"
)

type CreateCustomerRequest struct {
    CustomerType         CustomerType `json:"customer_type" validate:"required"`
    CustomerName         string    `json:"customer_name" validate:"required,min=2,max=150"`
    Gender               string    `json:"gender" validate:"required,oneof=Male Female Other"`
    BirthDate            string    `json:"birth_date" validate:"required"`
    IdentificationNumber string    `json:"identification_number" validate:"required,min=6,max=30"`
    Email                string    `json:"email,omitempty" validate:"omitempty,email,max=150"`
    Phone                string    `json:"phone,omitempty" validate:"omitempty,e164,max=20"`
    Address              string    `json:"address,omitempty" validate:"omitempty,max=200"`
}

func (req CreateCustomerRequest) Validate() error {
	v := validator.New()
	return v.Struct(req)
}