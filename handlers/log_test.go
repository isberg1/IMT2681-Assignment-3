package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/isberg1/IMT2681-Assignment-3/database"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_log(t *testing.T) {

	database.Connect()

	querry := Querry{Para{B: "no match in switch case"}, ""}
	strc := DialogflowPostStruct{ResponseID: "testid", FulfillmentText: "no match in switch case", QueryResult: querry}

	// convert to json
	str, err := json.Marshal(strc)
	if err != nil {
		t.Error("failed to marshal string to json")
	}

	// Creates a request that is passed to the handler.
	request, err2 := http.NewRequest("POST", "/dialogflow", bytes.NewBuffer(str))
	if err2 != nil {
		t.Error("failed http.NewRequest", err2)
	}
	// Creates the recorder and router.
	recorder := httptest.NewRecorder()
	router := mux.NewRouter()

	// Tests the function.
	router.HandleFunc("/dialogflow", Dialogflow).Methods("POST")
	router.ServeHTTP(recorder, request)

	// Check the status code is what we expect (200).
	status := recorder.Code
	if status != http.StatusBadRequest {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}
}

func Test_checkLog(t *testing.T) {

	// Creates a request that is passed to the handler.
	request, _ := http.NewRequest("GET", "/log", nil)

	// Creates the recorder and router.
	recorder := httptest.NewRecorder()
	router := mux.NewRouter()

	// Tests the function.
	router.HandleFunc("/log", Log).Methods("GET")
	router.ServeHTTP(recorder, request)

	// Check the status code is what we expect (200).
	status := recorder.Code
	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check if the content-type is what we expect (application/json).
	content := recorder.HeaderMap.Get("content-type")
	if content != "text/plain" {
		t.Errorf("Handler returned wrong content-type: got %s want %s",
			content, "text/plain")
	}

	// Check if returned value is empty.
	actual := recorder.Body.String()
	if actual == "" {
		t.Error("handler returned no data")
	}

	if !strings.Contains(actual, "no match in switch case") {
		t.Error("handler incorrect data")
	}

}
