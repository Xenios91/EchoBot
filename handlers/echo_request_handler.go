package Handlers

import (
	EchoRequest "EchoBot/echo_request"
	Service "EchoBot/service"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type echoURL struct {
	URL string
}

func echoRequestHandler(w http.ResponseWriter, r *http.Request) {
	var requestParameters url.Values = r.URL.Query()
	var token string = requestParameters.Get("token")

	if value, ok := Service.GetEchoRequestService().GetEchoMap()[token]; ok {
		time.Sleep(time.Duration(time.Duration(value.Delay) * time.Second))

		w.Header().Set("Content-Type", value.ContentType)
		w.Write([]byte(value.Message))
	} else {
		errorMessage := fmt.Sprintf("An invalid token [%s] has been submitted, please try again!\n", token)
		http.Error(w, errorMessage, http.StatusBadRequest)
	}

}

func createEchoRequestHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		tmpl, _ := template.ParseFiles("static/template/default_template.html", "static/template/echo_request.html")
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
		var performance string = r.FormValue("performance")
		var contentType string = r.FormValue("content-type")
		var requestBody string = r.FormValue("responseBodyRequested")

		echoRequestService := Service.GetEchoRequestService()
		echoRequest := EchoRequest.New(strings.Split((r.RemoteAddr), ":")[0], requestBody, contentType, performance)

		token := echoRequestService.AddToMap(echoRequest)
		requestURL := fmt.Sprintf("http://localhost:8080/echo?token=%s", token)

		tmpl, _ := template.ParseFiles("static/template/default_template.html", "static/template/echo_response.html")
		echoURL := echoURL{URL: requestURL}

		if err := tmpl.ExecuteTemplate(w, "layout", echoURL); err != nil {
			log.Printf("Error: %s", err)
			http.Error(w, "ParseFormError", http.StatusInternalServerError)
		}
	}
}
