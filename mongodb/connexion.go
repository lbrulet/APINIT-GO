package mongodb

import (
	"golang.org/x/crypto/bcrypt"
	mgo "gopkg.in/mgo.v2"

	"github.com/lbrulet/APINIT-GO/configs"
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
	Database.Config(configs.Config.DatabaseHost, configs.Config.DatabaseName, configs.Config.DatabaseCollection)
	Database.Connect()
	if hash, err := bcrypt.GenerateFromPassword([]byte("admin123"), 10); err == nil {
		Database.FindOrCreate("admin", hash, "admin@apinit-go.eu", true)
	}
	return &Database
}
