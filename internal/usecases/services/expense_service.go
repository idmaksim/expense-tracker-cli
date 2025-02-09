package services

import (
	"errors"
	"sort"

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
	err := s.repo.Create(expense)
	if err != nil {
		return err
	}

	return nil
}

func (s *ExpenseService) Summary() (float32, error) {
	total, err := s.repo.Summary()
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (s *ExpenseService) DeleteExpense(id int) error {
	return s.repo.Delete(id)
}

func (s *ExpenseService) FindAll() ([]*models.Expense, error) {
	expenses, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	if len(expenses) == 0 {
		return nil, errors.New("no expenses found")
	}

	sort.Slice(expenses, func(i, j int) bool {
		return expenses[i].Date < expenses[j].Date
	})

	return expenses, nil
}
