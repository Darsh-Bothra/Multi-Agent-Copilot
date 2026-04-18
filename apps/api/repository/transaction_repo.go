package repository

import (
	"api/internal/db"
	"api/models"
)

func GetAllTransaction() ([]models.Transaction, error) {
	row, err := db.DB.Query("SELECT id, amount, merchant FROM transactions")

	if err != nil {
		return nil, err;
	}

	// for memory cleanup
	defer row.Close()                                            

	var transactions []models.Transaction

	for row.Next() {
		var t models.Transaction

		// maps DB columns → struct fields	
		err := row.Scan(&t.ID, &t.Amount, &t.Amount)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, t)
	}
	return transactions, nil;
}


func InsertTransaction(amount float64, merchant string) error {
	// SQL Placeholder $1, $2
	// prevents SQL injection
	// safe query binding
	_, err := db.DB.Exec(
		"INSERT INTO transactions (amount, merchant) VALUES ($1, $2)",
		amount,
		merchant,
	)

	return err
}