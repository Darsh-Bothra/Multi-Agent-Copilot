package service

import (
	"api/repository"
	"fmt"
	"math"
	"sort"
)

// what service will the group have?-> create_grp, calc_balances, settlements

type Settlement struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

// 1. Create group and add memebers
func CreateGroup(name string, createdBy string, members []string) (string, error) {

	// create a group first
	grpId, err := repository.CreateGroup(name, createdBy)

	if err != nil {
		return "", err
	}

	// add members in the group
	err = repository.AddMembers(grpId, members)
	if err != nil {
		return "", err
	}

	return grpId, nil
}


// 2. Adding the expense
func AddExpense(grpId string, paidBy string, amount float64, description string, split map[string]float64) (string, error) {

	// validate the amount
	var total float64
	for _, val := range split {
		total += val
	}

	if math.Abs(total - amount) > 0.0001 {
		return "", fmt.Errorf("split total does not match expense amount")
	}

	// first we create the expense
	expId, err := repository.CreateExpense(grpId, paidBy, amount, description)

	if err != nil {
		return "", err
	}

	// create the splits
	err = repository.CreateSplits(expId, split)

	if err != nil {
		return "", err
	}

	return expId, nil

}


// 3. Calculate balances for each member
func CalculateBalances(grpId string) (map[string]float64, error) {
	// first get the expenses using the group id
	expenses, err := repository.GetExpensesByGroup(grpId)

	if err != nil {
		return nil, err
	}

	// now using expense id for each expense in the expenses we'll get the split
	balances := make(map[string]float64)

	for _, expense := range expenses {
		// calulate the expediture of each member
		balances[expense.PaidBy] += expense.Amount

		// get the splits
		splits, err := repository.GetSplitsByExpense(expense.ID)

		if err != nil {
			return nil, err
		}

		// simple logic balance = paid - owed, if balance < 0: red else green
		for _, split := range splits {
			balances[split.UserId] -= split.AmountOwed
		}
	}
	return balances, nil
}


// 4. Settlements 
func ClarifyDebts(balances map[string]float64) []Settlement {
	var creditors []struct {
		user string
		amount float64
	}
	var debtors []struct {
		user string
		amount float64
	}

	// seperate debtors and creditors
	for  user, balance := range balances {
		if balance < 0 {
			debtors = append(debtors, struct{user string; amount float64}{user, -balance});
		} else if balance > 0 {
			creditors = append(creditors, struct{user string; amount float64}{user, balance})
		}
	}

	// sort creditors and debtors for NlogN compute
	sort.Slice(creditors, func(i, j int) bool {return creditors[i].amount > creditors[j].amount});
	sort.Slice(debtors, func(i, j int) bool {return debtors[i].amount > debtors[j].amount});

	var settlements []Settlement
	i, j := 0, 0;

	for i < len(debtors) && j < len(creditors) {
		debt, credit := debtors[i], creditors[j];

		minAmt := math.Min(debt.amount, credit.amount)

		settlements = append(settlements, Settlement{
			From: debt.user,
			To: credit.user,
			Amount: minAmt,
		})

		debt.amount -= minAmt;
		credit.amount -= minAmt;

		debtors[i].amount = debt.amount
		creditors[j].amount = credit.amount

		const epsilon = 0.0001
		if debtors[i].amount < epsilon {
			i++
		}
		if creditors[j].amount < epsilon {
			j++
		}
	}
	return settlements
}