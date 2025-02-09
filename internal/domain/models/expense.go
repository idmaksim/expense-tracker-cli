package models

import "time"

type Expense struct {
	ID          int     `json:"id"`
	Date        string  `json:"date"`
	Description string  `json:"description"`
	Amount      float32 `json:"amount"`
}

func NewExpense(description string, amount float32) *Expense {
	today := time.Now()
	return &Expense{
		Description: description,
		Amount:      amount,
		Date:        today.Format("2006-01-02"),
	}
}
