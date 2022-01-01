package Handlers

import (
	"net/http"
)

func LoadHandlers() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/echo", echoRequestHandler)
	http.HandleFunc("/createEchoRequest", createEchoRequestHandler)
	http.HandleFunc("/getAPIResponse", getAPIResponse)
	http.Handle("/static/images/", http.StripPrefix("/static/images/", http.FileServer(http.Dir("static/images/"))))
}
