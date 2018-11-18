package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/isberg1/IMT2681-Assignment-3/database"
	"net/http"
	"net/http/httptest"
	"testing"
)

// online test returns correct result, but reports 0 % test courage
func Test_GetChuckNorrisJoke(t *testing.T) {


	database.Connect()

	contentType := "application/json"
	querry := Querry{Para{B: "joke"}, ""}
	strc := DialogflowPostStruct{ResponseID: "testid", FulfillmentText: "joke", QueryResult: querry}

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

	// Check if the content-type is what we expect (text/plain).
	content := recorder.HeaderMap.Get("content-type")
	if content != contentType {
		t.Errorf("Handler returned wrong content-type: got %s want %s",
			content, contentType)
	}

	var responseStruct DialogFlowResponceStruct

	err3 := json.NewDecoder(recorder.Body).Decode(&responseStruct)
	if err3 != nil {
		t.Error("unable do unmarshal dialogflow type json message")
	}

	if responseStruct.FulfillmentText == "" || responseStruct.Payload.Slack.Text == "" {
		t.Error("received no content form API")
	}

	fmt.Println(responseStruct.FulfillmentText)

}
