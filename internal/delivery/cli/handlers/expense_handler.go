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
	err := handler.service.CreateExpense(expense)
	if err != nil {
		return fmt.Errorf("failed to create expense: %w", err)
	}

	fmt.Println("Expense created successfully")
	return nil
}

func (handler *ExpenseHandler) ListExpenses() error {
	expenses, err := handler.service.FindAll()
	if err != nil {
		return fmt.Errorf("failed to list expenses: %w", err)
	}

	fmt.Println("ID\tDate\t\tAmount\t\tDescription")
	for _, expense := range expenses {
		fmt.Printf("%d\t%s\t%.2f\t\t%s\n", expense.ID, expense.Date, expense.Amount, expense.Description)
	}
	return nil
}

func (handler *ExpenseHandler) DeleteExpense(id int) error {
	return handler.service.DeleteExpense(id)
}

func (handler *ExpenseHandler) Summary() error {
	total, err := handler.service.Summary()
	if err != nil {
		return fmt.Errorf("failed to get summary: %w", err)
	}

	fmt.Printf("Total expenses: %.2f\n", total)
	return nil
}
