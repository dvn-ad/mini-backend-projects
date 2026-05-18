package expense

import "time"

type Expense struct {
	ID int 				`json:"id"`
	Date time.Time		`json:"date"`
	Description string 	`json:"description"`
	Amount int		`json:"amount"`
}

type Summary struct{
	Expenses int 	`json:"expenses"`
	UpdatedAt time.Time `json:"UpdatedAt`
}