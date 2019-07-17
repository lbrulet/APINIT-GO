package mongodb

import (
	mgo "gopkg.in/mgo.v2"
)

type DatabaseService struct {
	host       string
	database   string
	collection string
	db         *mgo.Database
}

var (
	// Users is a struct
	Database DatabaseService
)

func Connect() *DatabaseService {
	Database.Config("localhost", "ROLLEAT", "users")
	Database.Connect()
	return &Database
}
