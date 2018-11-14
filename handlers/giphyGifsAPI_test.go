// Unit tests for giphy gifs API.

package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// Function to test: getFunnyCatGif().
// Test to check the returned status code, content-type and data for the function.
func Test_getFunnyCatGif(t *testing.T) {
	// Creates a request that is passed to the handler.
	request, _ := http.NewRequest("POST", "/dialogflow", nil)

	// Creates the recorder and router.
	recorder := httptest.NewRecorder()
	router := mux.NewRouter()

	// Tests the function.
	router.HandleFunc("/dialogflow", getFunnyCatGif).Methods("POST")
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

// Function to test: getFunnyDogGif().
// Test to check the returned status code, content-type and data for the function.
func Test_getFunnyDogGif(t *testing.T) {
	// Creates a request that is passed to the handler.
	request, _ := http.NewRequest("POST", "/dialogflow", nil)

	// Creates the recorder and router.
	recorder := httptest.NewRecorder()
	router := mux.NewRouter()

	// Tests the function.
	router.HandleFunc("/dialogflow", getFunnyDogGif).Methods("POST")
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

// Function to test: getHackerGif().
// Test to check the returned status code, content-type and data for the function.
func Test_getHackerGif(t *testing.T) {
	// Creates a request that is passed to the handler.
	request, _ := http.NewRequest("POST", "/dialogflow", nil)

	// Creates the recorder and router.
	recorder := httptest.NewRecorder()
	router := mux.NewRouter()

	// Tests the function.
	router.HandleFunc("/dialogflow", getHackerGif).Methods("POST")
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

// Function to test: getTrendingGif().
// Test to check the returned status code, content-type and data for the function.
func Test_getTrendingGif(t *testing.T) {
	// Creates a request that is passed to the handler.
	request, _ := http.NewRequest("POST", "/dialogflow", nil)

	// Creates the recorder and router.
	recorder := httptest.NewRecorder()
	router := mux.NewRouter()

	// Tests the function.
	router.HandleFunc("/dialogflow", getTrendingGif).Methods("POST")
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
