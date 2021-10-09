package app

import (
	"errors"
	"fmt"
	"sync"

	"points-ws/points"
	"points-ws/templates"
)

type User map[string]*points.Statement

type App struct {
	RW *sync.RWMutex
	User
	Name string
}

// NewApps uses the functional options pattern to load up different templates
// templates.ProductionDirectory() used to load up test specific directories
// templates.TestDirectory() used to load up test specific directories
func NewApps(opts ...templates.Templates) *App {
	for _, opt := range opts {
		opt()
	}
	return &App{RW: new(sync.RWMutex), User: make(User, 0)}
}

// assign a name to active user
func (a *App) assign(name string) {
	a.Name = name
}

// Create user, also switches to user data if user exists
// assigns the user as a current user
func (a *App) Create(name string) {
	a.RW.Lock()
	_, ok := a.User[name]
	if !ok {
		a.User[name] = points.NewStatement()
	}
	// HTTP is stateless, so this name assignment is stupid
	// but it will suffice since this is run locally without cookie generation or security
	a.assign(name)
	a.RW.Unlock()
}

// Delete a user
// Deletes all transactions and balance data
// assigns "No User" if current user was deleted
func (a *App) Delete(name string) (response string) {
	a.RW.Lock()
	defer a.RW.Unlock()
	_, ok := a.User[name]
	if ok {
		a.User[name].Transactions = nil
		a.User[name].Balance = nil
		delete(a.User, name)
		if a.Name == name {
			fallback := "user"
			_, ok := a.User[fallback]
			if !ok {
				a.User[fallback] = points.NewStatement()
			}
			// HTTP is stateless, so this name assignment is stupid
			// but it will suffice since this is run locally without cookie generation or security
			a.assign(fallback)
		}
		return "Successful"
	}
	return "Not Found"
}

// transactions is current users transactions data
func (a *App) transactions() points.Transactions {
	return a.User[a.Name].Transactions
}

// balance is current users balance data
func (a *App) balance() points.Balance {
	return a.User[a.Name].Balance
}

// Update accepts necessary PointsRecord parameters and returns the PointsRecord
func (a *App) Update(payer string, points int, timestamp string) {
	a.RW.Lock()
	a.User[a.Name].Update(payer, points, timestamp)
	a.RW.Unlock()
}

// CurrentPoints will sum all points within the current users balance
func (a *App) CurrentPoints() (points int) {
	balance := a.ReadBalance()
	for _, n := range balance {
		points += n.Points
	}
	return
}

// Deduction from the transactions and balances
// Presents a transaction list of deductions
// Presented as Payer, Points deducted
func (a *App) Deduction(deduct int) ([]points.Transaction, error) {
	if deduct < 0 {
		return nil, errors.New("deduction must be greater than 0, must be considered positive")
	}

	if a.CurrentPoints() < deduct {
		return nil, errors.New("not enough points to cover this transaction")
	}

	var present []points.Transaction
	timestamps := points.SortedTimestamps(a.transactions())
	fmt.Println(timestamps)
	for _, ts := range timestamps {
		entry := a.transactions()[ts]
		if deduct > 0 {
			var deducted int
			if deduct > entry.Points {
				deducted -= entry.Points
				// Balance needs to be deducted based on payers points
				a.balance().Deduct(entry.Payer, entry.Points)
				delete(a.transactions(), ts)
				deduct -= entry.Points
				fmt.Println(deduct)
			} else {
				deducted -= deduct
				// Balance needs to be deducted based on remaining deduction value
				a.balance().Deduct(entry.Payer, deduct)
				entry.Points -= deduct
				a.transactions()[ts] = entry
				deduct = 0
				fmt.Println(deduct)
			}
			present = append(present, points.Transaction{Payer: entry.Payer, Points: deducted})
			fmt.Println("Deduction Method,", deducted, entry.Payer, deduct, entry.Points)
		} else {
			break
		}
	}
	return present, nil
}

// ReadBalance returns a Transaction list because
// only a sorted Payer and Points are necessary information
func (a *App) ReadBalance() []points.Transaction {
	a.RW.RLock()
	listBalance := a.User[a.Name].ListBalance()
	a.RW.RUnlock()
	return listBalance
}
