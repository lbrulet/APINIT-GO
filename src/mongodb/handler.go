package mongodb

import (
	"errors"
	"fmt"
	"log"

	"github.com/lbrulet/APINIT-GO/src/models"
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
	if valid := bson.IsObjectIdHex(id); valid == false {
		return models.User{}, errors.New("Invalid id")
	}
	err := m.getCollection().Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&user)
	return user, err
}

// FindOrCreate create a new admin user if he does not exist
func (m *DatabaseService) FindOrCreate(username string, password []byte, email string, isAdmin bool) error {
	var user models.User
	err := m.getCollection().Find(bson.M{"username": username}).One(&user)
	if err != nil {
		user.ID = bson.NewObjectId()
		user.Username = username
		user.Password = password
		user.Email = email
		user.AuthMethod = models.LOCAL
		user.Verified = true
		if isAdmin {
			user.Admin = true
		} else {
			user.Admin = false
		}
		err := m.getCollection().Insert(&user)
		return err
	}
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
