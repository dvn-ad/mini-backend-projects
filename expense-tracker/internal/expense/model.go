package expense

import "time"

type Expense struct {
	ID int 				`json:"id"`
	Date time.Time		`json:"date"`
	Description string 	`json:"description"`
	Amount string		`json:"amount"`
}