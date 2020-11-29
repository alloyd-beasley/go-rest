package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/krdo-93/go-rest.git/server/controllers/fdacontrollers"
)

//FDAServer defines type FDAServer (duh)
type FDAServer struct {
	http.Handler
}

//NewFDAServer creates new server and attaches handlers
func NewFDAServer() *FDAServer {
	s := new(FDAServer)
	router := mux.NewRouter()

	router.Handle("/getlimit", fdacontrollers.FdaRootController(fdacontrollers.GetLimit))

	s.Handler = router

	return s
}
