package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/isberg1/IMT2681-Assignment-3/database"
	"gopkg.in/mgo.v2/bson"
)

//teller used to count nr. of times a API has been used, Warning must be initialized before use
var teller map[string]int // to be replaced with database

// intitialites the couter map
func init() {
	teller = map[string]int{}
}

// Stats displays the count usage for each type of API request
func Stats(w http.ResponseWriter, r *http.Request) {

	stats, err := database.QueryForStats()
	if err != nil {
		logging(err.Error())
	}

	fmt.Fprintln(w, stats)
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
func Statistic(str string) {
	// DISCLAIMER this is a really bad way of doing this, idealy one should query the
	// database using .Find(bson.M{"value" : value]}.Limit(1). However could not get
	// it to work and therefore had to cheese it.

	//Returns all the objects in the database
	stats, err := database.QueryForStats()
	if err != nil {
		logging(err.Error())
	}

	// Puts the command field of every object into an array
	var statts []string
	for i := 0; i < len(stats); i++ {
		statts = append(statts, (stats[i].Command))
	}

	var exist bool

	// Checks if any on the objects had a command value that corresponded to the sent variable
	for i := 0; i < len(statts); i++ {
		if statts[i] == str {
			exist = true
			break
		}
	}

	// If no object existed, make a new one and insert it
	if !exist {
		fmt.Println("No entires in DB")

		// A struct for creating the object
		stat := database.Statistics{
			ID:        bson.NewObjectId(),
			Timestamp: time.Now().UnixNano() / int64(time.Millisecond),
			Command:   str,
			Visitors:  "1",
		}

		// Inserts the object to the db
		err = database.InsertStatistics(stat)
		if err != nil {
			logging(err.Error())
		}

		// If there is already and existing stat object, take it out and update it's values
	} else {
		// Gets the stat object from the database
		stat, err := database.GetStatObject()
		if err != nil {
			logging(err.Error())
		}
		// Generates a new timestamp to update
		newTimeStamp := time.Now().UnixNano() / int64(time.Millisecond)
		// Gets the visitors value, converts it to int to add 1, then converts it back to string to update
		howMany := stat.Visitors
		HM, err := strconv.Atoi(howMany)
		if err != nil {
			logging(err.Error())
		}
		newVisitors := HM + 1
		newVisitorss := strconv.Itoa(newVisitors)
		// Updates the object in the database
		database.UpdateStats(stat.Command, newTimeStamp, newVisitorss)
	}

}
