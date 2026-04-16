package service

import (
	"api/models"
	"api/repository"
)


func GetTransactions() ([]models.Transaction, error) {
	return repository.GetAllTransaction()
}

func CreateTransaction(amount float64, merchant string) error {
	// we have to save the transaction in the repository (the database)
	return repository.InsertTransaction(amount, merchant)
}