package handlers

import (
	"fmt"
	"net/http"
)

//teller used to count nr. of times a API has been used, Warning must be initialized before use
var teller map[string]int // to be replaced with database

// intitialites the couter map
func init() {
	teller= map[string]int{}
}

// Stats displays the count usage for each type of API request
func Stats(w http.ResponseWriter, r *http.Request) {

  	// explains this sites functionality
	fmt.Fprintln(w, "get counter for /stat ")

	// print statistic for all the different API calls
	for index, valu := range teller {
		fmt.Fprintln(w, index, valu)
	}
	// if no API call has been called display message
	if len(teller) == 0 {
		fmt.Fprintln(w, "no API has been used ")
 	}

}


// Statistic counts the nr of times a given API has been called
func Statistic(str string)  {
	// if it isn't already registered, the create the entry, and set it to 1
	_, ok := teller[str]
	if !ok {
		teller[str] = 1
	// if it already is registered, add 1 to its count
	} else {
		teller[str] += 1
	}

}
