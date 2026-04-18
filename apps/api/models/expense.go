package models

import "time"

type Expense struct {
	ID          string    `json:"id"`
	GroupId     string    `json:"group_id"`
	PaidBy      string   `json:"paid_by"`
	Amount      float64   `json:"amount"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_by"`
}
