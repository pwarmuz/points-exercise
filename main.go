// Author Phil Warmuz
// Programming exercise to develop a http web service, present and manipulate points

package main

const port string = ":8080"

func main() {
	pointsMUX := NewRoute(port)
	pointsMUX.Dispatch()
}
