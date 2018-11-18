package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetWiki(w http.ResponseWriter, r *http.Request) {

	srcString := "Norway"
	url := "https://en.wikipedia.org/w/api.php?action=opensearch&limit=1&format=json&search=" + srcString
//	fmt.Println("URL: ", url)

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "404 Page not found", 404)
		return
	}
	defer resp.Body.Close()



	body, err := ioutil.ReadAll(resp.Body)
	bodyStr := string(body)

	convertedStr := strings.Split(bodyStr, "[")
	finalStr := strings.Trim(string(convertedStr[3]), "],")
	fmt.Println(finalStr)

	postToDialogflow(w, finalStr)
}
