package dto

type CreateTransactionRequest struct {
	// amount with the validaton, like required and > 0
	Amount float64 `json:"amount" validate:"required,gt=0"`
	Merchant string `json:"merchant" validata:"required,min=2"`
}	