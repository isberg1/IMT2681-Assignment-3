// API to get dad jokes.

package handlers

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// GetRandomDadJoke gets a random dad joke as json.
func GetRandomDadJoke(w http.ResponseWriter, r *http.Request) {
	// Creates a client and sends a new request.
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://icanhazdadjoke.com/slack", nil)
	if err != nil {
		http.Error(w, "error from dadjokes API", 404)
		logging("error from dadjokes API")
		return
	}

	// Set the Accept header to get json format.
	req.Header.Set("Accept", "application/json")
	res, err := client.Do(req)
	if err != nil {
		http.Error(w, "error from dadjokes API", 404)
		logging("error from dadjokes API")
		return
	}

	// Grabs the contents from the GET request.
	output, err := ioutil.ReadAll(res.Body)
	if err != nil {
		http.Error(w, "could not read response body", 500)
		logging("error could not read response body")
		return
	}

	// Closes the response body.
	defer res.Body.Close()

	// Chech if the joke was received.
	if strings.Contains(string(output), "\"status\": 404") {
		http.Error(w, "could not retrieve joke", 404)
		logging("error could not retrieve joke")
		return
	}

	// Sends the joke back to dialogflow.
	postToDialogflow(w, string(output))
}
