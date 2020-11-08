package main

import "github.com/krdo-93/go-rest.git/server"

func main() {
	var server *server.Server = &server.Server{}

	server.Start()
}