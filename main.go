package main

import (
	handlers "EchoBot/handlers"
	service "EchoBot/service"
	"flag"
	"fmt"
	"log"
	"net/http"
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

	var port = flag.Int("port", 8080, "Port number for EchoBot to run on. [Example: -port=8888")

	startServices()

	log.Printf("EchoBot server started on port %d!\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
