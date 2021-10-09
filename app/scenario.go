package app

import "points-ws/templates"

// Scenario is an implied state based on the exercise document
// user is created as "user"
// Dannon and Unilever deposit points
// a deduction of 200 is made
// Miller Coors and Dannon deposit points
// the user can then manually deduct 5000 points to finish up the scenario
func Scenario() *App {
	app := NewApps(templates.ProductionDirectory())
	app.Create("user")
	app.Update("DANNON", 300, "10/31 10AM")
	app.Update("UNILEVER", 200, "10/31 11AM")
	app.Deduction(200)
	app.Update("MILLER COORS", 10000, "11/1 2PM")
	app.Update("DANNON", 1000, "11/2 2PM")
	return app
}
