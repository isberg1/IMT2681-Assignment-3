package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/isberg1/IMT2681-Assignment-3/database"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_wikeAPI(t *testing.T) {

	// ensure connection to database
	database.Connect()
	// expected content-type to be sendt back
	contentType := "application/json"
	// make a struck for queering the API
	querry := Querry{Para{B: "search"}, "search norge"}
	strc := DialogflowPostStruct{ResponseID: "testid", FulfillmentText: "search", QueryResult: querry}

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
	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check if the content-type is what we expect (application/json).
	content := recorder.HeaderMap.Get("content-type")
	if content != contentType {
		t.Errorf("Handler returned wrong content-type: got %s want %s",
			content, contentType)
	}
	// struct used for storing reply
	var responseStruct DialogFlowResponceStruct
	// unmarshal from json to struct
	err3 := json.NewDecoder(recorder.Body).Decode(&responseStruct)
	if err3 != nil {
		t.Error("unable do unmarshal dialogflow type json message")
	}
	// check if there is any content in the response
	if responseStruct.FulfillmentText == "" || responseStruct.Payload.Slack.Text == "" {
		t.Error("received no content form API")
	}

}
