package app

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"points-ws/templates"

	"github.com/julienschmidt/httprouter"
)

func (a *App) Index(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	if req.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
		return
	}
	if len(a.Name) == 0 {
		// use "user" for exercise scenario or create a new user for fresh data
	}

	data := template.FuncMap{
		"points": a.CurrentPoints(),
		"user":   a.Name,
	}
	templates.Points(w, data)
}

func (a *App) Balance(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	if req.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
		return
	}
	if len(a.Name) == 0 {
		// use "user" for exercise scenario or create a new user for fresh data
	}

	data := template.FuncMap{
		"balance": a.ReadBalance(),
	}
	templates.Balance(w, data)
}

func (a *App) Add(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
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

	a.Update(payer, points, timestamp)
	http.Redirect(w, req, "/balance", http.StatusFound)
}

func (a *App) Deduct(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	if req.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
		return
	}
	alphaPoints := template.HTMLEscapeString(req.FormValue("points"))

	pts, err := strconv.Atoi(alphaPoints)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
		return
	}

	entries, err := a.Deduction(pts)
	if err != nil {
		fmt.Fprint(w, "Deduction Error:", err)
		return
	}

	data := template.FuncMap{
		"entries": entries,
	}
	templates.Deduct(w, data)
}

func (a *App) Creation(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	if req.Method != "PUT" {
		http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
		return
	}
	// get requested information from form-URLEncoded Body
	req.ParseForm()
	username := req.FormValue("username")
	fmt.Println(username)
	// Do creation action
	a.Create(username)

	// Present it
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, username+" Created, Request complete")
}

func (a *App) Deletion(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	if req.Method != "DELETE" {
		http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
		return
	}
	// Read JSON from JavaScript as an example
	body, _ := ioutil.ReadAll(req.Body)
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	for k, v := range keyVal {
		fmt.Println(k, v)
	}
	// Do deletion action
	username := keyVal["username"]
	resp := a.Delete(username)

	// Present JSON payload response to JavaScript
	p := struct {
		Action    string `json:"Action"`
		Attempted string `json:"Attempted"`
		Status    string `json:"Status"`
		Username  string `json:"Username"`
		Points    int    `json:"Points"`
	}{
		Action:    "DELETE",
		Attempted: username,
		Status:    resp,
		Username:  a.Name,
		Points:    a.CurrentPoints(),
	}
	// Write it
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(p)
}
