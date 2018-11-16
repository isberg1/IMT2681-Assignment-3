package handlers

import (
	"encoding/json"
	"net/http"
)

/*
chuck norris api jason: response

{
"icon_url" : "https://assets.chucknorris.host/img/avatar/chuck-norris.png",
"id" : "vuSpFINERKarK6mXqQrbhg",
"url" : "https://api.chucknorris.io/jokes/vuSpFINERKarK6mXqQrbhg"
"value" : "The Mona Lisa is based on a peice of toilet paper used by Chuck Norris."
}
*/
//Chuch struct is used for receiving json strings form Chuch Norris API
type Chuch struct {
	Icon_url string `json:"icon_url"`
	ID       string `json:"id"`
	URL      string `json:"url"`
	Value    string `json:"value"`
}

// gets a Chuch Norris json string, extracts the relevant value from it and
// sends it to be formatted for dialogflow
func GetChuckNorrisJoke(w http.ResponseWriter, r *http.Request) {

	res, err1 := http.Get("https://api.chucknorris.io/jokes/random")
	if err1 != nil {
		http.Error(w, "error form chuchnorris api", 500)
		logging("error form chuchnorris api")
		return
	} else {
		defer res.Body.Close()
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
