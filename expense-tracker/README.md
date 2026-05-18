# Expense Tracker CLI

A simple command-line interface to track and manage your personal expenses.

This project is a solution to the [Expense Tracker](https://roadmap.sh/projects/expense-tracker) challenge from [roadmap.sh](https://roadmap.sh).

## Features

- Add expenses with a description and amount.
- List all recorded expenses.
- View a summary of all expenses.
- View a summary of expenses for a specific month.
- Delete an existing expense by its ID.
- Efficient data handling with incremental updates for summaries.

## Installation

Ensure you have [Go](https://go.dev/doc/install) installed on your system (version 1.26 or higher).

1. Clone the repository:
   ```bash
   git clone <your-repository-url> # Replace with your actual repository URL
   cd expense-tracker
   ```

2. Build the application:
   ```bash
   go build -o expense-tracker
   ```

## Usage

You can run the application using `./expense-tracker` (after building) or `go run main.go`.

### Commands

| Action | Command Example |
| :--- | :--- |
| **Add a new expense** | `expense-tracker add --description "Lunch" --amount 20` |
| **List all expenses** | `expense-tracker list` |
| **View total summary** | `expense-tracker summary` |
| **View monthly summary** | `expense-tracker summary --month 5` |
| **Delete an expense** | `expense-tracker delete --id 1` |
| **Update an expense** | `expense-tracker update --id 1 --description "Updated Lunch" --amount 25` *(TODO: Implementation pending)* |

## Data Storage

Expenses are stored in `data/data.json` and aggregated summaries in `data/summary.json`. The application will create these files automatically if they don't exist.