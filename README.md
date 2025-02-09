# 💰 Expense Tracker CLI

A command-line expense tracking application built with Go using clean architecture principles.

## ✨ Features

- 📝 Add new expenses
- 🔄 Update existing expenses
- 🗑️ Delete expenses
- 📊 View all expenses
- 💰 Get total expenses summary

## 🚀 Getting Started

### Prerequisites

- Go 1.23 or higher

### Installation

```
go install github.com/idmaksim/expense-tracker-cli/cmd/expense-cli
```

## 💡 Usage

### Managing Expenses

```
# Add a new expense
expense-cli add "Groceries" 100.50

# Update existing expense
expense-cli update 1 "Groceries and household items" 150.75

# Delete expense
expense-cli delete 1

```

### Viewing Information

```
# List all expenses
expense-cli list

# Show total expenses
expense-cli summary
```

## 🏗️ Project Structure

```
.

├── cmd/
│   └── expense-cli/        # Application entry point
├── internal/
│   ├── domain/            # Business logic and entities
│   │   ├── models/        # Domain models
│   │   └── repositories/  # Repository interfaces
│   ├── usecases/         # Application use cases
│   │   └── services/     # Service layers

│   ├── infrastructure/   # External implementations
│   │   └── storage/      # JSON storage
│   └── delivery/         # CLI interface
│       └── cli/
│           ├── commands/ # CLI commands
│           └── handlers/ # Command handlers
```

## 📊 Expense Properties

Each expense record includes:

- 🔑 Unique ID
- 📝 Description
- 💵 Amount
- 📅 Creation date

## 🛠️ Technical Details

- **Clean Architecture**: Clear separation of application layers
- **Thread Safety**: Mutex-protected JSON storage
- **Error Handling**: Proper error handling and user-friendly messages
- **CLI Framework**: Built using Cobra
- **Data Storage**: Local JSON storage in user's home directory
