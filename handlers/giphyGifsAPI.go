// API to get gifs from giphy.

package handlers

import (
	"encoding/json"
	"fmt"
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

// GetRandomGif gets a random gif from giphy.
func GetRandomGif() {

	// Should support searching for gifs.
	URL := "http://api.giphy.com/v1/gifs/search?q=funny+cat&api_key=nZOgnI2vBKhdH5DtG7zZvsdwICp95xO5"

	// Creates a client and sends a new request.
	client := &http.Client{}
	req, _ := http.NewRequest("GET", URL, nil)

	// Set the Accept header to get json format.
	res, err := client.Do(req)
	if err != nil {
		// logging(err) and return
	}

	// Grabs the contents from the GET request.
	output, err := ioutil.ReadAll(res.Body)
	if err != nil {
		// logging(err) and return
	}

	// Closes the response body.
	defer res.Body.Close()

	// Converts from json to struct.
	var gifData gif
	err = json.Unmarshal(output, &gifData)
	if err != nil {
		// logging(err) and return
	}

	// Gets the lenth of the returned array.
	len := len(gifData.Data)

	// Check if there was any data returned.
	if len < 1 {
		// logging(err), errormessage and return.
	}

	// Generates a random number from 0 to len - 1.
	rand.Seed(time.Now().Unix())
	randomNum := rand.Intn(len - 1)

	// Gets the ID for the choisen ID, converts it to string.
	ID := gifData.Data[randomNum]["id"].(string)

	// Creates the URL that should be posted.
	url := string("https://media1.giphy.com/media/" + ID + "/200.gif")

	// Sends the URL to be posted in slack.
	// gif or couldnt find any.
	fmt.Println(url)
}
