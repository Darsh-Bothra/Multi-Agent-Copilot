package models

type Split struct {
	ID         string  `json:"id"`
	ExpenseId  string  `json:"expense_id"`
	UserId     string  `json:"user_id"`
	AmountOwed float64 `json:"amount_owed"`
}
