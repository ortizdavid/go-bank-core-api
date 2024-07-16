package entities

type ChangeAccountStatusRequest struct {
	AccountNumber string       `json:"account_number" validate:"required"`
	AccountStatus AccountStatus `json:"account_status" validate:"required"`
}

type ChangeAccountTypeRequest struct {
	AccountNumber string      `json:"account_number" validate:"required"`
	AccountType   AccountType `json:"account_type" validate:"required"`
}

type CreateAccountRequest struct {
	CustomerId  int         `json:"customer_id" validate:"required"`
	AccountType AccountType `json:"account_type" validate:"required"`
	Currency    string      `json:"currency" validate:"required"`
}

type CreateAccountWithCustomerRequest struct {
	CustomerName         string      `json:"customer_name" validate:"required,max=150"`
	IdentificationNumber string      `json:"identification_number" validate:"required,max=30"`
	Email                string      `json:"email" validate:"required,email,min=10,max=150"`
	Phone                string      `json:"phone" validate:"required,max=20"`
	Address              string      `json:"address" validate:"required,max=200"`
	AccountType          AccountType `json:"account_type_id" validate:"required"`
	Currency             string      `json:"currency" validate:"required,max=3"`
}
