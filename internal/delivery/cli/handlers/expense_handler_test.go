package handlers

import (
	"testing"

	"github.com/idmaksim/expense-tracker-cli/internal/domain/repositories/mocks"
	"github.com/idmaksim/expense-tracker-cli/internal/usecases/services"
)

func TestExpenseHandler_CreateExpense(t *testing.T) {
	repo := mocks.NewMockExpenseRepository()
	service := services.NewExpenseService(repo)
	handler := NewExpenseHandler(service)

	err := handler.CreateExpense("Test expense", 100.50)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestExpenseHandler_ListExpenses(t *testing.T) {
	repo := mocks.NewMockExpenseRepository()
	service := services.NewExpenseService(repo)
	handler := NewExpenseHandler(service)

	handler.CreateExpense("Test1", 100.50)
	handler.CreateExpense("Test2", 200.50)

	err := handler.ListExpenses()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func TestExpenseHandler_UpdateExpense(t *testing.T) {
	repo := mocks.NewMockExpenseRepository()
	service := services.NewExpenseService(repo)
	handler := NewExpenseHandler(service)

	handler.CreateExpense("Test1", 100.50)
	err := handler.UpdateExpense(1, "Updated Test", 150.75)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
