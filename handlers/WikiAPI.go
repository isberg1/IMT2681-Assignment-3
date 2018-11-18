package handlers

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// GetWiki possesses a wikipedia search request
func GetWiki(w http.ResponseWriter, r *http.Request, s string) {
	// Receives the search string and normalizes the query.(Exchanges spaces with '_')
	s = strings.Replace(s, "search", " ", 1)
	s = strings.Replace(s, " ", "_", -1)
	// Appends the normalized query to the URL of the http request.
	url := "https://en.wikipedia.org/w/api.php?action=opensearch&limit=1&format=json&search=" + s

	// Sets up a new request with the correct headers; JSON formatted response.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		logging("could not retrieve wikipage, 400")
		return
	}

	req.Header.Set("Content-Type", "application/json")
	// Opens a client towards the API so that the headers information can be passed through.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "", http.StatusNotFound)
		logging("could not retrieve wikipage, 404")
		return
	}
	defer resp.Body.Close()
	// Transfers the http body content to a variable that is converted to string.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		logging("could not retrieve wikipage, 500")
		return
	}
	bodyStr := string(body)
	// Splits out the string into slices and removes all unnecessary text.
	convertedStr := strings.Split(bodyStr, "[")
	finalStr := strings.Trim(string(convertedStr[3]), "],")
	// Sends the resulting string to the dialogflowHandler.
	postToDialogflow(w, finalStr)
}
