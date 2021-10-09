package main

import (
	"fmt"
	"log"
	"net/http"
	"points-ws/app"

	"github.com/julienschmidt/httprouter"
)

type Route struct {
	port   string
	router *httprouter.Router
}

func NewRoute(port string) *Route {
	return &Route{
		port:   port,
		router: httprouter.New(),
	}
}

func (rt *Route) Dispatch() {
	rt.router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "Not Found")
	})

	app := app.Scenario()
	rt.router.GET("/", app.Index)
	rt.router.GET("/balance", app.Balance)
	rt.router.POST("/add", app.Add)
	rt.router.POST("/deduct", app.Deduct)
	rt.router.PUT("/user/create", app.Creation)
	rt.router.DELETE("/user/delete", app.Deletion)
	fmt.Printf("Listen and Serve on port %s\n", rt.port)
	log.Fatal(http.ListenAndServe(rt.port, rt.router))
}
