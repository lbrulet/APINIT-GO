package mongodb

import (
	"fmt"
	"log"

	"github.com/lbrulet/APINIT-GO/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Connect will etablish a new session
func (m *DatabaseService) Connect() {
	session, err := mgo.Dial(m.host)
	if err != nil {
		log.Fatal(err)
	}
	m.db = session.DB(m.database)
	fmt.Printf("[MONGODB] Connected to the database - %s | %s | %s\n", m.host, m.database, m.collection)
}

// Config is a function that is used to config the database
func (m *DatabaseService) Config(host, database, collection string) {
	m.host = host
	m.database = database
	m.collection = collection
}

func (m *DatabaseService) getCollection() *mgo.Collection {
	return m.db.C(m.collection)
}

// GetDatabase return the current session
func (m *DatabaseService) GetDatabase() *mgo.Session {
	return m.db.Session
}

// FindAll return an array of user
func (m *DatabaseService) FindAll() ([]models.User, error) {
	var users []models.User
	err := m.getCollection().Find(bson.M{}).All(&users)
	return users, err
}

// FindByID return a specific user fund by his ID
func (m *DatabaseService) FindByID(id string) (models.User, error) {
	var user models.User
	err := m.getCollection().Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&user)
	return user, err
}

// FindOrCreate create a new user if he does not exist
func (m *DatabaseService) FindOrCreate(id string, username string, service string) error {
	m.getCollection().Insert()
	return nil
}

// FindOrCreate create a new admin user if he does not exist
func (m *DatabaseService) FindOrCreateAdmin(username string, password string) error {
	return nil
}

// FindByKey return a specific user depending on the key used
func (m *DatabaseService) FindByKey(key, data string) (models.User, error) {
	var user models.User
	err := m.getCollection().Find(bson.M{key: data}).One(&user)
	return user, err
}

// Insert create a new user
func (m *DatabaseService) Insert(user models.User) error {
	err := m.getCollection().Insert(&user)
	return err
}

// Delete remove a user
func (m *DatabaseService) Delete(user models.User) error {
	err := m.getCollection().Remove(&user)
	return err
}

// Update update a user
func (m *DatabaseService) Update(user models.User) error {
	err := m.getCollection().UpdateId(user.ID, &user)
	return err
}
