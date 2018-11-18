package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/isberg1/IMT2681-Assignment-3/database"
	"gopkg.in/mgo.v2/bson"
)

// ShowDog shows a picture of a random dog
func ShowDog(w http.ResponseWriter, r *http.Request) {
	type dogs struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}

	// Sends a request to the Dop API
	resp, err := http.Get("https://dog.ceo/api/breeds/image/random")
	if err != nil {
		fmt.Printf("Got no dog")
		logging(err.Error())
	}

	//fmt.Println(resp.Body)
	defer resp.Body.Close()

	// Decodes the response into the struct
	dog := &dogs{}
	err = json.NewDecoder(resp.Body).Decode(dog)
	if err != nil {
		fmt.Println("Error decoding json")
		logging(err.Error())
	}
	// Send the resoponse to dialogflow
	postToDialogflow(w, dog.Message)
}

// AddDog adds a dog to the database
func AddDog(w http.ResponseWriter, r *http.Request) {

	// Struct made to decode the values into
	type dogs struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}

	// Send a request to get values from dog API as json
	resp, err := http.Get("https://dog.ceo/api/breeds/image/random")
	if err != nil {
		fmt.Printf("Got no dog")
		logging(err.Error())
	}

	//fmt.Println(resp.Body)
	defer resp.Body.Close()

	// Decodes the values into the struct
	dog := &dogs{}
	err = json.NewDecoder(resp.Body).Decode(dog)
	if err != nil {
		fmt.Println("Error decoding json")
		logging(err.Error())
	}

	// Created the object to insert into the database
	dogg := database.Dog{
		ID:      bson.NewObjectId(),
		Picture: dog.Message,
	}

	// Tries to insert a new object into the database
	if err := database.Insert(dogg); err != nil {
		fmt.Println("Error inserting new dog")
		logging(err.Error())
	}

	// Creates a string, and sends it to dialogflow as a response to the user
	returnString := strings.Join([]string{
		"New dog added ", dog.Message}, "")
	postToDialogflow(w, returnString)
}

// AdoptDog removes a dog object from the database,
// and gives the user a response that they adopted a dog
func AdoptDog(w http.ResponseWriter, r *http.Request) {
	// the database is a FIFO list,
	// So the oldest dog gets adopted first
	dog, err := database.FindOldestDog()
	if err != nil {
		fmt.Println("Could not find latest")
		logging(err.Error())
		return
	}
	// Removes the dog object from the database based on ID
	dog, err = database.DeleteDogWithID(dog.ID.Hex())
	if err != nil {
		fmt.Println("Coult not delete dog")
		logging(err.Error())
		return
	}

	// Returns a picture of the dog deleted from the database
	returnString := strings.Join([]string{
		"Congratulations, you adopted a new dog ", dog.Picture}, "")
	postToDialogflow(w, returnString)
}

// GetCount returns a count of how many dogs are currently in the database
func GetCount(w http.ResponseWriter, r *http.Request) {
	// Returns a count of how many objects are in the DB
	dogCount, err := database.FindCount("dogs")
	if err != nil {
		postToDialogflow(w, "No Dogs in shelter")
	}
	// Makes the message to send to dialogFlow
	returnString := strings.Join([]string{
		"There are ", strconv.Itoa(dogCount),
		" Dogs currently in the shelter"},
		"")
	postToDialogflow(w, returnString)
}

// ShowAllDogs Returns a picture of all dogs currently in the database
func ShowAllDogs(w http.ResponseWriter, r *http.Request) {
	// Returns all object in the db
	dogs, err := database.FindAll()
	if err != nil {
		postToDialogflow(w, "No dogs in shelter")
	}

	var returnString string

	// Adds the url to all the pictures from the object into a string separated by newline
	for i := 0; i < len(dogs); i++ {
		returnString = strings.Join([]string{returnString, "\n", dogs[i].Picture}, "")
	}

	// sends it to dialogflow
	postToDialogflow(w, returnString)

}
