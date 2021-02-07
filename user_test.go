package main

import (
	"testing"
)

func TestUserCreation(t *testing.T) {
	//var users Users
	users := NewUsers()
	users.Create("user")
	users.Create("user")
	users.Create("user2")
	for i := range users.User {
		t.Log("Name", i)
	}
}
func TestUserDeletion(t *testing.T) {
	//var users Users
	users := NewUsers()
	users.Create("doodles")
	users.Entry("DANNON", 300, "10/31 10AM")
	for i := range users.User {
		if i != "doodles" {
			t.Error("doodles should exist at this point")
		}
	}
	deduction, _ := users.Deduct(300)
	if deduction == nil {
		t.Error("entry for doodles should not have been deleted")
	}
	users.Delete("doodles")
	afterDeletion, _ := users.Deduct(300)
	if afterDeletion != nil {
		t.Error("entry for doodles should have been deleted")
	}
}
