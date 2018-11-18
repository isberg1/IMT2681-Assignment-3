package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

/*
"queryResult": {
    "queryText": "joke",
    "action": "actionAndParameters",
    "parameters": {
      "b": "joke"
    },
    "allRequiredParamsPresent": true,
    "fulfillmentMessages": [{
      "text": {
        "text": [""]
      }
    }],
*/

// Para is used for receiving requests form dialogflow
type Para struct {
	B string `json:"b"`
}

// Querry is used for receiving requests form dialogflow
type Querry struct {
	Parameters Para `json:"parameters"`
}

// DialogflowPostStruct is used for receiving requests form dialogflow
type DialogflowPostStruct struct {
	ResponseID      string `json:"responseId"`
	FulfillmentText string `json:"fulfillmentText"`
	QueryResult     Querry `json:"queryResult"`
}

// Dialogflow gets json string form POST body from the dialoflow chat-bot,
// extracts relevant values, then processes the value.
func Dialogflow(w http.ResponseWriter, r *http.Request) {
	// read the Post content
	read, err3 := ioutil.ReadAll(r.Body)
	if err3 != nil {
		http.Error(w, "error reading r.body", 500)
		logging("error reading r.body")
		return
	}
	// unmarshal json object to string
	var str DialogflowPostStruct
	err := json.Unmarshal(read, &str)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		logging("error unmarshaling ")
		return
	}

	postMemory = str.QueryResult.Parameters.B
	// select action based on request parameter
	switch str.QueryResult.Parameters.B {
	case "joke":
		// return a chuch norris joke
		GetChuckNorrisJoke(w, r)

	case "dad":
		// Returns a dad joke.
		GetRandomDadJoke(w, r)

	case "cat gif":
		// Returns a funny cat gif.
		GetFunnyCatGif(w, r)

	case "dog gif":
		// Returns a funny dog gif.
		GetFunnyDogGif(w, r)

	case "hacker gif":
		// Returns a hacker gif.
		GetHackerGif(w, r)

	case "trending gif":
		// Returns a trending gif.
		GetTrendingGif(w, r)

	case "show dog":
		// Shows a random picture of a dog
		ShowDog(w, r)

	case "add dog":
		AddDog(w, r)

	case "adopt":
		AdoptDog(w, r)

	case "how many":
		GetCount(w, r)

	case "show all":
		ShowAllDogs(w, r)

	default:
		http.Error(w, "", http.StatusBadRequest)
		logging("incorrect value sent to Dialogflow switch case function")
		return
	}

	// register the
	Statistic(str.QueryResult.Parameters.B)

	// log request
}

/*
correct jason response to dialogflow for slack

{
  "fulfillmentText": "joke",
  "payload": {
    "slack": {
      "text": "joke"
    }
  }
}
*/

// PayLoade is used to reply to the dialogflow chat-bot
type PayLoade struct {
	Slack SlackMessage `json:"slack"`
}

// SlackMessage is used to reply to the dialogflow chat-bot
type SlackMessage struct {
	Text string `json:"text"`
}

// DialogFlowResponceStruct is used to reply to the dialogflow chat-bot
type DialogFlowResponceStruct struct {
	FulfillmentText string   `json:"fulfillmentText"`
	Payload         PayLoade `json:"payload"`
}

// gets a string as a parameter, formats it to the correct dialogflow
// json format, and sends it back to the dialogflow chat-bot
func postToDialogflow(w http.ResponseWriter, jsonString string) {
	w.Header().Set("content-type", "application/json")

	// make response struct that will be sent back to dialogflow
	var respToDialogflow DialogFlowResponceStruct
	respToDialogflow.FulfillmentText = jsonString

	var slack SlackMessage
	slack.Text = jsonString
	var payload PayLoade
	payload.Slack = slack

	respToDialogflow.Payload = payload

	// write response back to dialogflow
	json.NewEncoder(w).Encode(respToDialogflow)
}
