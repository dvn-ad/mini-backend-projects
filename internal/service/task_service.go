package task

import (
	"fmt"
	"task-cli/internal/storage"
	"task-cli/internal/task"
	"time"
)

func AddTask(description string) error {
		
		tasks, err := storage.LoadTasks()
		if err != nil{
			fmt.Println("Error loading tasks:", err)
			return err
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
			fmt.Println("Error saving tasks:",err)
			return err
		}
		fmt.Println("Task added successfully (ID:", newTask.ID, ")")
		return err
}

func ListTasks(status string) ([]task.Task, error){
	tasks, err:=storage.LoadTasks()
	if err != nil{
		return nil, err
	}

	
	if status == ""{
		return tasks,nil
	}else if status != "" {
		
		var filtered []task.Task
		
		for _, t := range tasks{
			if t.Status == status{
				filtered = append(filtered, t)
			}
		}
		return filtered, err
	}
	return nil, err
}