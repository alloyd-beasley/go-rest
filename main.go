package main

import (
	"log"
	"net/http"

	"github.com/krdo-93/go-rest.git/server"
)

func main() {
	server := server.NewFDAServer()

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
