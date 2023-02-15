package database

import "database/sql"

var query_FINDUSERNAME = `SELECT username FROM User
						WHERE username = ?`

func (db *appdbimpl) ExistsName(username string) (bool, error) {
	var existsName string
	err := db.c.QueryRow(query_FINDUSERNAME, username).Scan(&existsName)
	if err != nil && err == sql.ErrNoRows {
		return false, nil
	}
	return existsName != "", err
}
