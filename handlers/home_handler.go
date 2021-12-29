package Handlers

import (
	"html/template"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("static/template/home_template.html"))
	tmpl.Execute(w, nil)
}
