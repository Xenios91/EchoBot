package Handlers

import (
	Service "EchoBot/service"
	"encoding/json"
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
		jsonResp, err := json.Marshal(value.Message)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonResp)
	} else {
		fmt.Fprintf(w, "An invalid token [%s] has been submitted, please try again!\n", token)
	}

}
