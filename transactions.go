package main

import (
	"time"
)

// Points representation
type Points int

// Transaction is used to identify a specific entry
// It requires a payer and a points value
type Transaction struct {
	Payer string
	Points
}

// Transactions are the organization of a transaction entry
// based on it's creation via timestamp
type Transactions map[time.Time]Transaction

// NewTransactions initializes an empty Transactions
func NewTransactions() Transactions {
	t := make(Transactions, 0)
	return t
}
