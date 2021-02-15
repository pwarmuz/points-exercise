package main

import "testing"

func TestSortedTimestamps(t *testing.T) {
	// Create unsorted list
	unsorted := []string{
		"1/1 3AM",
		"1/1 7AM",
		"1/3 10AM",
		"1/13 4AM",
		"1/9 10PM",
		"1/9 7PM",
		"1/7 7PM",
	}

	// Create wanted sorted list
	wantSort := []string{
		"1/3 10AM",
		"1/7 7PM",
		"1/9 7PM",
		"1/9 10PM",
		"1/13 4AM",
	}

	// Build transactions from unsorted list
	transactions := NewTransactions()
	for _, n := range unsorted {
		transactions[parseLayout(n)] = Transaction{"Payer", Points(10)}
	}

	// Get sorted transactions
	sorted := sortedTimestamps(transactions, "1/1 8AM")
	for i, n := range sorted {
		// comparing time.Time
		if n != parseLayout(wantSort[i]) {
			t.Error("expected sort failed", wantSort[i])
		}
	}
}
