package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetTotal(t *testing.T) {
	server := NewFDAServer()

	t.Run("Response code 200", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/gettotal", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		_, err := ioutil.ReadAll(response.Body)
		
		if err != nil {
			t.Errorf("Error reading response body v%v: ", err)
		}

		if response.Code != 200 {
			fmt.Printf("CODE %v", response.Code)
			t.Errorf("response not 200")
		}		
	})
}
