package database

import (
	"database/sql"
	"errors"
	"fmt"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"

	"github.com/lbrulet/APINIT-GO/src/configs"
)

type databaseManager struct {
	DB *sql.DB
}

// Database variable that store the db instance
var Database databaseManager

// InitDB connect the api to mysql
func InitDB() error {
	var err error
	if len(configs.Config.DatabasePassword) == 0 {
		return errors.New("DATABASE_PASSWORD is missing")
	}
	dnsStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?multiStatements=true",
		configs.Config.DatabaseUser, configs.Config.DatabasePassword, configs.Config.DatabaseEndPoint, configs.Config.DatabaseName,
	)
	Database.DB, err = sql.Open("mysql", dnsStr)
	if err != nil {
		return err
	}
	err = Database.DB.Ping()
	if err != nil {
		return err
	}
	fmt.Printf("[DB] Connected to: %s | %s\n", configs.Config.DatabaseEndPoint, configs.Config.DatabaseName)
	return nil
}
