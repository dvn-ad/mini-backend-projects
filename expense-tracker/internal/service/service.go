package service

import (
	"fmt"
	"expense-tracker/internal/storage"
	"expense-tracker/internal/expense"
	"time"
)

// func countExpenses()(error){
// 	expenses,err:=storage.LoadExpenses()
// 	if err!=nil{
// 		return fmt.Errorf("failed to load data %w", err)
// 	}
// 	val:=0
// 	for _,e:=range expenses{
// 		val+=e.Amount
// 	}
	
// 	newSummmary:=expense.Summary{
// 		Expenses: val,
// 		UpdatedAt: time.Now(),
// 	}
// 	err=storage.UpdateSummmary(newSummmary)
// 	if err!=nil{
// 		return fmt.Errorf("failed to update summary, %w",err)
// 	}
// 	return nil
// }

func countExpenses() error {
    expenses, err := storage.LoadExpenses()
    if err != nil {
        return fmt.Errorf("failed to load data: %w", err)
    }
    total := 0
    // monthly := make(map[string]int)
	monthly := make(map[string]expense.MonthlySummary)
    for _, e := range expenses {
        total += e.Amount
		monthKey := e.Date.Format("2006-01")
		m := monthly[monthKey]
		m.Total += e.Amount
		m.Count += 1
		monthly[monthKey] = m
    }
    newSummary := expense.Summary{
        Total:     total,
        Monthly:   monthly,
        UpdatedAt: time.Now(),
    }
    if err := storage.UpdateSummary(newSummary); err != nil {
        return fmt.Errorf("failed to update summary: %w", err)
    }
    return nil
}

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
	err=countExpenses()
	if err!=nil{
		return err
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

func GetSummary()(expense.Summary,error){
	summaries,err:=storage.LoadSummary()
	if err!=nil{
		return expense.Summary{},fmt.Errorf("failed to load data %w", err)
	}
	return summaries,nil
}

func GetMonthlySummmary(month string)(expense.MonthlySummary,error){
	summaries,err:=storage.LoadSummary()
	if err!=nil{
		fmt.Printf("invalid month")
		return expense.MonthlySummary{},err
	}
	total:=summaries.Monthly[month]
	return total,nil
}
