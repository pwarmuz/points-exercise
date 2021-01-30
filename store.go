package main

import "fmt"

// Entries are part of the memory store that will be manipulated
// with addition and deduction. Entries are considered a journal of points transactions
type Entries []PointsRecord

var (
	// ENTRIES global
	ENTRIES Entries
)

func init() {
	ENTRIES = make(Entries, 0)
	reserved()
}

// reserved is the default state of the log as detailed in the exercise document
func reserved() {
	store(Entry("DANNON", 300, "10/31 10AM"))

	store(Entry("UNILEVER", 200, "10/31 11AM"))

	store(Entry("DANNON", -200, "10/31 3PM"))

	store(Entry("MILLER COORS", 10000, "11/1 2PM"))

	store(Entry("DANNON", 1000, "11/2 2PM"))
	fmt.Println("reserved loaded")
}

func store(entry PointsRecord) {
	ENTRIES = append(ENTRIES, entry)
}
