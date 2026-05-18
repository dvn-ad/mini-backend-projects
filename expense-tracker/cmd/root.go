package cmd

import (
	// "container/list"
	"expense-tracker/internal/service"
	"expense-tracker/internal/storage"
	"fmt"
	"os"
	"strconv"
)

func Run() {
	args := os.Args
	if len(args)<2{
		fmt.Println("Usage: expense-tracker <command> [arguments]")
		return
	}

	command :=args[1]
	cmdArgs := args[2:]

	switch command {
	case "add":
		var desc string
		var amount int
		for i:=range 4{
			switch cmdArgs[i] {
			case "--description":
				desc=cmdArgs[i+1]
			case "--amount":
				val,err:=strconv.Atoi(cmdArgs[i+1])
				if err!=nil{
					fmt.Printf("invalid amount")
					return
				}
				amount=val
			}
		}
		err:=service.AddExpense(desc, amount)
		if err!=nil{
			fmt.Printf("failed to laod data")
		}

	case "list":
		expenses,err:=storage.LoadExpenses()
		if err!=nil{
			fmt.Printf("failed to load data")
			return
		}
		fmt.Printf("ID\tDate\t\t\tDescription\tAmount\n")
		for _,e:=range expenses{
			fmt.Printf(
				"%d\t%s\t%s\t\t$%d\n",
				e.ID,e.Date.Format("2006-01-02 15:04"),e.Description,e.Amount,
			)
		}
	case "summary":
		summary,err:=service.ShowSummary()
		if err!=nil{
			fmt.Printf("couldn't load summary")
			return 
		}
		fmt.Printf("Total expenses: $%d\n", summary.Expenses)
	case "delete":
		
	}
}