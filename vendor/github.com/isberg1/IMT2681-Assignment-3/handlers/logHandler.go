package handlers

import (
	"fmt"
	"net/http"
)

// the logging entries can possibly be moved to a database



// stores all log messages
var logArray []string

// displays all log messages
func Log(w http.ResponseWriter, r *http.Request) {

	for _, val:= range logArray{
		fmt.Fprintln(w, val)
	}

}
// adds new log messages to storage
func logging(s string)  {
	logArray = append(logArray, s)
}

