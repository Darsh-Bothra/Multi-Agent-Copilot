package service

import (
	"api/models"
	"api/repository"
)


func GetTransactions() []models.Transaction {
	return repository.GetAllTransaction()
}