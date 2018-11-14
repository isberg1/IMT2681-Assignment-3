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
type Para struct {
	B string `json:"b"`
}
type Querry struct {
	Parameters Para `json:"parameters"`
}

type DialogflowPostStruct struct {
	ResponseId      string `json:"responseId"`
	FulfillmentText string `json:"fulfillmentText"`
	QueryResult Querry `json:"queryResult"`
}


// http POST handler for "/dialogflow"
func Dialogflow(w http.ResponseWriter, r *http.Request) {
	// read the Post content
	read, err3 := ioutil.ReadAll(r.Body)
	if err3!= nil {
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

	// select action based on request parameter
	switch str.QueryResult.Parameters.B {
	case "joke":
		// return a chuch norris joke
		getChuckNorrisJoke(w, r)
		// log request
		postMemory = append(postMemory,string(read))

	default:
		http.Error(w, "",http.StatusBadRequest)
	}


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

type PayLoade struct {
	Slack SlackMessage `json:"slack"`
}
type SlackMessage struct {
	Text string	`json:"text"`
}

type DialogFlowResponceStruct struct {
	FulfillmentText string	`json:"fulfillmentText"`
	Payload PayLoade		`json:"payload"`
}

func postToDialogflow(w http.ResponseWriter, jsonString string)  {
	w.Header().Set("content-type","application/json")

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