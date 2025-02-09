package storage

import (
	"encoding/json"
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

func (s *JSONStorage) Create(expense *models.Expense) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	expenses, err := s.readExpenses()
	if err != nil {
		return err
	}

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
