package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/krdo-93/go-rest.git/server/handlers"
)

type FDAServer struct {
	http.Handler
}

//Start starts the server on :8080
func NewFDAServer() *FDAServer {
	s := new(FDAServer)
	router := mux.NewRouter()
	
	router.HandleFunc("/gettotal", handlers.GetLimit).Methods("GET")

	s.Handler = router

	return s
}
