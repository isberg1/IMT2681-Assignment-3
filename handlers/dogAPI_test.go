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
func Test_db(t *testing.T) {

	database.Connect()
}


// tests the AddDog functionallyty
func Test_addDog(t *testing.T)  {


	// used to chech that a new dog was added
	nr, err := database.FindCount("dogs")
	if err != nil {
		t.Error("fail, cant find nr in collection dogs", err)
	}

	// chech the respose for command "add dog"
	chechResponse(t, "add dog")

	nr2, err2 := database.FindCount("dogs")
	if err2 != nil {
		t.Error("fail, cant find nr2 in collection dogs", err2)
	}

	if nr2 != (nr +1) {
		t.Error("wrong count in collection dogs")
	}

}
// shows a dog image
func Test_showDog(t *testing.T)  {

	chechResponse(t, "show dog")
}

func Test_adoptDog(t *testing.T)  {

	// add a dog to ensure there is a dog in the database
	chechResponse(t, "add dog")

	// used to chech that a new dog was added
	nr1, err := database.FindCount("dogs")
	if err != nil {
		t.Error("fail, cant find nr in collection dogs", err)
	}

	chechResponse(t, "adopt")

	nr2, err1 := database.FindCount("dogs")
	if err1 != nil {
		t.Error("fail, cant find nr in collection dogs", err1)
	}

	if nr2 != (nr1 -1) {
		t.Error("fail, wrong number of dogs in DB", err1)
	}

}

func chechResponse(t *testing.T, command string) {

	contentType := "application/json"
	querry := Querry{Para{B: command}}
	strc := DialogflowPostStruct{ResponseID: "testid", FulfillmentText: command, QueryResult: querry}

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
	// unmarshal response
	err3 := json.NewDecoder(recorder.Body).Decode(&responseStruct)
	if err3 != nil {
		t.Error("unable do unmarshal dialogflow type json message")
	}

	// check to see if there is any content
	if responseStruct.FulfillmentText == "" || responseStruct.Payload.Slack.Text == "" {
		t.Error("received no content form API")
	}
}