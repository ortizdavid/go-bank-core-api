package entities

import (
	"github.com/go-playground/validator/v10"
)

//-------- Contact ---------
type UpdateCustomerContactRequest struct {
    CustomerId    int64 `json:"customer_id" validate:"required"`
    Email         string `json:"email,omitempty" validate:"omitempty,email,max=150"`
    Phone         string `json:"phone,omitempty" validate:"omitempty,e164,max=20"`
}

func (req UpdateCustomerContactRequest) Validate() error {
	v := validator.New()
	return v.Struct(req)
}

//-------- Status ---------
type ChangeCustomerStatusRequest struct {
    CustomerId  int64 `json:"customer_id" validate:"required"`
    NewStatus   CustomerStatus `json:"new_status" validate:"required"`
}

func (req ChangeCustomerStatusRequest) Validate() error {
	v := validator.New()
	return v.Struct(req)
}

//-------- Type ---------
type ChangeCustomerTypeRequest struct {
    CustomerId  int64 `json:"customer_id" validate:"required"`
    NewType     CustomerType `json:"new_type" validate:"required"`
}

func (req ChangeCustomerTypeRequest) Validate() error {
	v := validator.New()
	return v.Struct(req)
}

