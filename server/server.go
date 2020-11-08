package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Server type
type Server struct {
	Router *mux.Router
}

//Start starts the server on :8080
func (s *Server) Start() {
	//TD: should probably make the port value a param in the future

	s.Router = mux.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", s.Router), )
	log.Println("server running on port :8080")
}
