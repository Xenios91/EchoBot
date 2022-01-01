package Handlers

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func getAPIResponse(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, _ := template.ParseFiles("static/template/default_template.html", "static/template/api_response.html")
		if err := tmpl.ExecuteTemplate(w, "layout", nil); err != nil {
			log.Printf("Error: %s", err)
			http.Error(w, "Server Error", http.StatusInternalServerError)
		}
	} else {
		if err := r.ParseForm(); err != nil {
			log.Printf("Error: %s", err)
			http.Error(w, "ParseFormError", http.StatusBadRequest)
			return
		}
		var httpMethod string = r.FormValue("http-method")
		var url string = r.FormValue("url")
		var contentType string
		var body io.Reader

		switch httpMethod {
		case "GET":
			if resp, err := http.Get(url); err != nil {
				log.Println("An error has occured")
				http.Error(w, "Server Error", http.StatusInternalServerError)
				return
			}

		case "POST":
			if resp, err := http.Post(url, contentType, body); err != nil {
				log.Println("An error has occured")
				http.Error(w, "Server Error", http.StatusInternalServerError)
				return
			}
		}
	}
}
