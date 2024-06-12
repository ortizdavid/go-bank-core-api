package entities

type DepositRequest struct {
	AccountNumber string  `json:"account_number" validate:"required,max=18"`
	Amount        float64 `json:"amount" validate:"required"`
	Currency      string  `json:"currency" validate:"required,len=3"`
}

type TransferIbanRequest struct {
	SourceIban      string  `json:"source_iban" validate:"required,max=31"`
	DestinationIban string  `json:"destination_iban" validate:"required,max=31"`
	Amount          float64 `json:"amount" validate:"required"`
	Currency        string  `json:"currency" validate:"required,len=3"`
}

type TransferNumberRequest struct {
	SourceNumber      string  `json:"source_number" validate:"required,max=18"`
	DestinationNumber string  `json:"destination_number" validate:"required,max=18"`
	Amount            float64 `json:"amount" validate:"required"`
	Currency          string  `json:"currency" validate:"required,len=3"`
}

type WithdrawRequest struct {
	AccountNumber string  `json:"account_number" validate:"required,max=18"`
	Amount        float64 `json:"amount" validate:"required"`
	Currency      string  `json:"currency" validate:"required,len=3"`
}
