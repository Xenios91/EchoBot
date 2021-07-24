package Handlers

import (
	EchoRequest "EchoBot/echo_request"
	Service "EchoBot/service"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home Page")
}

func echoRequestHandler(w http.ResponseWriter, r *http.Request) {
	var requestParameters url.Values = r.URL.Query()
	var token string = requestParameters.Get("token")
	if len(token) != 24 {
		if len(token) == 0 {
			token = "N/A"
		}
		fmt.Fprintf(w, "An invalid token [%s] has been submitted, please try again!\n", token)
	}
	echoRequestMessage := Service.GetEchoRequestService().GetEchoMap()[token].Message
	fmt.Fprintln(w, echoRequestMessage)
}

func echoRequestGenerator(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	fmt.Println(body)
	echoRequestService := Service.GetEchoRequestService()
	echoRequest := EchoRequest.EchoRequest{Ip: "test", Message: "test"}
	token := echoRequestService.AddToMap(echoRequest)
	fmt.Fprintln(w, *token)
}

func Load_handlers() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/echo", echoRequestHandler)
	http.HandleFunc("/createEchoRequest", echoRequestGenerator)
}
