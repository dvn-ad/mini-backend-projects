package cmd

import (
	"fmt"
	"os"
	"task-cli/internal/service"
)

func Run() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: task-cli <command> [arguments]")
		return
	}

	command := args[1]
	cmdArgs := args[2:]

	switch command {
	case "add":
		if len(cmdArgs) < 1 {
			fmt.Println("Error: missing task description")
			return
		}

		err:=task.AddTask(cmdArgs[0])
		if err != nil {
			fmt.Println(err)
			return
		}

	case "list":
		if len(cmdArgs) < 1 {
			tasks, err:=task.ListTasks("")
			if err != nil{
				fmt.Println("Error: failed to load tasks")
				return
			}
			for _, t := range tasks{
				fmt.Printf(
					"[%d] %s | %s | created: %s | updated: %s\n",
					t.ID,
					t.Description,
					t.Status,
					t.CreatedAt.Format("2006-01-02 15:04:05"),
					t.UpdatedAt.Format("2006-01-02 15:04:05"),
				)
			}
		}
		fmt.Println("Tasks:")		

	default:
		fmt.Println("Unknown command:", command)
	}
}