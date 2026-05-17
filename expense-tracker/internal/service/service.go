package service

import (
	"fmt"
	"expense-tracker/internal/storage"
	"expense-tracker/internal/expense"
	"time"
)

func AddExpense(desc string, amount int)error{
	expenses,err:=storage.LoadExpenses()
	if err!=nil{
		return fmt.Errorf("failed to load data %w", err)
	}
	maxID:=0
	for _,e:=range expenses{
		if e.ID>maxID{
			maxID=e.ID
		}
	}
	newExpense:=expense.Expense{
		ID:maxID+1,
		Description: desc,
		Amount: amount,
		Date: time.Now(),
	}
	expenses = append(expenses, newExpense)
	err=storage.SaveExpenses(expenses)
	if err!=nil{
		return fmt.Errorf("failed to save %w", err)
	}
	fmt.Printf("Expense added successfully (ID: %d)\n",maxID+1)
	return nil
}

func ListExpenses()([]expense.Expense,error){
	expenses,err:=storage.LoadExpenses()
	if err!=nil{
		return nil,fmt.Errorf("failed to load data %w", err)
	}
	return expenses,nil
}