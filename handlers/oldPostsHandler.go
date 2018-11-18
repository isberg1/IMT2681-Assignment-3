package handlers

import (
	"fmt"
	"net/http"
	"strings"
)

// prints everything that has previously been posted  to "/dialogflow (used for input analyse and debugging)
var postMemory string

// OldPosts prints all the preveous posts from dialogflow,
// used for analyzing structure of json input, and debugging
func OldPosts(w http.ResponseWriter, r *http.Request) {

	/*
		for _, val := range postMemory {
			fmt.Fprintln(w, val)
		}
	*/
	var commands string
	commands = strings.Join([]string{
		commands, "\n",
		postMemory, "\n"},
		"")
	fmt.Fprintln(w, postMemory)
	fmt.Fprintln(w, commands)
}
