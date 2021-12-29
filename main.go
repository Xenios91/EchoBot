package main

import (
	handlers "EchoBot/handlers"
	service "EchoBot/service"
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
	startServices()
	log.Printf("EchoBot server started on port %d!\n", 8080)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
