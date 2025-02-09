package handlers

import (
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
