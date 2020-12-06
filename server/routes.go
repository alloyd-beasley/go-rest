package server

import "github.com/gorilla/mux"

func (s *Server) routes() {
	s.router = mux.NewRouter()
	s.router.Handle("/getlimit", RootController(s.GetLimit))
}
