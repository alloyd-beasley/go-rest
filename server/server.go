package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/krdo-93/go-rest.git/server/handlers"
)

//Server type
type Server struct {
	Router *mux.Router
}

//Start starts the server on :8080
func (s *Server) Start() {
	//TD: should probably make the port value a param in the future

	s.Router = mux.NewRouter()

	s.Router.HandleFunc("/placeholder", handlers.PlaceholderGET).Methods("GET")
	s.Router.HandleFunc("/bydate", handlers.Bydate).Methods("GET")
	s.Router.HandleFunc("/placeholder/post", handlers.PlaceholderPOST).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", s.Router))
	log.Println("server running on port :8080")
}
