package main

import (
	"log"
	"net/http"
)

func main() {
	directory := ".dist"

	fileServer := http.FileServer(http.Dir(directory))

	http.Handle("/", fileServer)

	address := ":8080"

	log.Printf("Starting server on %s", address)
	if err := http.ListenAndServe(address, nil); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
