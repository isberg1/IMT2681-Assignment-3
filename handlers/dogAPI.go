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

func ShowDog(w http.ResponseWriter, r *http.Request) {
	type dogs struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}

	resp, err := http.Get("https://dog.ceo/api/breeds/image/random")
	if err != nil {
		fmt.Printf("Got no dog")
		panic(err)
	}

	//fmt.Println(resp.Body)
	defer resp.Body.Close()

	dog := &dogs{}
	err = json.NewDecoder(resp.Body).Decode(dog)
	if err != nil {
		fmt.Println("Error decoding json")
		panic(err)
	}
	postToDialogflow(w, dog.Message)
}

func AddDog(w http.ResponseWriter, r *http.Request) {

	type dogs struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}

	resp, err := http.Get("https://dog.ceo/api/breeds/image/random")
	if err != nil {
		fmt.Printf("Got no dog")
		panic(err)
	}

	//fmt.Println(resp.Body)
	defer resp.Body.Close()

	dog := &dogs{}
	err = json.NewDecoder(resp.Body).Decode(dog)
	if err != nil {
		fmt.Println("Error decoding json")
		logging(err.Error())
	}

	dogg := database.Dog{
		ID:      bson.NewObjectId(),
		Picture: dog.Message,
	}

	if err := database.Insert(dogg); err != nil {
		fmt.Println("Error inserting new dog")
		panic(err)
	}

	returnString := strings.Join([]string{
		"New dog added ", dog.Message}, "")
	postToDialogflow(w, returnString)
}

func AdoptDog(w http.ResponseWriter, r *http.Request) {
	dog, err := database.FindOldestDog()
	if err != nil {
		fmt.Println("Could not find latest")
		panic(err)
	}
	dog, err = database.DeleteDogWithId(dog.ID.Hex())
	if err != nil {
		fmt.Println("Coult not delete dog")
		panic(err)
	}

	returnString := strings.Join([]string{
		"Congratulations, you adopted a new dog ", dog.Picture}, "")
	postToDialogflow(w, returnString)
}

func getCount(w http.ResponseWriter, r *http.Request) {
	dogCount, err := database.FindCount("dogs")
	if err != nil {
		postToDialogflow(w, "No Dogs in shelter")
	}
	returnString := strings.Join([]string{
		"There are ", strconv.Itoa(dogCount),
		" Dogs currently in the shelter"},
		"")
	postToDialogflow(w, returnString)
}

func ShowAllDogs(w http.ResponseWriter, r *http.Request) {
	dogs, err := database.FindAll()
	if err != nil {
		postToDialogflow(w, "No dogs in shelter")
	}

	for i := 0; i < len(dogs); i++ {
		postToDialogflow(w, dogs[i].Picture)
	}

}
