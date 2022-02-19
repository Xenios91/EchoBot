package main

import (
	handlers "EchoBot/handlers"
	service "EchoBot/service"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"gopkg.in/Graylog2/go-gelf.v2/gelf"
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

func grayLogConfig(graylogAddr *string) {
	if len(*graylogAddr) > 0 {
		gelfWriter, err := gelf.NewTCPWriter(*graylogAddr)
		if err != nil {
			log.Fatalf("gelf.NewWriter: %s", err)
		}
		// log to both stderr and graylog2
		log.SetOutput(io.MultiWriter(os.Stderr, gelfWriter))
		log.Printf("logging to stderr & graylog2@'%s'\n", *graylogAddr)
	}
}

func startServer() {
	loadHandlers()

	var graylogServer = flag.String("graylog", "", "The graylog server and port to utilize [Example: -graylog=http:www.graylog.com:8080")
	var port = flag.Int("port", 8080, "Port number for EchoBot to run on. [Example: -port=8888")

	grayLogConfig(graylogServer)
	startServices()

	log.Printf("EchoBot server started on port %d!\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
