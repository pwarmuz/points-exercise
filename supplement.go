package main

import (
	"fmt"
	"sort"
	"time"
)

const (
	layoutTransaction = "1/2 3PM"
)

// parseLayout supports parsing a written timestamp,
// as defined in the layoutTransaction constant.
// It returns Time so it can be sorted
func parseLayout(timestamp string) time.Time {
	t, err := time.Parse(layoutTransaction, timestamp)
	if err != nil {
		fmt.Println(err)
	}
	return t
}

// sortedTimestamps performs the sort based on the current items in the global TRANSACTIONS
func sortedTimestamps(transactions Transactions, expireTS string) []time.Time {
	expired := parseLayout(expireTS)
	timestamps := make([]time.Time, 0)
	for ts := range transactions {
		if ts.After(expired) {
			timestamps = append(timestamps, ts)
		}
	}
	sort.Slice(timestamps, func(i, j int) bool {
		return timestamps[i].Before(timestamps[j])
	})
	return timestamps
}
