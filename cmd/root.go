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
		tasks, err :=storage.LoadTasks()
		if err !=nil{
			fmt.Println("Error loading tasks:",err)
		}
		if len(tasks)==0{
			fmt.Println("No tasks found")
			return
		}
		for _,t:=range tasks{
			fmt.Printf(
				"[%d] %s | %s | created: %s | updated: %s\n",
				t.ID,
				t.Description,
				t.Status,
				t.CreatedAt.Format("2006-01-02 15:04:05"),
				t.UpdatedAt.Format("2006-01-02 15:04:05"),
			)
		}
	default:
		fmt.Println("Unknown command:", command)
	}
}