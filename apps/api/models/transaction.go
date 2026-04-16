package models

type Transaction struct {
	ID       string     `json:"id"`
	Amount   float64 `json:"amount"`
	Merchant string  `json:"merchant"`
}
