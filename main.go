// Author Phil Warmuz
// Programming exercise to develop a http web service, present and manipulate points

package main

func main() {
	pointsMUX := NewRoute()
	pointsMUX.Register()
}
