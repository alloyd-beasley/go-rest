package httperror

/*
 https://medium.com/@ozdemir.zynl/rest-api-error-handling-in-go-behavioral-type-assertion-509d93636afd

 Link to walkthrough for this error typing pattern
*/

import (
	"encoding/json"
)

/*
	ClientError defines methods for custom errors.
	Also implements Error method, inheriting from Go error type
*/
type ClientError interface {
	Error() string
	ResponseBody() ([]byte, error)
	ResponseHeaders() (int, map[string]string)
}

//HTTPError defines our custom error
type HTTPError struct {
	Actual  error
	Message string
	Status  int
}

//Error defines error method for HTTPError, HTTPError inherits from Error interface now
func (e *HTTPError) Error() string {
	if e.Actual == nil {
		return e.Message
	}

	return e.Error()
}

//ResponseBody returns body of error if provided, or generic error
func (e *HTTPError) ResponseBody() ([]byte, error) {
	body, err := json.Marshal(e)

	if err != nil {
		return nil, err
	}

	return body, nil
}

//ResponseHeaders returns the Status Code of the response and a map of the headers
func (e *HTTPError) ResponseHeaders() (int, map[string]string) {
	return e.Status, map[string]string{
		"Content-Type": "application/json; charset=utf-8",
	}
}

//NewHTTPError returns the custom HTTP Error
func NewHTTPError(err error, message string, status int) error {
	return &HTTPError{
		err,
		message,
		status,
	}
}
