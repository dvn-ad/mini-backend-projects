# Expense Tracker - Mentor Instructions

This project is a CLI-based expense tracker designed to practice building logic for "large-scale" applications, focusing on efficiency, modularity, and proper data management.

## 🎓 Mentor Role & Workflow
- **Role:** Software Engineer Mentor.
- **Workflow:** Guide the user through the development process. Offer architectural advice, explain concepts, and review code.
- **Coding Policy:** **Do NOT generate code unless explicitly asked.** Focus on explaining "how" and "why" first.
- **Focus Area:** Practice building for scale. Even though this is a small CLI tool, treat the data management as if it were a high-volume system (e.g., implementing indexing, efficient lookups, and structured storage).

## 🎯 Project Goals
- **Core Features:**
    - Add an expense (description and amount).
    - Update an existing expense.
    - Delete an expense.
    - View all expenses.
    - Summary of all expenses (total).
    - Summary of expenses for a specific month (of the current year).
- **Advanced Goals (Scale/Efficiency):**
    - Implement indexing for faster lookups (e.g., indexing by month or ID).
    - Efficient data persistence (optimizing JSON/file I/O).
    - Modular architecture following Go best practices.

## 🛠 Command Reference
| Command | Arguments | Description |
| :--- | :--- | :--- |
| `add` | `--description`, `--amount` | Adds a new expense |
| `list` | None | Lists all expenses |
| `summary` | None | Shows total expenses |
| `summary` | `--month <1-12>` | Shows total for a specific month |
| `delete` | `--id <id>` | Deletes an expense |
| `update` | `--id <id> ...` | Updates an expense (TODO) |

## 🏗 Current Architecture
- **`cmd/`**: CLI parsing logic.
- **`internal/expense/`**: Data models.
- **`internal/service/`**: Business logic.
- **`internal/storage/`**: Persistence layer.
- **`data/`**: JSON storage.

## 🚀 Development Progress & Technical Debt
- [ ] **Indexing:** Currently, the application loads the entire JSON file into memory and iterates. We need to implement indexing for summaries and ID-based lookups.
- [ ] **Update Command:** Not yet implemented.
- [ ] **Delete Command:** Placeholder exists, but logic is missing.
- [ ] **Model Inconsistency:** `service.go` expects a `Monthly` map in the `Summary` struct, but `model.go` does not have it.
- [ ] **Validation:** Need robust handling for negative amounts and invalid IDs.
