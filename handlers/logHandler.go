package handlers

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

// the logging entries can possibly be moved to a database

// stores all log messages
var logArray string
var logTimestamp string

// Log displays all log messages, used for storing errors, and debugging
func Log(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "text/plain")
	fmt.Fprintln(w, "Errors: ")
	fmt.Fprintln(w, logArray)

}

// adds new log messages to storage
func logging( /* w http.ResponseWriter, r *http.Request,*/ s string) {
	logTimestamp = time.Now().String()
	logArray = strings.Join([]string{
		logTimestamp, "\t", s, "\n"},
		"")
}
