package main

import (
	"testing"
)

var ()

// TestScenario will test both deductions and balance of the senario
// as described in the excerise details
func TestScenario(t *testing.T) {
	wantDeductions := []struct {
		payer  string
		points Points
	}{
		{"DANNON", -100},
		{"UNILEVER", -200},
		{"MILLER COORS", -4700},
	}

	wantBalance := []struct {
		payer  string
		points Points
	}{
		{"DANNON", 1000},
		{"UNILEVER", 0},
		{"MILLER COORS", 5300},
	}

	for x := 0; x <= 10; x++ { // check for randomness
		t.Log("Starting user test scenario ", x)
		users := NewUsers()
		users.Create("test")
		users.Entry("DANNON", 300, "10/31 10AM")
		users.Entry("UNILEVER", 200, "10/31 11AM")
		users.Deduct(200)
		users.Entry("MILLER COORS", 10000, "11/1 2PM")
		users.Entry("DANNON", 1000, "11/2 2PM")

		// Verify that the proper points are deducted and presented correctly
		verifyDeduction, _ := users.Deduct(5000)
		for i, n := range verifyDeduction {
			if wantDeductions[i].payer != n.Payer || wantDeductions[i].points != n.Points {
				t.Error("incorrect transactions order:", i, " wanted (", wantDeductions[i].payer, "/", wantDeductions[i].points, ") got (", n.Payer, "/", n.Points, ")")
			}
		}

		// Verify the payers are in priority order and points are a sum total
		balance := users.ReadBalance()
		for i, n := range balance {
			if wantBalance[i].payer != n.Payer || wantBalance[i].points != n.Points {
				t.Error("incorrect balance order:", i, " wanted (", wantBalance[i].payer, "/", wantBalance[i].points, ") got (", n.Payer, "/", n.Points, ")")
			}
		}
		// Deleting test user
		users.Delete("test")
	}
	t.Log("Completed user test scenario")
}
