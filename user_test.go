package main

import (
	"testing"
)

func TestUserCreation(t *testing.T) {
	users := NewUsers()
	users.Create("user")
	users.Create("user")
	users.Create("user2")
	for i := range users.User {
		t.Log("Name", i)
	}
}
func TestUserDeletion(t *testing.T) {
	users := NewUsers()
	// create unique test user named doodles
	users.Create("doodles")
	// enter generic entry information
	users.Entry("DANNON", 300, "10/31 10AM")
	for i := range users.User {
		// ensure doodles exists
		if i != "doodles" {
			t.Error("doodles should exist at this point")
		}
	}
	// deduct points to prove doodles has transaction data
	deduction, _ := users.Deduct(300)
	if deduction == nil {
		t.Error("entry for doodles should not have been deleted")
	}
	// actually delete doodles
	users.Delete("doodles")
	// recheck transactions to ensure deletion
	afterDeletion, _ := users.Deduct(300)
	if afterDeletion != nil {
		t.Error("entry for doodles should have been deleted")
	}
	for i := range users.User {
		// ensure doodles exists
		if i == "doodles" {
			t.Error("doodles should not exist at this point")
		}
	}
}
