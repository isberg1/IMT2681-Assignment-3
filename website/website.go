package website

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/mux"
	"github.com/isberg1/IMT2681-Assignment-3/handlers"
)

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
	var joke map[string]string
	output := recorder.Body.Bytes()
	json.Unmarshal(output, &joke)

	// Returns the joke.
	return joke["fulfillmentText"]
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
	var joke map[string]string
	output := recorder.Body.Bytes()
	json.Unmarshal(output, &joke)

	// Returns the joke.
	return joke["fulfillmentText"]
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
	var joke map[string]string
	output := recorder.Body.Bytes()
	json.Unmarshal(output, &joke)

	// Returns the joke.
	return joke["fulfillmentText"]
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
	var joke map[string]string
	output := recorder.Body.Bytes()
	json.Unmarshal(output, &joke)

	// Returns the joke.
	return joke["fulfillmentText"]
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
	var joke map[string]string
	output := recorder.Body.Bytes()
	json.Unmarshal(output, &joke)

	// Returns the joke.
	return joke["fulfillmentText"]
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
	var joke map[string]string
	output := recorder.Body.Bytes()
	json.Unmarshal(output, &joke)

	// Returns the joke.
	return joke["fulfillmentText"]
}

// TEMP
func TestingWebsite(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, testChuckNorrisJoke()+"\n")
	fmt.Fprintf(w, testDadJoke()+"\n")
	fmt.Fprintf(w, testCatGif()+"\n")
	fmt.Fprintf(w, testDogGif()+"\n")
	fmt.Fprintf(w, testHackerGif()+"\n")
	fmt.Fprintf(w, testTrendingGif()+"\n")
}
