package models

type Transaction struct {
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
}

type TransactionFilters struct {
	Description string
	Category    string

	MinAmount float64
	MaxAmount float64
}
