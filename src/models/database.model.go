package models

import "database/sql"

// DatabaseManager struct to handle the database
type DatabaseManager struct {
	DB *sql.DB
}
