package Handlers

import (
	EchoRequest "EchoBot/echo_request"
	Service "EchoBot/service"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type echoUrl struct {
	Url string
}

func createEchoRequestHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	var performance string = r.FormValue("performance")
	var contentType string = r.FormValue("content-type")
	var requestBody string = r.FormValue("responseBodyRequested")

	echoRequestService := Service.GetEchoRequestService()
	echoRequest := EchoRequest.EchoRequest{Ip: strings.Split((r.RemoteAddr), ":")[0], Message: requestBody}
	echoRequest.SetPerformance(performance)
	echoRequest.SetContentType(contentType)

	token := echoRequestService.AddToMap(&echoRequest)
	requestUrl := fmt.Sprintf("http://localhost:8080/echo?token=%s", token)

	tmpl, _ := template.ParseFiles("static/template/default_template.html", "static/template/response_template.html")
	echoUrl := echoUrl{Url: requestUrl}

	if err := tmpl.ExecuteTemplate(w, "layout", echoUrl); err != nil {
		log.Printf("Error: %s", err)
	}
}
