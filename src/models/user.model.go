package models

import (
	"database/sql"
	"errors"
)

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
	ID         int        `json:"id" bson:"_id"`
	Username   string     `json:"username" bson:"username" binding:"required"`
	Email      string     `json:"email" bson:"email" binding:"required"`
	Password   string     `json:"-" bson:"password" minLen:"8"`
	Admin      bool       `json:"admin" bson:"admin"`
	Verified   bool       `json:"verified" bson:"verified"`
	AuthMethod AuthMethod `json:"-" bson:"auth_method" binding:"required"`
	_exists    bool
	_deleted   bool
}

// Exists determines if the User exists in the database.
func (u *User) Exists() bool {
	return u._exists
}

// Deleted provides information if the User has been deleted from the database.
func (u *User) Deleted() bool {
	return u._deleted
}

// Insert inserts the User to the database.
func (u *User) Insert(db *sql.DB) error {
	var err error

	// if already exist, bail
	if u._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO apinit_go.users (` +
		`username, email, admin, verified, auth_method, password` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	res, err := db.Exec(sqlstr, u.Username, u.Email, u.Admin, u.Verified, u.AuthMethod, u.Password)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	u.ID = int(id)
	u._exists = true

	return nil
}

// Update updates the User in the database.
func (u *User) Update(db *sql.DB) error {
	var err error

	// sql query
	const sqlstr = `UPDATE apinit_go.users SET ` +
		`username = ?, email = ?, admin = ?, verified = ?, auth_method = ?, password = ?` +
		` WHERE id = ?`

	// run query
	_, err = db.Exec(sqlstr, u.Username, u.Email, u.Admin, u.Verified, u.AuthMethod, u.Password, u.ID)
	return err
}

// Save saves the User to the database.
func (u *User) Save(db *sql.DB) error {
	if u.Exists() {
		return u.Update(db)
	}

	return u.Insert(db)
}

// Delete deletes the User from the database.
func (u *User) Delete(db *sql.DB) error {
	var err error

	// sql query
	const sqlstr = `DELETE FROM apinit_go.users WHERE id = ?`

	// run query
	_, err = db.Exec(sqlstr, u.ID)
	if err != nil {
		return err
	}

	// set deleted
	u._deleted = true

	return nil
}

// UserUpdate is a struct used to update a user
type UserUpdate struct {
	Username string `json:"username" exemple:"sankamille"`
	Email    string `json:"email" exemple:""`
	Password string `json:"password" exemple:"test"`
	Admin    bool   `json:"admin"`
	Verified bool   `json:"verified"`
}
