package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/isberg1/IMT2681-Assignment-3/handlers"
)

var AS = AnimalShelter{
	Address:  os.Getenv("MONGO_ADDRESS"),
	Database: os.Getenv("MONGO_DATABASE"),
	Username: os.Getenv("MONGO_USER"),
	Password: os.Getenv("MONGO_PASSWORD"),
}

func main() {
	AS.databse.Connect()
	r := mux.NewRouter()

	r.HandleFunc("/", handlers.Frontpage).Methods("GET")
	r.HandleFunc("Dialogflow", handlers.Dialogflow).Methods("POST")
	r.HandleFunc("/OldPosts", handlers.OldPosts).Methods("GET")
	r.HandleFunc("/log", handlers.Log).Methods("GET")
	r.HandleFunc("/statistics", handlers.Stats).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := http.ListenAndServe(":"+port, r); err != nil {
		panic(err)
	}
}
