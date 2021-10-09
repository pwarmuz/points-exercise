package templates

import (
	"html/template"
	"log"
	"net/http"
)

var (
	tmplPoints  *template.Template
	tmplBalance *template.Template
	tmplDeduct  *template.Template
)

type Templates func()

func ProductionDirectory() Templates {
	return func() {
		tmplPoints = template.Must(template.ParseFiles("./html/boilerplate.tmpl", "./html/points.gohtml"))
		tmplBalance = template.Must(template.ParseFiles("./html/boilerplate.tmpl", "./html/balance.gohtml"))
		tmplDeduct = template.Must(template.ParseFiles("./html/boilerplate.tmpl", "./html/deduct.gohtml"))
	}
}

func TestDirectory() Templates {
	return func() {
		tmplPoints = template.Must(template.ParseFiles("../html/boilerplate.tmpl", "../html/points.gohtml"))
		tmplBalance = template.Must(template.ParseFiles("../html/boilerplate.tmpl", "../html/balance.gohtml"))
		tmplDeduct = template.Must(template.ParseFiles("../html/boilerplate.tmpl", "../html/deduct.gohtml"))
	}
}

func Points(w http.ResponseWriter, data interface{}) {
	if err := template.Must(tmplPoints.Clone()).ExecuteTemplate(w, "boilerplate", data); err != nil {
		log.Fatal(err)
	}
}
func Balance(w http.ResponseWriter, data interface{}) {
	if err := template.Must(tmplBalance.Clone()).ExecuteTemplate(w, "boilerplate", data); err != nil {
		log.Fatal(err)
	}
}
func Deduct(w http.ResponseWriter, data interface{}) {
	if err := template.Must(tmplDeduct.Clone()).ExecuteTemplate(w, "boilerplate", data); err != nil {
		log.Fatal(err)
	}
}
