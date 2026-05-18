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
		if len(cmdArgs) == 0 {
			summary, err := service.GetSummary()
			if err != nil {
				fmt.Printf("couldn't load summary: %v\n", err)
				return
			}
			fmt.Printf("Total expenses: $%d\n", summary.Total)
			return
		}

		if cmdArgs[0] == "--month" {
			if len(cmdArgs) < 2 {
				fmt.Println("usage: summary --month <1-12>")
				return
			}
			monthInt, err := strconv.Atoi(cmdArgs[1])
			if err != nil || monthInt < 1 || monthInt > 12 {
				fmt.Println("invalid month (1-12)")
				return
			}
			summaries,err:=storage.LoadSummary()
			if err != nil {
				fmt.Printf("failed to load summary (%w)",err)
				return 
			}
			monthKey := fmt.Sprintf("2026-%02d", monthInt)
			amount:=summaries.Monthly[monthKey].Total
			months := []string{
				"", "January", "February", "March", "April",
				"May", "June", "July", "August", "September",
				"October", "November", "December",
			}
			fmt.Printf(
				"Total expenses for %s: $%d\n",
				months[monthInt],
				amount,
			)
			return
		}

	case "delete":
		if cmdArgs[0]=="--id"{
			intId,err:=strconv.Atoi(cmdArgs[1])
			if err!=nil{
				fmt.Printf("invalid id %s\n",cmdArgs[1])
				fmt.Printf("usage: delete --id <id>\new(type)")
				return
			}
			err=service.DeleteExpense(intId)
			if err!=nil{
				fmt.Printf("failed to delete (%v)\n",err)
				return
			}
			fmt.Println("Expense deleted successfully")
		}else{
			fmt.Println("usage: delete --id <id>")
			return
		}
		return
	
	default:
		fmt.Println("invalid command")
	}
}