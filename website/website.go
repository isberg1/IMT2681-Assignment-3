package website

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gorilla/mux"
	"github.com/isberg1/IMT2681-Assignment-3/handlers"
)

// Struct containing the json text.
type pageVariable struct {
	Text string
}

// PageVars is the variables that should be sent to the HTML template.
var PageVars pageVariable

// Function to test the chuck norris joke.
func testChuckNorrisJoke() string {
	// Creates a request that is passed to the handler.
	request, _ := http.NewRequest("POST", "/dialogflow", nil)

	// Creates the temporary recorder.
	recorder := httptest.NewRecorder()
	router := mux.NewRouter()

	// Sends the request internally.
	router.HandleFunc("/dialogflow", handlers.GetChuckNorrisJoke).Methods("POST")
	router.ServeHTTP(recorder, request)

	// Gets the respond.
	var returnedText map[string]string
	output := recorder.Body.Bytes()
	json.Unmarshal(output, &returnedText)

	// Returns the text.
	return returnedText["fulfillmentText"]
}

// Function to test the dad joke.
func testDadJoke() string {
	// Creates a request that is passed to the handler.
	request, _ := http.NewRequest("POST", "/dialogflow", nil)

	// Creates the temporary recorder.
	recorder := httptest.NewRecorder()
	router := mux.NewRouter()

	// Sends the request internally.
	router.HandleFunc("/dialogflow", handlers.GetRandomDadJoke).Methods("POST")
	router.ServeHTTP(recorder, request)

	// Gets the respond.
	var returnedText map[string]string
	output := recorder.Body.Bytes()
	json.Unmarshal(output, &returnedText)

	// Returns the text.
	return returnedText["fulfillmentText"]
}

// Function to test the cat gif.
func testCatGif() string {
	// Creates a request that is passed to the handler.
	request, _ := http.NewRequest("POST", "/dialogflow", nil)

	// Creates the temporary recorder.
	recorder := httptest.NewRecorder()
	router := mux.NewRouter()

	// Sends the request internally.
	router.HandleFunc("/dialogflow", handlers.GetFunnyCatGif).Methods("POST")
	router.ServeHTTP(recorder, request)

	// Gets the respond.
	var returnedText map[string]string
	output := recorder.Body.Bytes()
	json.Unmarshal(output, &returnedText)

	// Returns the text.
	return returnedText["fulfillmentText"]
}

// Function to test the dog gif.
func testDogGif() string {
	// Creates a request that is passed to the handler.
	request, _ := http.NewRequest("POST", "/dialogflow", nil)

	// Creates the temporary recorder.
	recorder := httptest.NewRecorder()
	router := mux.NewRouter()

	// Sends the request internally.
	router.HandleFunc("/dialogflow", handlers.GetFunnyDogGif).Methods("POST")
	router.ServeHTTP(recorder, request)

	// Gets the respond.
	var returnedText map[string]string
	output := recorder.Body.Bytes()
	json.Unmarshal(output, &returnedText)

	// Returns the text.
	return returnedText["fulfillmentText"]
}

// Function to test the hacker gif.
func testHackerGif() string {
	// Creates a request that is passed to the handler.
	request, _ := http.NewRequest("POST", "/dialogflow", nil)

	// Creates the temporary recorder.
	recorder := httptest.NewRecorder()
	router := mux.NewRouter()

	// Sends the request internally.
	router.HandleFunc("/dialogflow", handlers.GetHackerGif).Methods("POST")
	router.ServeHTTP(recorder, request)

	// Gets the respond.
	var returnedText map[string]string
	output := recorder.Body.Bytes()
	json.Unmarshal(output, &returnedText)

	// Returns the text.
	return returnedText["fulfillmentText"]
}

// Function to test the trending gif.
func testTrendingGif() string {
	// Creates a request that is passed to the handler.
	request, _ := http.NewRequest("POST", "/dialogflow", nil)

	// Creates the temporary recorder.
	recorder := httptest.NewRecorder()
	router := mux.NewRouter()

	// Sends the request internally.
	router.HandleFunc("/dialogflow", handlers.GetTrendingGif).Methods("POST")
	router.ServeHTTP(recorder, request)

	// Gets the respond.
	var returnedText map[string]string
	output := recorder.Body.Bytes()
	json.Unmarshal(output, &returnedText)

	// Returns the text.
	return returnedText["fulfillmentText"]
}

