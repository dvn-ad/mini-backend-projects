package service

import (
	"fmt"
	"expense-tracker/internal/storage"
	"expense-tracker/internal/expense"
	"time"
)

// func countSummary()(error){
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

	// Count the total summaries and monthly summary without counting each data in data.json
	summaries,err:=storage.LoadSummary()
	if err!=nil{
		return fmt.Errorf("failed to open summary(%w)",err)
	}
	summaries.Total+=amount
	targetMonth := time.Now().Format("2006-01")
	if summaries.Monthly == nil{
		summaries.Monthly=make(map[string]expense.MonthlySummary)
	}
	if monthly, exists := summaries.Monthly[targetMonth]; exists {
		monthly.Total += amount 
		monthly.Count += 1
		summaries.Monthly[targetMonth] = monthly
	} else {
		summaries.Monthly[targetMonth] = expense.MonthlySummary{
			Total: amount,
			Count: 1,
		}
	}
	err=storage.UpdateSummary(summaries)
	if err!=nil{
		return fmt.Errorf("failed to update summary (%w)",err)
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

func GetMonthlySummary(month string)(expense.MonthlySummary,error){
	summaries,err:=storage.LoadSummary()
	if err!=nil{
		fmt.Printf("invalid month")
		return expense.MonthlySummary{},err
	}
	total:=summaries.Monthly[month]
	return total,nil
}

func DeleteExpense(id int)error{
	expenses,err:=storage.LoadExpenses()
	if err!=nil{
		return err
	}
	var targetMonth string
	var targetAmount int
	found:=false
	for i,j:=range expenses{
		if j.ID == id{
			expenses=append(expenses[:i],expenses[i+1:]...)
			targetMonth=j.Date.Format("2006-01")
			targetAmount=j.Amount
			found=true
		}
	}
	if !found{
		return fmt.Errorf("expense with id %d not found",id)
	}
	err=storage.SaveExpenses(expenses)
	if err!=nil{
		return fmt.Errorf("failed to save expenses (%w)",err)
	}

	// handle Summary
	summaries,err:=storage.LoadSummary()
	if err!=nil{
		return fmt.Errorf("failed to load summary (%w)",err)
	}
	summaries.Total-=targetAmount
	if monthly, exists :=summaries.Monthly[targetMonth]; exists {
        monthly.Total -= targetAmount
        monthly.Count -= 1
        summaries.Monthly[targetMonth] = monthly //
    }
    err = storage.UpdateSummary(summaries)
	if err!=nil{
		return fmt.Errorf("failed to save summary (%w)",err)
	}
	return nil
}
