package services

import (
	"github.com/idmaksim/expense-tracker-cli/internal/domain/models"
	"github.com/idmaksim/expense-tracker-cli/internal/domain/repositories"
)

type ExpenseService struct {
	repo repositories.ExpenseRepository
}

func NewExpenseService(repo repositories.ExpenseRepository) *ExpenseService {
	return &ExpenseService{repo: repo}
}

func (s *ExpenseService) CreateExpense(expense *models.Expense) error {
	return s.repo.Create(expense)
}
