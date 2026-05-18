package expense

import "time"

type Expense struct {
	ID int 				`json:"id"`
	Date time.Time		`json:"date"`
	Description string 	`json:"description"`
	Amount int			`json:"amount"`
}

// type Summary struct{
// 	Expenses int 	`json:"expenses"`
// 	UpdatedAt time.Time `json:"UpdatedAt`
// }

type MonthlySummary struct {
    Total int `json:"total"`
    Count int `json:"count"`
}

type Summary struct {
    Total     int                      	`json:"total"`
    Monthly   map[string]MonthlySummary	`json:"monthly"`
    UpdatedAt time.Time                	`json:"updated_at"`
}