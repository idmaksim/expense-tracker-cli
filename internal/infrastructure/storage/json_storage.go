package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/idmaksim/expense-tracker-cli/internal/domain/models"
)

type JSONStorage struct {
	filePath string
	mu       sync.RWMutex
}

func NewJSONStorage(filePath string) *JSONStorage {
	return &JSONStorage{filePath: filePath}
}

func (s *JSONStorage) Init() error {
	if _, err := os.Stat(s.filePath); os.IsNotExist(err) {
		return s.writeExpenses([]*models.Expense{})
	}
	return nil
}

func (s *JSONStorage) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	expenses, err := s.readExpenses()
	if err != nil {
		return err
	}

	for i, expense := range expenses {
		if expense.ID == id {
			expenses = append(expenses[:i], expenses[i+1:]...)
			return s.writeExpenses(expenses)
		}
	}

	return fmt.Errorf("expense with id %d not found", id)
}

func (s *JSONStorage) FindAll() ([]*models.Expense, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.readExpenses()
}

func (s *JSONStorage) Summary() (float32, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	expenses, err := s.readExpenses()
	if err != nil {
		return 0, err
	}

	total := float32(0.0)
	for _, expense := range expenses {
		total += expense.Amount
	}

	return total, nil
}

func (s *JSONStorage) Update(expense *models.Expense) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	expenses, err := s.readExpenses()
	if err != nil {
		return err
	}

	updated := false

	for i, e := range expenses {
		if e.ID == expense.ID {
			expenses[i] = expense
			updated = true
		}
	}

	if !updated {
		return fmt.Errorf("expense with id %d not found", expense.ID)
	}

	return s.writeExpenses(expenses)
}

func (s *JSONStorage) Create(expense *models.Expense) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	expenses, err := s.readExpenses()
	if err != nil {
		return err
	}

	maxId := 0
	for _, expense := range expenses {
		if expense.ID > maxId {
			maxId = expense.ID
		}
	}

	expense.ID = maxId + 1

	expenses = append(expenses, expense)

	return s.writeExpenses(expenses)
}

func (s *JSONStorage) writeExpenses(tasks []*models.Expense) error {
	data, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	return os.WriteFile(s.filePath, data, 0644)
}

func (s *JSONStorage) readExpenses() ([]*models.Expense, error) {
	data, err := os.ReadFile(s.filePath)
	if err != nil {
		return nil, err
	}

	var expenses []*models.Expense
	if err := json.Unmarshal(data, &expenses); err != nil {
		return nil, err
	}

	return expenses, nil
}
