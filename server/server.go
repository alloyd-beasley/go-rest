package server

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/alloyd-beasley/go-rest.git/db"

	fdahandlers "github.com/alloyd-beasley/go-rest.git/handlers"
	httperror "github.com/alloyd-beasley/go-rest.git/util"
	"github.com/gorilla/mux"
)

//RootController defines receiver type for ServeHTTP
type RootController func(w http.ResponseWriter, r *http.Request) error

//Server defines type for server
type Server struct {
	router *mux.Router
	db     *sql.DB
}

//Init intializes new server
func Initialize() {
	server := new(Server)
	server.routes()
	server.db = db.Initialize()
	if err := http.ListenAndServe(":5000", server.router); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}

//GetLimit returns number of records corresponding to limit query
func (s *Server) GetLimit(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return httperror.NewHTTPError(nil, "This HTTP Method is not allowed", 405)
	}

	limit := strings.TrimSpace(r.URL.Query().Get("limit"))

	if _, err := strconv.Atoi(limit); err != nil && len(limit) != 0 {
		return httperror.NewHTTPError(err, "Limit query must be an integer", 400)
	}

	results, err := fdahandlers.GetLimit(limit)

	if err != nil {
		return err
	}

	output, err := json.Marshal(results)

	if err != nil {
		return httperror.NewHTTPError(err, "Something went wrong while Unmarshaling response json", 400)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)

	return nil
}

/*
 https://medium.com/@ozdemir.zynl/rest-api-error-handling-in-go-behavioral-type-assertion-509d93636afd

 ServeHTTP method pattern comes from above walkthrough on error handling pattern
*/

func (fn RootController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//Call the controller
	err := fn(w, r)

	//If the controller does not return an error, allow fn controller to write response
	if err == nil {
		return
	}

	log.Printf("An error ocurred in an FDA controller: %v", err)

	//type check for client error
	clientError, ok := err.(httperror.ClientError)

	//if the error is not of type ClientError, we have a server related error, return with 500
	if !ok {
		log.Printf("Error occured on the Server")
		w.WriteHeader(500)
		return
	}

	body, err := clientError.ResponseBody()

	//Something went wrong parsing the error response body, return with 500
	if err != nil {
		log.Printf("Error retrieving client error response body: %v", err)
		w.WriteHeader(500)
		return
	}

	status, headers := clientError.ResponseHeaders()

	for i, v := range headers {
		w.Header().Set(i, v)
	}

	//Finally, write headers and body from error, respond
	w.WriteHeader(status)
	w.Write(body)
}
