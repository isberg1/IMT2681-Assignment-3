package database

import "gopkg.in/mgo.v2/bson"

type AnimalShelter struct {
	Address  string
	Database string
	Username string
	Password string
}

type Dog struct {
	ID      bson.ObjectId `bson:"_id"`
	Picture string        `bson:"picture" json:"picture"`
}

type Statistics struct {
	ID        bson.ObjectId `bson:"_id"`
	Timestamp int64         `bson:"timestamp json:"timestamp`
	Text      string        `bson:"text"	json"text"`
}
