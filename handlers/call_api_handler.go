package Handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"regexp"
)

type response struct {
	StatusCode  int
	ContentType string
	Body        string
}

func checkURL(url string) bool {
	regex := "((http|https)://)(www.)?[a-zA-Z0-9@:%._\\+~#?&//=]{2,256}\\.[a-z]{2,6}\\b([-a-zA-Z0-9@:%._\\+~#?&//=]*)"
	match, err := regexp.MatchString(regex, url)

	return match && (err == nil)
}

func submitHTTPRequest(httpMethod, url, contentType string, body io.Reader) (*response, error) {
	var b []byte
	var err error
	var resp *http.Response

	switch httpMethod {
	case "GET":
		if resp, err = http.Get(url); err != nil {
			log.Printf("An error has occured [%s]\n", err)
		} else {
			b, err = io.ReadAll(resp.Body)
			if err != nil {
				log.Printf("An error has occured [%s]\n", err)
			}

			response := &response{StatusCode: resp.StatusCode, ContentType: resp.Header.Get("Content-Type"), Body: string(b)}
			return response, err
		}

	case "POST":
		if resp, err = http.Post(url, contentType, body); err != nil {
			log.Printf("An error has occured [%s]\n", err)
		} else {
			b, err = io.ReadAll(resp.Body)
			if err != nil {
				log.Printf("An error has occured [%s]\n", err)
			}

			response := &response{StatusCode: resp.StatusCode, ContentType: resp.Header.Get("content-type"), Body: string(b)}
			return response, err
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
			http.Error(w, "invalid url", http.StatusBadRequest)
			return
		}

		if resp, err := submitHTTPRequest(httpMethod, url, contentType, body); err != nil {
			http.Error(w, "an error has occurred", http.StatusBadRequest)
		} else {
			jsonResponse, err := json.Marshal(*resp)
			if err != nil {
				http.Error(w, "an error has occurred", http.StatusInternalServerError)
			} else {
				fmt.Fprint(w, string(jsonResponse))
			}
		}
	}
}
