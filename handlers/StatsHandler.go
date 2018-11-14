package handlers

import (
	"fmt"
	"net/http"
)

//Counter used to count nr. of times a API has been used, Warning must be initialized before use
var teller map[string]int // to be replaced with database

func Stats(w http.ResponseWriter, r *http.Request) {

	// TODO get value form DB instead of Counter

	for index, valu := range teller {
		fmt.Fprintln(w, index, valu)
	}
}

func init() {
	teller= map[string]int{}
}

func Statistic(str string)  {

	switch str {
	case "chuchNorrisAPI":
		teller[str] ++
	case "dad":
		teller[str] ++

	default:
		logging("function called with incorrect string")

	}


}
