package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/isberg1/IMT2681-Assignment-3/database"
	"github.com/isberg1/IMT2681-Assignment-3/handlers"
	"github.com/isberg1/IMT2681-Assignment-3/website"
)

func main() {
	// Connects to the mongoDB, for the moment hosted in mLabs
	database.Connect()
	r := mux.NewRouter()

	// Allow static files (css and pictures) for the website.
	r.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./website"))))
	// Serves the testing website.
	r.Handle("/website.html", http.FileServer(http.Dir("./website/templates")))
	// Handles the display of content in the website.
	r.HandleFunc("/website.html/", website.APIContent).Methods("GET")

	r.HandleFunc("/", handlers.Frontpage).Methods("GET")
	r.HandleFunc("/dialogflow", handlers.Dialogflow).Methods("POST")
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
