package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", frontpage).Methods("GET")
	r.HandleFunc("/dialogflow", dialogflow).Methods("POST")
	r.HandleFunc("/oldPosts", oldPosts).Methods("GET")
	r.HandleFunc("/log", log).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := http.ListenAndServe(":"+port, r); err != nil {
		panic(err)
	}
}


//_____________________________________________________________________________________


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
func dialogflow(w http.ResponseWriter, r *http.Request) {
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

//_____________________________________________________________________________________




// prints everything that has previously been posted  to "/dialogflow (used for debugging)
var postMemory []string
func oldPosts(w http.ResponseWriter, r *http.Request) {

	for _,val := range postMemory{
		fmt.Fprintln(w, val)
	}

}

//_____________________________________________________________________________________



// stores all log messages
var logArray []string

// displays all log messages
func log(w http.ResponseWriter, r *http.Request) {

	for _, val:= range logArray{
		fmt.Fprintln(w, val)
	}

}
// adds new log messages to storage
func logging(s string)  {
	logArray = append(logArray, s)
}

// webpage explaining api functionality
func frontpage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "post to /dialogflow ")

}



//_____________________________________________________________________________________



/*
chuch norris api jason: response

{
"icon_url" : "https://assets.chucknorris.host/img/avatar/chuck-norris.png",
"id" : "vuSpFINERKarK6mXqQrbhg",
"url" : "https://api.chucknorris.io/jokes/vuSpFINERKarK6mXqQrbhg"
"value" : "The Mona Lisa is based on a peice of toilet paper used by Chuck Norris."
}
 */
type Chuch struct {
	Icon_url string `json:"icon_url"`
	ID string `json:"id"`
	URL string `json:"url"`
	Value string `json:"value"`
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

func getChuckNorrisJoke(w http.ResponseWriter, r *http.Request)  {

	res, err1 := http.Get("https://api.chucknorris.io/jokes/random")
	if err1 != nil {
		http.Error(w, "error form chuchnorris api", 500)
		logging("error form chuchnorris api")
		return
	}
	var joke Chuch
	err2 := json.NewDecoder(res.Body).Decode(&joke)
	if err2 != nil {
		http.Error(w, "error deconding Joke", 500)
		logging("error deconding Joke")
		return
	}

	// write response back to dialogflow
	postToDialogflow(w, joke.Value)


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