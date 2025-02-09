package repositories

import "github.com/idmaksim/expense-tracker-cli/internal/domain/models"

type ExpenseRepository interface {
	Create(expense *models.Expense) error
	FindAll() ([]*models.Expense, error)
	Delete(id int) error
}
