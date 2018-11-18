package database

import "gopkg.in/mgo.v2/bson"

// AnimalShelter stores data about the DB
type AnimalShelter struct {
	Address  string
	Database string
	Username string
	Password string
}

// Dog is a model for the dog document in the DB
type Dog struct {
	ID      bson.ObjectId `bson:"_id"`
	Picture string        `bson:"picture" json:"picture"`
}

// Statistics is a model for the stat document in the DB
type Statistics struct {
	ID        bson.ObjectId `bson:"_id"`
	Timestamp int64         `bson:"timestamp" json:"timestamp"`
	Command   string        `bson:"command" json:"command"`
	Visitors  string        `bson:"visitors" json:"visitors"`
}
