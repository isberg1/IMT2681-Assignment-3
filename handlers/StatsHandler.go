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

// displaces the count usage for each type of API request
func Stats(w http.ResponseWriter, r *http.Request) {


	fmt.Fprintln(w, "get counter for /stat ")

	//fmt.Fprintln(w, len(teller))

		if len(teller) > 0 {

			for index, valu := range teller {
				fmt.Fprintln(w, index, valu)
			}
		}

		if len(teller) == 0 {
			fmt.Fprintln(w, "no content")

		} else {
			fmt.Fprintln(w, "some content")

		}


}


// Statistic counts the nr of times a givven API has been called
func Statistic(str string)  {

	_, ok := teller[str]
	if !ok {
		teller[str] = 0
	}


	switch str {
	case "joke":
		teller[str] += 1

	case "dad":
		teller[str] += 1

	default:
		logging("function called with incorrect string")
	}

	if len(teller) < 0 {
		fmt.Print(teller[str])
	}

}
