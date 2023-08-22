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
	// set app config
	var app application

	// read from command line

	// connect to db

	app.Domain = "example.com"

	log.Printf("Starting server on port %d", port)

	http.HandleFunc("/", Hello)

	// start server
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatal(err)
	}

}
