package models

import "gopkg.in/mgo.v2/bson"

// AuthMethod is a authentication mark
type AuthMethod int

const (
	// LOCAL authentication
	LOCAL AuthMethod = 0
	// GOOGLE authentication
	GOOGLE AuthMethod = 1
	// FACEBOOK authentication
	FACEBOOK AuthMethod = 2
)

// User is a struct about a user
type User struct {
	ID         bson.ObjectId `json:"id" bson:"_id"`
	Username   string        `json:"username" bson:"username" binding:"required"`
	Email      string        `json:"email" bson:"email" binding:"required"`
	Password   string        `json:"-" bson:"password" minLen:"8"`
	Admin      bool          `json:"admin" bson:"admin"`
	Verified   bool          `json:"verified" bson:"verified"`
	AuthMethod AuthMethod    `json:"-" bson:"auth_method" binding:"required"`
}
