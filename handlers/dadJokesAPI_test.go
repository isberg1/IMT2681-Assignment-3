// Unit tests for dadjokes API.

package handlers

import (
	"github.com/isberg1/IMT2681-Assignment-3/database"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// Function to test: GetRandomDadJoke().
// Test to check the returned status code, content-type and data for the function.
func Test_GetRandomDadJoke(t *testing.T) {

	database.Connect()

	// Creates a request that is passed to the handler.
	request, _ := http.NewRequest("POST", "/dialogflow", nil)

	// Creates the recorder and router.
	recorder := httptest.NewRecorder()
	router := mux.NewRouter()

	// Tests the function.
	router.HandleFunc("/dialogflow", GetRandomDadJoke).Methods("POST")
	router.ServeHTTP(recorder, request)

	// Check the status code is what we expect (200).
	status := recorder.Code
	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check if the content-type is what we expect (application/json).
	content := recorder.HeaderMap.Get("content-type")
	if content != "application/json" {
		t.Errorf("Handler returned wrong content-type: got %s want %s",
			content, "application/json")
	}

	// Check if returned value is empty.
	actual := recorder.Body.String()
	if actual == "" {
		t.Error("handler returned no data")
	}
}
