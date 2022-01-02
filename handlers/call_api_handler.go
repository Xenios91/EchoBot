package Handlers

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"regexp"
)

func checkURL(url string) bool {
	regex := `https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`

	match, err := regexp.MatchString(regex, url)

	return match && (err == nil)
}

func submitHTTPRequest(httpMethod, url, contentType string, body io.Reader) (*string, error) {
	var b []byte
	var err error
	var resp *http.Response

	switch httpMethod {
	case "GET":
		if resp, err = http.Get(url); err != nil {
			log.Println("An error has occured")
		} else {
			b, err = io.ReadAll(resp.Body)
			if err != nil {
				log.Fatalln(err)
			}

			resp := string(b)
			return &resp, err
		}

	case "POST":
		if resp, err = http.Post(url, contentType, body); err != nil {
			log.Println("An error has occured")
		} else {
			b, err = io.ReadAll(resp.Body)
			if err != nil {
				log.Fatalln(err)
			}

			resp := string(b)
			return &resp, err
		}

	default:
		break
	}
	return nil, errors.New("unknown error")
}

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
		var contentType string = r.FormValue("content-type")
		var body io.Reader

		if !checkURL(url) {
			http.Error(w, "Invalid URL", http.StatusBadRequest)
			return
		}

		if resp, err := submitHTTPRequest(httpMethod, url, contentType, body); err != nil {
			http.Error(w, "An error has occured", http.StatusBadRequest)
		} else {
			fmt.Fprint(w, *resp)
		}
	}
}
