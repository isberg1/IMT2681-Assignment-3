package handlers

import (
	"fmt"
	"net/http"
)

// prints everything that has previously been posted  to "/dialogflow (used for input analyse and debugging)
var postMemory []string
func OldPosts(w http.ResponseWriter, r *http.Request) {

	for _,val := range postMemory{
		fmt.Fprintln(w, val)
	}

}

