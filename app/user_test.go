package app

import (
	"testing"

	"points-ws/templates"
)

func TestUserCreation(t *testing.T) {
	app := NewApps(templates.TestDirectory())
	app.Create("user")
	app.Create("user")
	app.Create("user2")
	for i := range app.User {
		t.Log("Name", i)
		if i != "user" && i != "user2" {
			t.Error("failed to create user and user2")
		}
	}
}
func TestUserDeletion(t *testing.T) {
	app := NewApps(templates.TestDirectory())
	// create unique test user named doodles
	app.Create("doodles")
	// enter generic entry information
	app.Update("DANNON", 400, "10/31 10AM")
	for i := range app.User {
		// ensure doodles exists
		if i != "doodles" {
			t.Error("doodles should exist at this point")
		}
	}
	// deduct points to prove doodles has transaction data
	deduction, err := app.Deduction(300)
	if err != nil || deduction == nil {
		t.Error("deduction of points failed to list")
	}

	// actually delete doodles
	app.Delete("doodles")
	// recheck transactions to ensure deletion

	_, ok := app.User["doodles"]
	if ok {
		t.Error("doodles should not exist at this point")
	}
}
