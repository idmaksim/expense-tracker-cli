package models

import (
	"testing"
	"time"
)

func TestNewExpense(t *testing.T) {
	description := "Test expense"
	amount := float32(100.50)

	expense := NewExpense(description, amount)

	today := time.Now().Format("2006-01-02")

	if expense.Description != description {
		t.Errorf("Expected description %s, got %s", description, expense.Description)
	}

	if expense.Amount != amount {
		t.Errorf("Expected amount %f, got %f", amount, expense.Amount)
	}

	if expense.Date != today {
		t.Errorf("Expected date %s, got %s", today, expense.Date)
	}
}
