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
)

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

func (m *AnimalShelter) Insert(dog Dog) error {
	err := db.C(COLLECTION).Insert(&dog)
	return err
}

func (m *AnimalShelter) FindOldestDog() (Dog, error) {
	var dog Dog
	err := db.C(COLLECTION).Find(nil).Sort("_id").One(&dog)
	return dog, err
}

func (m *AnimalShelter) DeleteDogWithId(id string) (Dog, error) {
	var dog Dog
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&dog)
	err = db.C(COLLECTION).RemoveId(bson.ObjectIdHex(id))
	return dog, err
}

func (m *AnimalShelter) FindCount() (int, error) {
	trackCount, err := db.C(COLLECTION).Count()
	return trackCount, err
}

func (m *AnimalShelter) FindAll() ([]Dog, error) {
	fmt.Println("Trying to find all")
	var dogs []Dog
	// Using the nil parameter in find gets all tracks
	err := db.C(COLLECTION).Find(nil).All(&dogs)
	return dogs, err
}

func (m *AnimalShelter) DeleteAll() (*mgo.ChangeInfo, error) {
	rem, err := db.C(COLLECTION).RemoveAll(nil)
	return rem, err
}
