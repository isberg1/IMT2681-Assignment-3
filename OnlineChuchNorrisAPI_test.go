package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/heroku/chuch/go-getting-started/handlers"
	"net/http"
	"testing"
)


// online test returns correct result, but reports 0 % test courage
func Test_getChuckNorrisJoke(t *testing.T)  {
	expextedStatusCode := http.StatusOK

	webSite:= "https://warm-meadow-53471.herokuapp.com/dialogflow"
	contentType:="application/json"
	querry := handlers.Querry{handlers.Para{B:"joke"}}
	strc := handlers. DialogflowPostStruct{ResponseId:"testid",FulfillmentText: "joke",QueryResult:querry}


	str, err := json.Marshal(strc)
	if err != nil {
		t.Error("failed to marshal string to json")
	}

	res, err2 := http.Post(webSite,contentType,bytes.NewBuffer(str))
	if err2 != nil {
		t.Error("failed to post to /dialogflow")
	}

	if res.StatusCode != expextedStatusCode {
		t.Error("incorrect status code form heroku post",res.StatusCode)
	}

	var responseStruct handlers.DialogFlowResponceStruct

	err3 := json.NewDecoder(res.Body).Decode(&responseStruct)
	if err3 != nil {
		t.Error("unable do unmarshal dialogflow type json message",err3)
	}

	if responseStruct.FulfillmentText == "" || responseStruct.Payload.Slack.Text == "" {
		t.Error("received no content form API")
	}

	fmt.Println(responseStruct.FulfillmentText)


}



