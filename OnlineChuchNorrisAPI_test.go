package main

import (
	"bytes"
	"encoding/json"
	"github.com/heroku/chuch/go-getting-started/handlers"
	"net/http"
	"testing"
)

// Test_getChuckNorrisJoke online test,
// returns correct result, but reports 0 % test courage
func Test_getChuckNorrisJoke(t *testing.T) {
	expextedStatusCode := http.StatusOK
	//
	webSite := "https://warm-meadow-53471.herokuapp.com/dialogflow"
	contentType := "application/json"

	// make a struct with test values
	para := handlers.Para{B: "joke"}
	querry := handlers.Querry{para}
	strc := handlers.DialogflowPostStruct{ResponseId: "testid", FulfillmentText: "joke", QueryResult: querry}

	// make json string in correct format( imitating a request from dialogflow chat-bot)
	str, err := json.Marshal(strc)
	if err != nil {
		t.Error("failed to marshal string to json")
	}
	// post to online API
	res, err2 := http.Post(webSite, contentType, bytes.NewBuffer(str))
	if err2 != nil {
		t.Error("failed to post to /dialogflow")
	}
	// check http status code
	if res.StatusCode != expextedStatusCode {
		t.Error("incorrect status code form heroku post", res.StatusCode)
	}

	var responseStruct handlers.DialogFlowResponceStruct
	// unmarshal response form online API
	err3 := json.NewDecoder(res.Body).Decode(&responseStruct)
	if err3 != nil {
		t.Error("unable do unmarshal dialogflow type json message", err3)
	}
	// check if response contains some values
	if responseStruct.FulfillmentText == "" || responseStruct.Payload.Slack.Text == "" {
		t.Error("received no content form API")
	}

}
