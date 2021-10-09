package points

import (
	"testing"
)

// TestSortedTimestamps will test both supplemental functions ParseLayout() and SortedTime()
func TestSortedTimestamps(t *testing.T) {
	// Create unsorted list
	cases := struct {
		unsorted []string
		sorted   []string
	}{
		unsorted: []string{"1/3 10AM", "1/13 4AM", "1/9 10PM", "1/9 7PM", "1/7 7PM"},
		sorted:   []string{"1/3 10AM", "1/7 7PM", "1/9 7PM", "1/9 10PM", "1/13 4AM"},
	}

	// Build transactions from unsorted list
	transactions := NewTransactions()
	for _, n := range cases.unsorted {
		// the actual Transaction details doesn't matter
		transactions[ParseLayout(n)] = &Transaction{"Payer", 10}
	}

	// Get sorted transactions
	sorted := SortedTimestamps(transactions)
	for i, n := range sorted {
		// comparing time.Time
		if n != ParseLayout(cases.sorted[i]) {
			t.Error("expected sort failed", cases.sorted[i])
		}
	}
}
