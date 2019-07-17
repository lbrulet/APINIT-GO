package mongodb

import (
	"fmt"
	"log"

	"github.com/lbrulet/APINIT-GO/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (m *DatabaseService) Connect() {
	session, err := mgo.Dial(m.host)
	if err != nil {
		log.Fatal(err)
	}
	m.db = session.DB(m.database)
	fmt.Printf("Connected to the database\n")
}

func (m *DatabaseService) Config(host, database, collection string) {
	m.host = host
	m.database = database
	m.collection = collection
}

func (m *DatabaseService) getCollection() *mgo.Collection {
	return m.db.C(m.collection)
}

func (m *DatabaseService) GetDatabase() *mgo.Session {
	return m.db.Session
}

func (m *DatabaseService) FindAll() ([]models.User, error) {
	var users []models.User
	err := m.getCollection().Find(bson.M{}).All(&users)
	return users, err
}

func (m *DatabaseService) FindByID(id string) (models.User, error) {
	var user models.User
	err := m.getCollection().Find(bson.M{"_id": id}).One(&user)
	return user, err
}

func (m *DatabaseService) FindOrCreate(id string, username string, service string) error {
	m.getCollection().Insert()
	return nil
}

func (m *DatabaseService) FindOrCreateAdmin(username string, password string) error {
	return nil
}

func (m *DatabaseService) FindByKey(key, data string) (models.User, error) {
	var user models.User
	err := m.getCollection().Find(bson.M{key: data}).One(&user)
	return user, err
}

func (m *DatabaseService) FindByUsername(username string) (models.User, error) {
	var user models.User
	err := m.getCollection().Find(bson.M{"username": username}).One(&user)
	return user, err
}

func (m *DatabaseService) Insert(user models.User) error {
	err := m.getCollection().Insert(&user)
	return err
}

func (m *DatabaseService) Delete(user models.User) error {
	err := m.getCollection().Remove(&user)
	return err
}

func (m *DatabaseService) Update(user models.User) error {
	err := m.getCollection().UpdateId(user.ID, &user)
	return err
}
