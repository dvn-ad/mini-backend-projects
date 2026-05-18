# Task Tracker CLI

A simple command-line interface to track your tasks and manage your to-do list.

This project is a solution to the [Task Tracker](https://roadmap.sh/projects/task-tracker) challenge from [roadmap.sh](https://roadmap.sh).

## Features

- Add, Update, and Delete tasks
- Mark tasks as in-progress or done
- List all tasks
- List tasks by status (todo, in-progress, done)

## Installation

Ensure you have [Go](https://go.dev/doc/install) installed on your system.

1. Clone the repository:
   ```bash
   git clone https://github.com/dvn-ad/mini-backend-projects.git
   cd mini-backend-projects/task-cli
   ```

2. Build the application:
   ```bash
   go build -o task-cli
   ```

## Usage

You can run the application using `./task-cli` (after building) or `go run main.go`.

### Commands

| Action | Command |
| :--- | :--- |
| **Add a new task** | `task-cli add "Buy groceries"` |
| **Update a task** | `task-cli update 1 "Buy groceries and cook dinner"` |
| **Delete a task** | `task-cli delete 1` |
| **Mark in progress** | `task-cli mark-in-progress 1` |
| **Mark as done** | `task-cli mark-done 1` |
| **List all tasks** | `task-cli list` |
| **List by status** | `task-cli list todo`, `task-cli list in-progress`, `task-cli list done` |

## Data Storage

Tasks are stored in a JSON file located at `data/tasks.json`. The application will create this file automatically if it doesn't exist.