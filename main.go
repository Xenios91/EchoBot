package main

import (
	handlers "EchoBot/handlers"
	service "EchoBot/service"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	log.Println("Starting EchoBot server...")
	startServer()
}

func loadHandlers() {
	handlers.LoadHandlers()
}

func startServices() {
	service.RunTimedTask()
}

func startServer() {
	loadHandlers()
	startServices()

	var port string
	if len(os.Args) > 1 {
		if _, err := strconv.Atoi(os.Args[1]); err != nil {
			log.Println(fmt.Sprintf("Port cannot be [%s], assigning default port 8080", os.Args[1]))
		} else {
			port = os.Args[1]
		}
	} else {
		port = "8080"
	}

	log.Printf("EchoBot server started on port %s!\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
