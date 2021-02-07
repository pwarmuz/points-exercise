package main

import (
	"testing"
)

func TestBalanceList(t *testing.T) {
	verifyOrder := []Transaction{
		{"payer1", 100},
		{"payer2", 101},
		{"payer3", 102},
	}
	balance := NewBalance()
	balance["payer1"] = Collection{Points(100), 0}
	balance["payer2"] = Collection{Points(101), 1}
	balance["payer3"] = Collection{Points(102), 2}

	for i := 0; i <= 10; i++ {
		balanceList := balance.List()
		for i, listed := range balanceList {
			if verifyOrder[i].Payer != listed.Payer {
				t.Error("incomparable balance in index", i, " expected:", verifyOrder[i].Payer, " got: ", listed.Payer)
			}
		}
		if len(balanceList) != len(balance) {
			t.Error("unequal balance list length")
		}
	}
}
