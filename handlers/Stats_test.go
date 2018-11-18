package handlers

import (
	"github.com/gorilla/mux"
	"github.com/isberg1/IMT2681-Assignment-3/database"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func Test_statdb(t *testing.T) {

	database.Connect()
}

func Test_statHandler(t *testing.T)  {


	// Creates a request that is passed to the handler.
	request, err2 := http.NewRequest("GET", "/statistics", nil)
	if err2 != nil {
		t.Error("failed http.NewRequest", err2)
	}
	// Creates the recorder and router.
	recorder := httptest.NewRecorder()
	router := mux.NewRouter()

	// Tests the function.
	router.HandleFunc("/statistics", Stats).Methods("GET")
	router.ServeHTTP(recorder, request)

	// Check the status code is what we expect (200).
	status := recorder.Code
	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check if returned value is empty.
	actual := recorder.Body.String()
	if actual == "" {
		t.Error("handler returned no data")
	}

}
// checks that a statistic entry is updated when a command is used
func Test_addStatisticEntry(t *testing.T) {
	// type of entry to check
	commandType := "joke"
	// ensure that i entry exits
	Statistic(commandType)

	// count nr of existing entries
	jokeStat, err := database.GetStatObject(commandType)
	if err != nil {
		t.Error("fail, can not retrieve entry for",commandType)
	}
	// convert to int
	firstCount, err2 := strconv.Atoi(jokeStat.Visitors)
	if err2 != nil {
		t.Error("fail, unable to convert visitor count to int ", err2)
	}

	// add +1 to entry type
	Statistic(commandType)
	// count nr of existing entries
	jokeStat2, err3 := database.GetStatObject(commandType)
	if err3 != nil {
		t.Error("fail, can not retrieve entry for",commandType, err3)
	}
	// convert to int
	secondCount, err4 := strconv.Atoi(jokeStat2.Visitors)
	if err4 != nil {
		t.Error("fail, unable to convert visitor count to int ",err4)
	}
	// check if nr of count is correct
	if secondCount != ( firstCount+1) {
		t.Error("fail, wrong visitor count ")
	}

}

