package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_FrontPage(t *testing.T)  {

	// Creates a request that is passed to the handler.
	request, err2 := http.NewRequest("GET", "/", nil)
	if err2 != nil {
		t.Error("failed http.NewRequest", err2)
	}
	// Creates the recorder and router.
	recorder := httptest.NewRecorder()
	router := mux.NewRouter()

	// Tests the function.
	router.HandleFunc("/", Frontpage).Methods("GET")
	router.ServeHTTP(recorder, request)

	// Check the status code is what we expect (200).
	status := recorder.Code
	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check if returned value is empty.
	actual := recorder.Body.String()
	if actual == "" {
		t.Error("handler returned no data")
	}

}
