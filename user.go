package main

import (
	"errors"
	"sync"
)

// Details is required to further manipulate Transactions and Balances
type Details struct {
	Transactions
	Balance
}
type User map[string]Details
type Users struct {
	mutex *sync.RWMutex
	User  User
	Name  string
}

func NewUsers() *Users {
	return &Users{mutex: new(sync.RWMutex), User: make(User, 0)}
}

func (u *Users) assign(name string) {
	u.Name = name
}
func (u *Users) Create(name string) {
	u.mutex.Lock()
	_, ok := u.User[name]
	if !ok {
		u.User[name] = Details{Transactions: NewTransactions(), Balance: NewBalance()}
	}
	// HTTP is stateless, so this name assignment is stupid
	// but it will suffice since this is run locally without cookie generation
	u.assign(name)
	u.mutex.Unlock()
}

// Delete a user
func (u *Users) Delete(name string) {
	u.mutex.Lock()
	for ts := range u.User[name].Transactions {
		delete(u.User[name].Transactions, ts)
	}
	for payer := range u.User[name].Balance {
		delete(u.User[name].Balance, payer)
	}
	delete(u.User, name)
	u.assign("No User")
	u.mutex.Unlock()
}
func (u *Users) transactions() Transactions {
	return u.User[u.Name].Transactions
}
func (u *Users) balance() Balance {
	return u.User[u.Name].Balance
}

// Entry accepts necessary PointsRecord parameters and returns the PointsRecord
func (u *Users) Entry(payer string, points int, timestamp string) {
	ts := parseLayout(timestamp)
	u.mutex.Lock()
	// Appends entry data into a transaction
	u.transactions()[ts] = Transaction{payer, Points(points)}

	// Updates entry data into the payers balance
	u.balance()[payer] = u.balance().Update(payer, points)

	// Associates data to user
	u.User[u.Name] = Details{u.transactions(), u.balance()}
	u.mutex.Unlock()
}

func (u *Users) CurrentPoints() Points {
	var current Points
	balance := u.ReadBalance()
	for _, n := range balance {
		current += n.Points
	}
	return current
}

// Deduct from the transactions and balances
// Presents a transaction list of deductions
// Presented as Payer, Points deducted
func (u *Users) Deduct(deduct Points) ([]Transaction, error) {
	if deduct < 0 {
		return nil, errors.New("deduction must be greater than 0, must be considered positive")
	}

	if u.CurrentPoints() < deduct {
		return nil, errors.New("not enought points to cover this transaction")
	}

	var present []Transaction
	//userTransactions := u.User[name].Transactions
	timestamps := sortedTimestamps(u.transactions())
	for _, ts := range timestamps {
		entry := u.transactions()[ts]
		if deduct > 0 {
			var deducted Points
			if deduct > entry.Points {
				deducted -= entry.Points
				present = append(present, Transaction{entry.Payer, deducted})
				// Balance needs to be deducted based on payers points
				u.balance().Deduct(entry.Payer, entry.Points)
				delete(u.transactions(), ts)
				deduct -= entry.Points
			} else {
				deducted -= deduct
				present = append(present, Transaction{entry.Payer, deducted})
				// Balance needs to be deducted based on remaining deduction value
				u.balance().Deduct(entry.Payer, deduct)
				entry.Points -= deduct
				u.transactions()[ts] = entry
				deduct = 0
			}
		} else {
			break
		}
	}
	return present, nil
}

// ReadBalance returns a Transaction list because
// only a sorted Payer and Points are necessary information
func (u *Users) ReadBalance() []Transaction {
	u.mutex.RLock()
	listBalance := u.balance().List()
	u.mutex.RUnlock()
	return listBalance
}
