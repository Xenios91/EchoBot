package Handlers

import (
	EchoRequest "EchoBot/echo_request"
	Service "EchoBot/service"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func echoRequestHandler(w http.ResponseWriter, r *http.Request) {
	var requestParameters url.Values = r.URL.Query()
	var token string = requestParameters.Get("token")
	var echoRequestMessage string

	if value, ok := Service.GetEchoRequestService().GetEchoMap()[token]; ok {
		echoRequestMessage = *value.Message
		fmt.Fprintln(w, echoRequestMessage)
	} else {
		fmt.Fprintf(w, "An invalid token [%s] has been submitted, please try again!\n", token)
	}
}

func echoRequestGenerator(w http.ResponseWriter, r *http.Request) {
	var requestParameters url.Values = r.URL.Query()
	var performance string = requestParameters.Get("performance")

	if body, err := ioutil.ReadAll(r.Body); err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	} else {
		requestBody := string(body)
		echoRequestService := Service.GetEchoRequestService()
		echoRequest := EchoRequest.EchoRequest{Ip: &strings.Split((r.RemoteAddr), ":")[0], Message: &requestBody}
		echoRequest.SetPerformance(performance)
		token := echoRequestService.AddToMap(&echoRequest)
		fmt.Fprintln(w, *token)
	}

}
