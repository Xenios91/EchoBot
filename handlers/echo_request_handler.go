package Handlers

import (
	Service "EchoBot/service"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func echoRequestHandler(w http.ResponseWriter, r *http.Request) {
	var requestParameters url.Values = r.URL.Query()
	var token string = requestParameters.Get("token")

	if value, ok := Service.GetEchoRequestService().GetEchoMap()[token]; ok {
		time.Sleep(time.Duration(time.Duration(value.Delay) * time.Second))

		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(value.Message))
	} else {
		fmt.Fprintf(w, "An invalid token [%s] has been submitted, please try again!\n", token)
	}

}
