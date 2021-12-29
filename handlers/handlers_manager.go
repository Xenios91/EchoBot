package Handlers

import (
	"net/http"
)

func LoadHandlers() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/echo", echoRequestHandler)
	http.HandleFunc("/createEchoRequest", echoRequestGenerator)
}
