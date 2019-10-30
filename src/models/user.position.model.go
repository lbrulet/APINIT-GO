package models

import (
	"database/sql"
	"errors"
	"time"
)

// UserPosition is a struct that store the geolocalisation of a user
type UserPosition struct {
	ID        int       `json:"id"`
	Latitude  float32   `json:"latitude"`
	Longitude float32   `json:"longitude"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	_exists   bool
	_deleted  bool
}

// Exists determines if the UserPosition exists in the database.
func (p *UserPosition) Exists() bool {
	return p._exists
}

// Deleted provides information if the UserPosition has been deleted from the database.
func (p *UserPosition) Deleted() bool {
	return p._deleted
}

// Insert inserts the User to the database.
func (p *UserPosition) Insert(db *sql.DB) error {
	var err error

	// if already exist, bail
	if p._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO apinit_go.user_position (` +
		`latitude, longitude, user_id` +
		`) VALUES (` +
		`?, ?, ?` +
		`)`

	// run query
	res, err := db.Exec(sqlstr, p.Latitude, p.Longitude, p.UserID)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	p.ID = int(id)
	p._exists = true

	return nil
}
