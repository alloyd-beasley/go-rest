package fdacontrollers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/krdo-93/go-rest.git/server/handlers"
	"github.com/krdo-93/go-rest.git/server/util/httperror"
)

//FdaRootController defines type for controlle wrapper
type FdaRootController func(w http.ResponseWriter, r *http.Request) error

//GetLimit returns number of records corresponding to limit query
func GetLimit(w http.ResponseWriter, r *http.Request) error {
	if r.Method != http.MethodGet {
		return httperror.NewHTTPError(nil, "This HTTP Method is not allowed", 405)
	}

	limit := r.URL.Query().Get("limit")

	if _, err := strconv.Atoi(limit); err != nil {
		return httperror.NewHTTPError(err, "Limit query must be an integer", 400)
	}

	results, err := handlers.GetLimit(limit)

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
func (fn FdaRootController) ServeHTTP(w http.ResponseWriter, r *http.Request) {

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
