// Author Phil Warmuz
// Programming exercise to develop a http web service, present and manipulate points

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	PORT string = ":8080"
)

func journal(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	for _, v := range ENTRIES {
		fmt.Fprintf(w, "payer %s, %d points, %s\n", v.Payer, v.Points, v.Transaction)
	}
}

func index(w http.ResponseWriter, req *http.Request, ps httprouter.Params) {
	if req.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusNotAcceptable), http.StatusNotAcceptable)
		return
	}
	fmt.Fprint(w, "index page")
}

func main() {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, "Not Found")
	})
	router.GET("/", index)
	router.GET("/log", journal)
	fmt.Printf("Listen and Serve on port %s\n", PORT)
	log.Fatal(http.ListenAndServe(PORT, router))
}
