package database

import (
	"database/sql"
	"errors"
	"fmt"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"

	"github.com/lbrulet/APINIT-GO/src/configs"
	"github.com/lbrulet/APINIT-GO/src/models"
)

// DBManager variable that store the db instance
var DBManager models.DatabaseManager

// InitDB connect the api to mysql
func InitDB() error {
	var err error
	if len(configs.Config.DatabasePassword) == 0 {
		return errors.New("DATABASE_PASSWORD is missing")
	}
	dnsStr := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		configs.Config.DatabaseUser, configs.Config.DatabasePassword, configs.Config.DatabaseEndPoint, configs.Config.DatabaseName,
	)
	DBManager.DB, err = sql.Open("mysql", dnsStr)
	if err != nil {
		return err
	}
	err = DBManager.DB.Ping()
	if err != nil {
		return err
	}
	fmt.Printf("[DB] Connected to: %s\n", configs.Config.DatabaseEndPoint)
	return nil
}
