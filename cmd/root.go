package cmd

import (
	"fmt"
	"os"
	"task-cli/internal/service"
	"strconv"
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
			// LIST ALL TASKS
			if len(cmdArgs) < 1 {
				tasks, err:=task.ListTasks("")
				if err != nil{
					fmt.Println("Error: failed to load tasks")
					return
				}
				fmt.Println("Tasks:")		
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
			}else if cmdArgs[0]=="done"{
				tasks, err:=task.ListTasks("done")
				if err != nil{
					fmt.Println("Error: failed to load tasks")
					return
				}
				fmt.Println("Tasks:")		
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
			}else if cmdArgs[0]=="in-progress"{
				tasks, err:=task.ListTasks("in-progress")
				if err != nil{
					fmt.Println("Error: failed to load tasks")
					return
				}
				fmt.Println("Tasks:")		
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
			}else if cmdArgs[0]=="todo"{
				tasks, err:=task.ListTasks("todo")
				if err != nil{
					fmt.Println("Error: failed to load tasks")
					return
				}
				fmt.Println("Tasks:")		
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

		
		case "mark-in-progress":
			idStr :=cmdArgs[0]
			id, err:=strconv.Atoi(idStr)
			if err!=nil{
				fmt.Println("Invalid id")
				return
			}
			task.UpdateTask(id, "in-progress")

		case "mark-done":
			idStr :=cmdArgs[0]
			id, err:=strconv.Atoi(idStr)
			if err!=nil{
				fmt.Println("Invalid id")
				return
			}
			task.UpdateTask(id, "done")

		default:
			fmt.Println("Unknown command:", command)
	}
}