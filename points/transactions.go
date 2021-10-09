package points

import (
	"fmt"
	"time"
)

// Points representation
//type Points int

// Transaction is used to identify a specific entry
// It requires a payer and a points value
type Transaction struct {
	Payer  string
	Points int
}

// Transactions are the organization of a transaction entry
// based on it's creation via timestamp
type Transactions map[time.Time]*Transaction

// Credit is a grouping of points and it's priority
// The Priority is used to identify where the entry belongs
type Credit struct {
	Points   int
	Priority int
}

// Balance revolves around the payer's name
// It manipulates the Credit
//		- the Points will be a sum
//    - the Priority is where the payer's name shows up in a list
type Balance map[string]*Credit

// Statement stores information for the transactions and balances
type Statement struct {
	Transactions
	Balance
}

// NewTransactions initializes an empty Transactions
func NewTransactions() Transactions {
	t := make(Transactions)
	return t
}

// NewBalance initializes an empty Balance
func NewBalance() Balance {
	b := make(Balance)
	return b
}

// NewStatement initializes an empty Balance
func NewStatement() *Statement {
	return &Statement{Transactions: NewTransactions(), Balance: NewBalance()}
}

// Update the priority and points of the payer within the collection
// priority determines the order in which the payer is presented
func (b Balance) Update(payer string, points int) {
	_, ok := b[payer]
	if !ok {
		b[payer] = &Credit{Priority: len(b)}
	}
	b[payer].Points += points
}

// Update the priority and points of the payer within the collection
// priority determines the order in which the payer is presented
func (b Balance) Deduct(payer string, points int) {
	_, ok := b[payer]
	if !ok {
		return
	}
	b[payer].Points -= points
}

// Update the priority and points of the payer within the collection
// priority determines the order in which the payer is presented
func (s Statement) Update(payer string, points int, timestamp string) {
	ts := ParseLayout(timestamp)
	s.Transactions[ts] = &Transaction{payer, points}
	s.Balance.Update(payer, points)
}

// ListBalance the balance based on priority
// Transaction is used output since payer and points are only necessary
func (s Statement) ListBalance() []Transaction {
	listTransaction := make([]Transaction, len(s.Balance))
	for payers, c := range s.Balance {
		fmt.Println(c.Priority, payers, c.Points)
		listTransaction[c.Priority] = Transaction{payers, c.Points}
	}
	return listTransaction
}
