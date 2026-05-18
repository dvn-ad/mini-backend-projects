package storage

import (
	"encoding/json"
	"expense-tracker/internal/expense"
	"os"
)

const datapath = "data/data.json"

func LoadExpenses()([]expense.Expense, error){
	if _,err:=os.Stat(datapath);os.IsNotExist(err){
		return []expense.Expense{},nil
	}

	data,err:=os.ReadFile(datapath)
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

func SaveExpenses(expenses []expense.Expense) error {
	data, err := json.MarshalIndent(expenses, "", " ")
	if err != nil {
		return err
	}
	err = os.WriteFile(datapath,data,0644)
	if err!=nil{
		return err
	}
	return nil
}


const summaryPath = "data/summary.json"
func UpdateSummary(summaries expense.Summary) error {
	summary, err := json.MarshalIndent(summaries, "", " ")
	if err != nil {
		return err
	}
	err = os.WriteFile(summaryPath,summary,0644)
	if err!=nil{
		return err
	}
	return nil
}

func LoadSummary()(expense.Summary,error){
	if _,err:=os.Stat(summaryPath);os.IsNotExist(err){
		return expense.Summary{},nil
	}

	data,err:=os.ReadFile(summaryPath)
	if err!=nil{
		return expense.Summary{}, err
	}

	var summaries expense.Summary
	if len(data) == 0 {	
		return expense.Summary{}, nil
	}
	err = json.Unmarshal(data, &summaries)
	if err != nil {
		return expense.Summary{}, err
	}
	return summaries, nil
}