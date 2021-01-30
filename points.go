package main

import (
	"fmt"
	"time"
)

const (
	layoutTransaction = "1/2 3PM"
)

// PointsRecord contains all data necessary to manipulate a stored log
type PointsRecord struct {
	Payer       string
	Points      int
	Transaction time.Time
}

// Entry accepts necessary PointsRecord parameters and returns the PointsRecord
func Entry(payer string, points int, timestamp string) PointsRecord {
	t, err := time.Parse(layoutTransaction, timestamp)
	if err != nil {
		fmt.Println(err)
	}
	return PointsRecord{
		Payer:       payer,
		Points:      points,
		Transaction: t,
	}
}