// Function to test the show dog from animal shelter.
func testShowDog() string {
	// Creates a request that is passed to the handler.
	request, _ := http.NewRequest("POST", "/dialogflow", nil)

	// Creates the temporary recorder.
	recorder := httptest.NewRecorder()
	router := mux.NewRouter()

	// Sends the request internally.
	router.HandleFunc("/dialogflow", handlers.ShowDog).Methods("POST")
	router.ServeHTTP(recorder, request)

	// Gets the respond.
	var returnedText map[string]string
	output := recorder.Body.Bytes()
	json.Unmarshal(output, &returnedText)

	// Returns the text.
	return returnedText["fulfillmentText"]
}

// Function to test the add dog from animal shelter.
func testAddDog() string {
	// Creates a request that is passed to the handler.
	request, _ := http.NewRequest("POST", "/dialogflow", nil)

	// Creates the temporary recorder.
	recorder := httptest.NewRecorder()
	router := mux.NewRouter()

	// Sends the request internally.
	router.HandleFunc("/dialogflow", handlers.AddDog).Methods("POST")
	router.ServeHTTP(recorder, request)

	// Gets the respond.
	var returnedText map[string]string
	output := recorder.Body.Bytes()
	json.Unmarshal(output, &returnedText)

	// Splits the text to grab the link.
	splitText := strings.Split(returnedText["fulfillmentText"], " ")

	// Returns the text.
	return splitText[3]
}

// Function to test the addopt dog from animal shelter.
func testAddoptDog() string {
	// Creates a request that is passed to the handler.
	request, _ := http.NewRequest("POST", "/dialogflow", nil)

	// Creates the temporary recorder.
	recorder := httptest.NewRecorder()
	router := mux.NewRouter()

	// Sends the request internally.
	router.HandleFunc("/dialogflow", handlers.AdoptDog).Methods("POST")
	router.ServeHTTP(recorder, request)

	// Gets the respond.
	var returnedText map[string]string
	output := recorder.Body.Bytes()
	json.Unmarshal(output, &returnedText)

	// Splits the text to grab the link.
	splitText := strings.Split(returnedText["fulfillmentText"], " ")

	// Returns the text.
	return splitText[6]
}

// Function to test the how many dog from animal shelter.
func testHowManyDogs() string {
	// Creates a request that is passed to the handler.
	request, _ := http.NewRequest("POST", "/dialogflow", nil)

	// Creates the temporary recorder.
	recorder := httptest.NewRecorder()
	router := mux.NewRouter()

	// Sends the request internally.
	router.HandleFunc("/dialogflow", handlers.GetCount).Methods("POST")
	router.ServeHTTP(recorder, request)

	// Gets the respond.
	var returnedText map[string]string
	output := recorder.Body.Bytes()
	json.Unmarshal(output, &returnedText)

	// Returns the text.
	return returnedText["fulfillmentText"]
}

// APIContent sends the correct parameters to the html template.
func APIContent(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path

	// Splits the URL path and stores it.
	input := strings.Split(message, "/")

	// Change the text based on the path.
	switch input[2] {
	case "chuckNorrisJoke":
		PageVars.Text = testChuckNorrisJoke()
		displayText(w, r)

	case "dadJoke":
		PageVars.Text = testDadJoke()
		displayText(w, r)

	case "catGif":
		PageVars.Text = testCatGif()
		displayImage(w, r)

	case "dogGif":
		PageVars.Text = testDogGif()
		displayImage(w, r)

	case "hackerGif":
		PageVars.Text = testHackerGif()
		displayImage(w, r)

	case "trendingGif":
		PageVars.Text = testTrendingGif()
		displayImage(w, r)

	case "showDog":
		PageVars.Text = testShowDog()
		displayImage(w, r)

	case "addDog":
		PageVars.Text = testAddDog()
		displayImage(w, r)

	case "addoptDog":
		PageVars.Text = testAddoptDog()
		displayImage(w, r)

	case "howManyDogs":
		PageVars.Text = testHowManyDogs()
		displayText(w, r)

	default:
		return
	}
}

// Function to display text.
func displayText(w http.ResponseWriter, r *http.Request) {
	// Parses the HTML template.
	html, err := template.ParseFiles("website/templates/displayText.html")
	if err != nil {
		log.Print("template parsing error: ", err)
		return
	}

	// Executes the template and sends the parameter to the HTML file.
	err = html.Execute(w, PageVars)
	if err != nil {
		log.Print("template executing error: ", err)
		return
	}
}

// Function to display images.
func displayImage(w http.ResponseWriter, r *http.Request) {
	// Parses the HTML template.
	html, err := template.ParseFiles("website/templates/displayImage.html")
	if err != nil {
		log.Print("template parsing error: ", err)
		return
	}

	// Executes the template and sends the parameter to the HTML file.
	err = html.Execute(w, PageVars)
	if err != nil {
		log.Print("template executing error: ", err)
		return
	}
}
