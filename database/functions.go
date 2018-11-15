package database

import (
	"fmt"
	"log"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (m *AnimalShelter) Connect() {
	session := &mgo.DialInfo{
		Addrs:    []string{m.Address},
		Timeout:  60 * time.Second,
		Database: m.Database,
		Username: m.Username,
		Password: m.Password,
	}

	connection, err := mgo.DialWithInfo(session)
	if err != nil {
		fmt.Println("Could not connect to DB")
		log.Fatal(err)
	}
	db = connection.DB(m.Database)
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
