package Handlers

import (
	"html/template"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("static/template/default_template.html", "static/template/home_template.html")
	if err := tmpl.ExecuteTemplate(w, "layout", nil); err != nil {
		log.Printf("Error: %s", err)
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}
}
