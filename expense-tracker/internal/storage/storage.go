package storage

import (
	"encoding/json"
	"os"
	"expense-tracker/internal/expense"
)

const filepath = "data/data.json"

func LoadExpenses()([]expense.Expense, error){
	if _,err:=os.Stat(filepath);os.IsNotExist(err){
		return []expense.Expense{},nil
	}

	data,err:=os.ReadFile(filepath)
	if err!=nil{
		return nil, err
	}

	var expenses []expense.Expense
	if len(data) == 0 {
		return []expense.Expense{}, nil
	}
	err = json.Unmarshal(data, &expenses)
	if err != nil {
		return nil, err
	}
	return expenses, nil
}

func SaveTasks(tasks []expense.Expense) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	err = os.WriteFile(filepath,data,0644)
	if err!=nil{
		return err
	}
	return nil
}