package main

// Collection is a grouping of points
// The Priority is used to identify where the entry belongs
type Collection struct {
	Points
	Priority int
}

// Balance revolves around the payer's name
// It manipulates a Collection
//		- the Points will be a sum
//    - the Priority is where the payer's name shows up in a list
type Balance map[string]Collection

// NewBalance initializes an empty Balance
func NewBalance() Balance {
	b := make(Balance, 0)
	return b
}

// Deduct the points from the payer and updates the user's collection
func (b Balance) Deduct(payer string, points Points) {
	b[payer] = Collection{b[payer].Points - points, b[payer].Priority}
}

// Update the priority and points of the payer within the collection
// priority determines the order in which the payer is presented
func (b Balance) Update(payer string, points int) Collection {
	_, ok := b[payer]
	if !ok {
		return Collection{Points: b[payer].Points + Points(points), Priority: len(b)}
	}

	return Collection{b[payer].Points + Points(points), b[payer].Priority}
}

// List the balance based on priority
// Transaction is used output since payer and points are only necessary
func (b Balance) List() []Transaction {
	listTransaction := make([]Transaction, len(b))
	for payers, c := range b {
		listTransaction[c.Priority] = Transaction{payers, c.Points}
	}
	return listTransaction
}
