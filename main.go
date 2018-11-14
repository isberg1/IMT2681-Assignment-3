package main

import (
	"github.com/gorilla/mux"
	"github.com/isberg1/IMT2681-Assignment-3/handlers"
	"net/http"
	"os"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/", handlers.Frontpage).Methods("GET")
	r.HandleFunc("/dialogflow", handlers.Dialogflow).Methods("POST")
	r.HandleFunc("/OldPosts", handlers.OldPosts).Methods("GET")
	r.HandleFunc("/log", handlers.Log).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := http.ListenAndServe(":"+port, r); err != nil {
		panic(err)
	}
}
