package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

const (
	PORT string = ":8080"
)

type Route struct {
	router *httprouter.Router
	users  *Users
}

func NewRoute() *Route {
	return &Route{
		router: httprouter.New(),
		users:  scenario(),
	}
}

func (rt *Route) Register() {
	rt.router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "Not Found")
	})
	rt.router.GET("/", rt.Index)
	rt.router.POST("/add", rt.Add)
	rt.router.POST("/deduct", rt.Deduct)
	fmt.Printf("Listen and Serve on port %s\n", PORT)
	log.Fatal(http.ListenAndServe(PORT, rt.router))
}

func (rt *Route) Index(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	if req.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
		return
	}
	if len(rt.users.name) == 0 {
		// use "user" for exercise scenario or create a new user for fresh data

	}

	fmt.Fprint(w, "index page")
}
func (rt *Route) Add(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	if req.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
		return
	}
	payer := template.HTMLEscapeString(req.FormValue("payer"))
	timestamp := template.HTMLEscapeString(req.FormValue("timestamp"))
	alphaPoints := template.HTMLEscapeString(req.FormValue("points"))
	points, err := strconv.Atoi(alphaPoints)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
		return
	}

	rt.users.Entry(payer, points, timestamp)
}

func (rt *Route) Deduct(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	if req.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
		return
	}
	alphaPoints := template.HTMLEscapeString(req.FormValue("points"))
	points, err := strconv.Atoi(alphaPoints)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
		return
	}

	rt.users.Deduct(Points(points))

}
