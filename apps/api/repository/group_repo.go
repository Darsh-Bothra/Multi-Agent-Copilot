package repository

import (
	"api/internal/db"
	"api/models"
)


// Inserts a new group in the record in the DB
func CreateGroup(name string, createdBy string) (string, error) {
	query := `
		INSERT INTO groups (name, createdBy)
		VALUES ($1, $2)
		RETURNING id
	`	
	var groupId string;
	// Execute the query
	err := db.DB.QueryRow(query, name, createdBy).Scan(&groupId)

	if err != nil {
		return "", err;
	}
	return groupId, nil;
}


// Adding the members in the group
func AddMembers(groupId string, userIds []string) error {
	query := `
		INSERT INTO group_members (group_id, user_id)
		VALUES ($1, $2)
	`

	for _, userId := range userIds {
		_, err := db.DB.Exec(query, groupId, userId)

		if err != nil {
			return err;
		}
	} 
	return nil;
}


// Creating a expense in a group
func CreateExpense(groupId string, paidBy string, amount float64, description string) (string, error) {
	query := `
		INSERT INTO expenses (group_id, paidBy, amount, description)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	var expenseId string;
	err := db.DB.QueryRow(query, groupId, paidBy, amount, description).Scan(&expenseId);
	
	if err != nil {
		return "", err;
	}
	return expenseId, nil;
}


// Creating splits
func CreateSplits(expenseId string, splits map[string]float64) error {
	query := `
		INSERT INTO splits (expense_id, user_id, amount_owed)
		VALUES ($1, $2, $3)
	`

	for userId, amount := range splits {
		_, err := db.DB.Exec(query, expenseId, userId, amount);

		if err != nil {
			return err;
		}
	}
	return nil;
}


// Get the expenses in a particular group 	
func GetExpensesByGroup(groupId string) ([]models.Expense, error) {
	query := `
		SELECT id, group_id, paid_by, amount, description, created_at
		FROM expenses
		WHERE group_id = $1
	`

	row, err := db.DB.Query(query, groupId);

	if err != nil {
		return nil, err
	}

	// releases connection back to pool
	defer row.Close();

	var expenses []models.Expense;

	for row.Next() {
		var expense models.Expense

		err := row.Scan(
			&expense.ID,
			&expense.GroupId,
			&expense.Amount,
			&expense.Description,
			&expense.CreatedAt,
		)

		if err != nil {
			return nil, err;
		}

		expenses = append(expenses, expense)
	}
	return expenses, nil;
}


// Get all the splits for a particular expense
func GetSplitsByExpense(expenseId string) ([]models.Split, error) {
	query := `
		SELECT id, expense_id, user_id, amount_owed
		FROM splits
		WHERE expense_id = $1
	`

	row, err := db.DB.Query(query, expenseId);

	if err != nil {
		return nil, err
	}

	// releases connection back to pool
	defer row.Close();

	var splits []models.Split;

	for row.Next() {
		var split models.Split

		err := row.Scan(
			&split.ID,
			&split.ExpenseId,
			&split.UserId,
			&split.AmountOwed,
		)

		if err != nil {
			return nil, err;
		}

		splits = append(splits, split)
	}
	return splits, nil;
}