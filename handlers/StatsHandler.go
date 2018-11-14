package handlers

import (
	"fmt"
	"net/http"
)

var Counter map[string]int // to be replaced with database




func Stats(w http.ResponseWriter, r *http.Request) {

	for index, valu := range Counter {
		fmt.Fprintln(w, index, valu)
	}
}









func Statistic(str string)  {

	switch str {
	case "chuchNorrisAPI":
		Counter[str] ++

	default:
		logging("function called with incorrect string")

	}
}
