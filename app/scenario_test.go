package app

import (
	"fmt"
	"testing"
	"time"

	"points-ws/points"
	"points-ws/templates"
)

// heartbeat will continue to pulse based on the ticker
// the intent is to use this during long tests where there is an uncertainty if the test has frozen due to lack of feedback
// the pulse will maintain a consistent feedback ensuring that the program is still functioning.
type heartbeat struct {
	tick *time.Ticker
	done chan struct{}
}

func newHeartbeat(ticking time.Duration) *heartbeat {
	return &heartbeat{time.NewTicker(ticking * time.Microsecond), make(chan struct{})}
}

func (h *heartbeat) pulse(msg string) {
	for {
		select {
		case <-h.done:
			return
		case t := <-h.tick.C:
			fmt.Println(msg, t)
		}
	}
}

func (h *heartbeat) close() {
	h.tick.Stop()
	h.done <- struct{}{}
	fmt.Println("Heartbeat finished")
}

// TestScenario will test both deductions and balance of the scenario
// as described in the exercise details
func TestScenario(t *testing.T) {
	// the scenario is as follows
	// + [0] DANNON 300
	// + [1] UNILEVER 200
	// then
	// - 200 points
	// which will result in
	// = [0] DANNON 100
	// = [1] UNILEVER 200
	// then
	// + [2] MILLER COORS 10000
	// + [3] DANNON 1000
	// which will result in
	// = [0] DANNON 100
	// = [1] UNILEVER 200
	// = [2] MILLER COORS 10000
	// = [3] DANNON 1000
	// then
	// - 5000 points
	// which will result in
	// = [0] DANNON 0
	// = [1] UNILEVER 0
	// = [2] MILLER COORS 5300
	// = [3] DANNON 1000

	wantDeductions := []points.Transaction{
		{"DANNON", -100},
		{"UNILEVER", -200},
		{"MILLER COORS", -4700}, // 4700
	}

	wantBalance := []points.Transaction{
		{"DANNON", 1000},
		{"UNILEVER", 0},
		{"MILLER COORS", 5300},
	}
	hb := newHeartbeat(10)
	// the pulse is running concurrently an will maintain it's ticking rate of 10 microseconds
	go hb.pulse("!!! Rest assure this is still processing...")
	for x := 0; x <= 10; x++ { // it's worth repeating the test because maps are unordered and a lucky order might falsely trigger a pass
		t.Log("Starting user test scenario ", x)
		app := NewApps(templates.TestDirectory())
		app.Create("test")
		app.Update("DANNON", 300, "10/31 10AM")
		app.Update("UNILEVER", 200, "10/31 11AM")
		app.Deduction(200)
		app.Update("MILLER COORS", 10000, "11/1 2PM")
		app.Update("DANNON", 1000, "11/2 2PM")
		// Verify that the proper points are deducted and presented correctly
		verifyDeduction, _ := app.Deduction(5000)
		for i, n := range verifyDeduction {
			if wantDeductions[i].Payer != n.Payer || wantDeductions[i].Points != n.Points {
				t.Error("incorrect transactions order:", i, " wanted (", wantDeductions[i].Payer, "/", wantDeductions[i].Points, ") got (", n.Payer, "/", n.Points, ")")
			}
		}

		// Verify the payers are in priority order and points are a sum total
		balance := app.ReadBalance()
		for i, n := range balance {
			if wantBalance[i].Payer != n.Payer || wantBalance[i].Points != n.Points {
				t.Error("incorrect balance order:", i, " wanted (", wantBalance[i].Payer, "/", wantBalance[i].Points, ") got (", n.Payer, "/", n.Points, ")")
			}
		}
		// Clean up, Deleting test user
		app.Delete("test")
		t.Log("Deleted user: test")
	}
	t.Log("Completed user test scenario")
	hb.close()
}
