package main

import (
	handlers "EchoBot/handlers"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting EchoBot server...")
	start_server()
}

func load_handlers() {
	handlers.Load_handlers()
}

func start_server() {
	load_handlers()
	log.Println("EchoBot server started!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
