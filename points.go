package main

import "time"

const (
	layoutTransaction = "1/2 3PM"
)

// PointsRecord contains all data necessary to manipulate a stored log
type PointsRecord struct {
	Payer       string
	Points      int
	Transaction time.Time
}
