package mocks

import (
	"github.com/idmaksim/expense-tracker-cli/internal/domain/models"
)

type MockExpenseRepository struct {
	expenses []*models.Expense
	err      error
}

func NewMockExpenseRepository() *MockExpenseRepository {
	return &MockExpenseRepository{
		expenses: make([]*models.Expense, 0),
	}
}

func (m *MockExpenseRepository) SetError(err error) {
	m.err = err
}

func (m *MockExpenseRepository) Create(expense *models.Expense) error {
	if m.err != nil {
		return m.err
	}
	expense.ID = len(m.expenses) + 1
	m.expenses = append(m.expenses, expense)
	return nil
}

func (m *MockExpenseRepository) FindAll() ([]*models.Expense, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.expenses, nil
}

func (m *MockExpenseRepository) Delete(id int) error {
	if m.err != nil {
		return m.err
	}
	for i, e := range m.expenses {
		if e.ID == id {
			m.expenses = append(m.expenses[:i], m.expenses[i+1:]...)
			return nil
		}
	}
	return nil
}

func (m *MockExpenseRepository) Update(expense *models.Expense) error {
	if m.err != nil {
		return m.err
	}
	for i, e := range m.expenses {
		if e.ID == expense.ID {
			m.expenses[i] = expense
			return nil
		}
	}
	return nil
}

func (m *MockExpenseRepository) Summary() (float32, error) {
	if m.err != nil {
		return 0, m.err
	}
	var total float32
	for _, e := range m.expenses {
		total += e.Amount
	}
	return total, nil
}
