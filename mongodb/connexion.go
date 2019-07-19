package mongodb

import (
	mgo "gopkg.in/mgo.v2"
)

// DatabaseService is a struct that containt all informations about the database
type DatabaseService struct {
	host       string
	database   string
	collection string
	db         *mgo.Database
}

var (
	// Database is the entier service
	Database DatabaseService
)

// Connect is used to config & connect the api to the database
func Connect() *DatabaseService {
	Database.Config("localhost", "ROLLEAT", "users")
	Database.Connect()
	return &Database
}
