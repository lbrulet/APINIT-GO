package database

import (
	"github.com/lbrulet/APINIT-GO/src/models"
)

func (db *databaseManager) CountUserByKey(key, data string) (int, error) {
	var count int

	row := db.DB.QueryRow(`SELECT COUNT(*) FROM users WHERE `+key+" = ?;", data)
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (db *databaseManager) FindUserByKey(key, data string) (*models.User, error) {
	var err error

	var sqlstr = `SELECT ` +
		`id, username, email, admin, verified, auth_method, password ` +
		`FROM apinit_go.users ` +
		`WHERE ` + key + ` = ?`

	u := models.User{}

	err = db.DB.QueryRow(sqlstr, data).Scan(&u.ID, &u.Username, &u.Email, &u.Admin, &u.Verified, &u.AuthMethod, &u.Password)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (db *databaseManager) GetAllUsers() ([]*models.User, error) {
	users := make([]*models.User, 0)
	rows, err := db.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := new(models.User)
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Admin, &user.Verified, &user.AuthMethod, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (db *databaseManager) GetUserByID(id int) (*models.User, error) {
	var err error

	const sqlstr = `SELECT ` +
		`id, username, email, admin, verified, auth_method, password ` +
		`FROM apinit_go.users ` +
		`WHERE id = ?`

	u := models.User{}

	err = db.DB.QueryRow(sqlstr, id).Scan(&u.ID, &u.Username, &u.Email, &u.Admin, &u.Verified, &u.AuthMethod, &u.Password)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
