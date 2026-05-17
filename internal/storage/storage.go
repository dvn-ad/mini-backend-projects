package storage

import (
	"encoding/json"
	"os"
	"task-cli/internal/task"
)

const filepath = "data/tasks.json"

func LoadTasks([]task.Task, error) {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return []task.Task{}, nil
	}
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var tasks []task.Task
	if len(data) == 0 {
		return []task.Task{}, nil
	}
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
func SaveTasks(tasks []task.Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nli {
		return err
	}
	err = os.WriteFile(filepath,data,0644)
	if err!=nil{
		return err
	}
	return nil
}
