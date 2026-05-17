package cmd

import (
	"fmt"
	"os"
	"task-cli/internal/storage"
	"task-cli/internal/task"
	"time"
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

		description:=cmdArgs[0]
		
		tasks, err := storage.LoadTasks()
		if err != nil{
			fmt.Println("Error loading tasks:", err)
			return
		}
		newTask:=task.Task{
			ID: len(tasks)+1,
			Description: description,
			Status: "todo",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		tasks = append(tasks,newTask)
		err = storage.SaveTasks(tasks)
		if err!=nil{
			fmt.Println("Error savin tasks:",err)
			return
		}
		fmt.Println("Task added successfully (ID:", newTask.ID, ")")
		
	case "list":
		fmt.Println("Tasks:")

	default:
		fmt.Println("Unknown command:", command)
	}
}