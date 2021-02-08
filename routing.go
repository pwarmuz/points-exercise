package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

const PORT string = ":8080"

type Route struct {
	router          *httprouter.Router
	users           *Users
	TemplatePoints  *template.Template
	TemplateBalance *template.Template
	TemplateDeduct  *template.Template
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
	rt.TemplatePoints = template.Must(template.ParseFiles("html/boilerplate.tmpl", "html/points.gohtml"))
	rt.TemplateBalance = template.Must(template.ParseFiles("html/boilerplate.tmpl", "html/balance.gohtml"))
	rt.TemplateDeduct = template.Must(template.ParseFiles("html/boilerplate.tmpl", "html/deduct.gohtml"))
	rt.router.GET("/", rt.Index)
	rt.router.GET("/balance", rt.Balance)
	rt.router.POST("/add", rt.Add)
	rt.router.POST("/deduct", rt.Deduct)
	rt.router.POST("/user/create", rt.Creation)
	rt.router.POST("/user/delete", rt.Deletion)
	fmt.Printf("Listen and Serve on port %s\n", PORT)
	log.Fatal(http.ListenAndServe(PORT, rt.router))
}

func (rt *Route) Index(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	if req.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
		return
	}
	if len(rt.users.Name) == 0 {
		// use "user" for exercise scenario or create a new user for fresh data
	}
	var current Points
	balance := rt.users.ReadBalance()
	for _, n := range balance {
		current += n.Points
	}
	data := template.FuncMap{
		"current": current,
		"user":    rt.users.Name,
	}
	if err := template.Must(rt.TemplatePoints.Clone()).ExecuteTemplate(w, "boilerplate", data); err != nil {
		log.Fatal(err)
	}
}

func (rt *Route) Balance(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	if req.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
		return
	}
	if len(rt.users.Name) == 0 {
		// use "user" for exercise scenario or create a new user for fresh data

	}
	balance := rt.users.ReadBalance()

	data := template.FuncMap{
		"balance": balance,
	}
	if err := template.Must(rt.TemplateBalance.Clone()).ExecuteTemplate(w, "boilerplate", data); err != nil {
		log.Fatal(err)
	}
}

func (rt *Route) Add(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	if req.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
		return
	}
	payer := strings.ToUpper(template.HTMLEscapeString(req.FormValue("payer")))
	timestamp := template.HTMLEscapeString(req.FormValue("timestamp"))
	alphaPoints := template.HTMLEscapeString(req.FormValue("points"))
	points, err := strconv.Atoi(alphaPoints)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
		return
	}

	rt.users.Entry(payer, points, timestamp)
	http.Redirect(w, req, "/balance", http.StatusFound)
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

	entries, err := rt.users.Deduct(Points(points))
	if err != nil {
		fmt.Fprint(w, "Deduction Error:", err)
		return
	}

	data := template.FuncMap{
		"entries": entries,
	}
	if err := template.Must(rt.TemplateDeduct.Clone()).ExecuteTemplate(w, "boilerplate", data); err != nil {
		log.Fatal(err)
	}
}

func (rt *Route) Creation(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	if req.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
		return
	}
	username := template.HTMLEscapeString(req.FormValue("username"))

	rt.users.Create(username)
	fmt.Fprintln(w, "Request complete")
}

func (rt *Route) Deletion(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	if req.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
		return
	}
	username := template.HTMLEscapeString(req.FormValue("username"))

	rt.users.Delete(username)
	fmt.Fprintln(w, "Request complete")
}
