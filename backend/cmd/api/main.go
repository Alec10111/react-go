package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type application struct {
	Domain string
}

func main() {
	// Set application config
	var app application

	// Read from command line (arguments are called flags)

	// Connect to the DB

	app.Domain = "example.com"

	log.Println("Starting application on port", port)

	// Start a web server (go built in http server is production ready)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}

}