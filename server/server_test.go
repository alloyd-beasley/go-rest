package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func formatResponseCodeSuccessMessage(t *testing.T, message string, code int) {
	t.Helper()

	fmt.Printf("CODE %v for: %v \n", code, message)
}

func formatResponseCodeErrorMessage(t *testing.T, message string, code int) {
	t.Helper()

	t.Errorf("INVALID CODE %v for: %v \n", code, message)
}

func TestGetLimit(t *testing.T) {
	server := new(Server)
	server.routes()

	t.Run("Response code 200 for valid request of default limit", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/getlimit", nil)
		response := httptest.NewRecorder()
		server.router.ServeHTTP(response, request)

		if response.Code != 200 {
			formatResponseCodeErrorMessage(t, "valid request of default limit", response.Code)
		} else {
			formatResponseCodeSuccessMessage(t, "default requst limit", response.Code)
		}
	})

	t.Run("Response code 200 for valid request of limit 2", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/getlimit?limit=2", nil)
		response := httptest.NewRecorder()

		server.router.ServeHTTP(response, request)

		if response.Code != 200 {
			formatResponseCodeErrorMessage(t, "valid request of linit 2", response.Code)
		} else {
			formatResponseCodeSuccessMessage(t, "request of limit 2", response.Code)
		}
	})

	t.Run("Response code 405 for invalid request Method", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/getlimit", nil)
		response := httptest.NewRecorder()

		server.router.ServeHTTP(response, request)

		if response.Code != 405 {
			formatResponseCodeErrorMessage(t, "invalid method request", response.Code)
		} else {
			formatResponseCodeSuccessMessage(t, "invalid method request", response.Code)
		}
	})

	t.Run("Response code 400 for invalid query param", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/getlimit?limit=xxx", nil)
		response := httptest.NewRecorder()

		server.router.ServeHTTP(response, request)

		if response.Code != 400 {
			formatResponseCodeErrorMessage(t, "invalid query param", response.Code)
		} else {
			formatResponseCodeSuccessMessage(t, "invalid query param", response.Code)
		}
	})
}
