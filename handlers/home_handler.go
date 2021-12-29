package Handlers

import (
	"html/template"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("static/template/default_template.html", "static/template/home_template.html")
	tmpl.ExecuteTemplate(w, "layout", nil)
}
