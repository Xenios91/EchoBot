package Handlers

import (
	"fmt"
	"net/http"
	"net/url"
)

func home_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home Page")
}

func echo_request_handler(w http.ResponseWriter, r *http.Request) {
	var requestParameters url.Values = r.URL.Query()
	var token string = requestParameters.Get("token")
	if len(token) != 24 {
		if len(token) == 0 {
			token = "N/A"
		}
		fmt.Fprintf(w, "An invalid token [%s] has been submitted, please try again!\n", token)
	}
}

func echo_request_generator(w http.ResponseWriter, r *http.Request) {
	//todo
}

func Load_handlers() {
	http.HandleFunc("/", home_handler)
	http.HandleFunc("/echo", echo_request_handler)
	http.HandleFunc("/createEchoRequest", echo_request_generator)
}
