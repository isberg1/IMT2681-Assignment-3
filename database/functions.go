package database

import (
	"fmt"
	"log"
	"os"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var db *mgo.Database

const (
	COLLECTION = "dogs"
	LOG        = "log"
	STAT       = "statistics"
)

// Connects to the database, gettint the env variables from heroku/other files
func Connect() {
	session := &mgo.DialInfo{
		Addrs:    []string{os.Getenv("MONGO_ADDRESS")},
		Timeout:  60 * time.Second,
		Database: os.Getenv("MONGO_DATABASE"),
		Username: os.Getenv("MONGO_USER"),
		Password: os.Getenv("MONGO_PASSWORD"),
	}

	connection, err := mgo.DialWithInfo(session)
	if err != nil {
		fmt.Println("Could not connect to DB")
		log.Fatal(err)
	}
	db = connection.DB(os.Getenv("MONGO_DATABASE"))
}

// inserts a dog into the db
func Insert(dog Dog) error {
	err := db.C(COLLECTION).Insert(&dog)
	return err
}

// inserts stats into the db
func InsertStatistics(stat Statistics) error {
	err := db.C(STAT).Insert(&stat)
	return err
}

// The stats database is supposed to only contain one object,
//therefore one can assume that no sorting or query method is required.
func GetStatObject() (Statistics, error) {
	var stat Statistics
	err := db.C(STAT).Find(nil).One(&stat)
	return stat, err
}

func QueryForStats() ([]Statistics, error) {
	var stats []Statistics
	// Using the nil parameter in find gets all tracks
	err := db.C(STAT).Find(nil).All(&stats)
	return stats, err
}

//Updates the stat object with the new values sent.
// usin bson.M to find the right object to update based on ID, then to find the right value to update.
func UpdateStats(command string, timestamp int64, visitors string) {
	db.C(STAT).Update(bson.M{"command": command}, bson.M{"$set": bson.M{"timestamp": timestamp}})
	db.C(STAT).Update(bson.M{"command": command}, bson.M{"$set": bson.M{"visitors": visitors}})
}

// Finds the oldest dog by sorting by id and taking the first
func FindOldestDog() (Dog, error) {
	var dog Dog
	err := db.C(COLLECTION).Find(nil).Sort("_id").One(&dog)
	return dog, err
}

// Deletes the dog with the id sent, and returnes the dog object.
func DeleteDogWithId(id string) (Dog, error) {
	var dog Dog
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&dog)
	err = db.C(COLLECTION).RemoveId(bson.ObjectIdHex(id))
	return dog, err
}

// Finds a count of how many object in current collection in the DB
func FindCount(coll string) (int, error) {
	trackCount, err := db.C(coll).Count()
	return trackCount, err
}

// Finds all Dog objects in the database
func FindAll() ([]Dog, error) {
	fmt.Println("Trying to find all")
	var dogs []Dog
	// Using the nil parameter in find gets all tracks
	err := db.C(COLLECTION).Find(nil).All(&dogs)
	return dogs, err
}

// Deletes all objects in a collection, not used yet
func DeleteAll() (*mgo.ChangeInfo, error) {
	rem, err := db.C(COLLECTION).RemoveAll(nil)
	return rem, err
}
