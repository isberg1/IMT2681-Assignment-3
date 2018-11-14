package handlers

import (
	"fmt"
	"net/http"
)

// webpage explaining api functionality
func Frontpage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintln(w, "post to /dialogflow ")

}

