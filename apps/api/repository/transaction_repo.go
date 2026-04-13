package repository

import (
	"api/models"
)

func GetAllTransaction() []models.Transaction {
	return []models.Transaction{
		{ID: 1, Amount: 2500, Merchant: "Swiggy"},
		{ID: 2, Amount: 250, Merchant: "Amazon"},
	}
}
