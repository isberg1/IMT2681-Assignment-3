// API to get gifs from giphy.

package handlers

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

// Metadata about a gif.
type gifMeta struct {
	Status     int    `json:"status"`
	Msg        string `json:"msg"`
	ResponseID string `json:"response_id"`
}

// Data about the gif returned.
type gif struct {
	Data       []map[string]interface{} `json:"data"`
	Pagination map[string]int           `json:"pagination"`
	Meta       gifMeta                  `json:"meta"`
}

// Sets the URL for the funny cat gifs and sends it to API function.
func getFunnyCatGif(w http.ResponseWriter, r *http.Request) {
	URL := "http://api.giphy.com/v1/gifs/search?q=funny+cat&api_key=nZOgnI2vBKhdH5DtG7zZvsdwICp95xO5"
	getGif(w, r, URL)
}

// Sets the URL for the funny dog gifs and sends it to API function.
func getFunnyDogGif(w http.ResponseWriter, r *http.Request) {
	URL := "http://api.giphy.com/v1/gifs/search?q=funny+dog&api_key=nZOgnI2vBKhdH5DtG7zZvsdwICp95xO5"
	getGif(w, r, URL)
}

// Sets the URL for the hacker gifs and sends it to API function.
func getHackerGif(w http.ResponseWriter, r *http.Request) {
	URL := "http://api.giphy.com/v1/gifs/search?q=hacker&api_key=nZOgnI2vBKhdH5DtG7zZvsdwICp95xO5"
	getGif(w, r, URL)
}

// Sets the URL for the trending gifs and sends it to API function.
func getTrendingGif(w http.ResponseWriter, r *http.Request) {
	URL := "http://api.giphy.com/v1/gifs/trending?&api_key=nZOgnI2vBKhdH5DtG7zZvsdwICp95xO5"
	getGif(w, r, URL)
}

// Sets the URL for the random gifs and sends it to API function.
func getRandomGif(w http.ResponseWriter, r *http.Request) {
	URL := "http://api.giphy.com/v1/gifs/random?&api_key=nZOgnI2vBKhdH5DtG7zZvsdwICp95xO5"
	getGif(w, r, URL)
}

// Gets a gif from giphy.
func getGif(w http.ResponseWriter, r *http.Request, URL string) {

	// Creates a client and sends a new request.
	client := &http.Client{}
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		http.Error(w, "error from giphy API", 404)
		logging("error from giphy API")
		return
	}

	// Set the Accept header to get json format.
	res, err := client.Do(req)
	if err != nil {
		http.Error(w, "error from giphy API", 404)
		logging("error from giphy API")
		return
	}

	// Grabs the contents from the GET request.
	output, err := ioutil.ReadAll(res.Body)
	if err != nil {
		http.Error(w, "could not read response body", 404)
		logging("error could not read response body")
		return
	}

	// Closes the response body.
	defer res.Body.Close()

	// Converts from json to struct.
	var gifData gif
	err = json.Unmarshal(output, &gifData)
	if err != nil {
		http.Error(w, "could not unmarshal json", 500)
		logging("error could not unmarshal json")
		return
	}

	// Gets the lenth of the returned array.
	len := len(gifData.Data)

	// Check if there was any data returned.
	if len < 1 {
		http.Error(w, "could not retrieve any content", 404)
		logging("error could not retrieve any content")
		return
	}

	// Generates a random number from 0 to len - 1.
	rand.Seed(time.Now().Unix())
	randomNum := rand.Intn(len - 1)

	// Gets the ID for the choisen ID, converts it to string.
	ID := gifData.Data[randomNum]["id"].(string)

	// Creates the URL that should be posted.
	url := string("https://media1.giphy.com/media/" + ID + "/200.gif")

	// Sends the gif back to dialogflow.
	postToDialogflow(w, url)
}
