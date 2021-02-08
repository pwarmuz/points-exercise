package main

// scenario is an implied state based on the exercise document
// user is created as "user"
// Dannon and Unilever deposit points
// a deduction of 200 is made
// Miller Coors and Dannon deposit points
// the user can then manually deduct 5000 points to finish up the scenario
func scenario() *Users {
	users := NewUsers()
	users.Create("user")
	users.Entry("DANNON", 300, "10/31 10AM")
	users.Entry("UNILEVER", 200, "10/31 11AM")
	users.Deduct(200)
	users.Entry("MILLER COORS", 10000, "11/1 2PM")
	users.Entry("DANNON", 1000, "11/2 2PM")
	return users
}
