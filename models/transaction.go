package models

type Transaction struct {
	Type     string  `json:"type"`
	Amount   float64 `json:"amount"`
	ParentId *int64  `json:"parent_id,omitempty"`
}

var Transactions = make(map[int]Transaction)
