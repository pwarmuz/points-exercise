package points

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
func ParseLayout(timestamp string) time.Time {
	t, err := time.Parse(layoutTransaction, timestamp)
	if err != nil {
		fmt.Println(err)
	}
	return t
}

// sortedTimestamps performs the sort based on the current items in the global TRANSACTIONS
func SortedTimestamps(transactions Transactions) []time.Time {
	timestamps := make([]time.Time, 0)
	//expired := ParseLayout("11/12 10PM")
	for ts := range transactions {
		//if expired.Before(ts) {
		timestamps = append(timestamps, ts)
		//}
	}
	fmt.Println(timestamps)
	sort.Slice(timestamps, func(i, j int) bool {
		return timestamps[i].Before(timestamps[j])
	})
	return timestamps
}
