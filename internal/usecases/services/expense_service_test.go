package services

import (
	"errors"
	"testing"

	"github.com/idmaksim/expense-tracker-cli/internal/domain/models"
	"github.com/idmaksim/expense-tracker-cli/internal/domain/repositories/mocks"
)

func TestExpenseService_CreateExpense(t *testing.T) {
	repo := mocks.NewMockExpenseRepository()
	service := NewExpenseService(repo)

	expense := models.NewExpense("Test", 100.0)

	err := service.CreateExpense(expense)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	repo.SetError(errors.New("database error"))
	err = service.CreateExpense(expense)
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestExpenseService_FindAll(t *testing.T) {
	repo := mocks.NewMockExpenseRepository()
	service := NewExpenseService(repo)

	expense1 := models.NewExpense("Test1", 100.0)
	expense2 := models.NewExpense("Test2", 200.0)

	repo.Create(expense1)
	repo.Create(expense2)

	expenses, err := service.FindAll()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(expenses) != 2 {
		t.Errorf("Expected 2 expenses, got %d", len(expenses))
	}
}

func TestExpenseService_Summary(t *testing.T) {
	repo := mocks.NewMockExpenseRepository()
	service := NewExpenseService(repo)

	repo.Create(models.NewExpense("Test1", 100.0))
	repo.Create(models.NewExpense("Test2", 200.0))

	total, err := service.Summary()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if total != 300.0 {
		t.Errorf("Expected total 300.0, got %f", total)
	}
}
