package handlers

import (
	"fmt"

	"github.com/idmaksim/expense-tracker-cli/internal/domain/models"
	"github.com/idmaksim/expense-tracker-cli/internal/usecases/services"
)

type ExpenseHandler struct {
	service *services.ExpenseService
}

func NewExpenseHandler(service *services.ExpenseService) *ExpenseHandler {
	return &ExpenseHandler{service: service}
}

func (handler *ExpenseHandler) CreateExpense(description string, amount float32) error {
	expense := models.NewExpense(description, amount)
	return handler.service.CreateExpense(expense)
}

func (handler *ExpenseHandler) ListExpenses() error {
	expenses, err := handler.service.FindAll()
	if err != nil {
		return fmt.Errorf("failed to list expenses: %w", err)
	}

	fmt.Println("ID\tDate\t\tDescription\tAmount")
	for _, expense := range expenses {
		fmt.Printf("%d\t%s\t%s\t\t%.2f\n", expense.ID, expense.Date, expense.Description, expense.Amount)
	}
	return nil
}

func (handler *ExpenseHandler) DeleteExpense(id int) error {
	return handler.service.DeleteExpense(id)
}
