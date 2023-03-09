package database

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
)

var query_ADDUSER = `INSERT INTO User (userID, username, bio)
					 VALUES (?, ?, ?);`
var query_MAXID = `SELECT MAX(userID) FROM User`

func (db *appdbimpl) CreateUser(u User) (User, error) {
	var user User
	user.Username = u.Username

	// ------FIND USERID---------//
	var _maxID = sql.NullInt64{Int64: 0, Valid: false}
	row, err := db.c.Query(query_MAXID)
	if err != nil {
		return user, err
	}

	var maxID int
	for row.Next() {
		if row.Err() != nil {
			return user, err
		}

		err = row.Scan(&_maxID)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return user, err
		}

		if !_maxID.Valid {
			maxID = 0
		} else {
			maxID = int(_maxID.Int64)
		}
	}

	// --------SET USERID------------//
	user.UserID = maxID + 1

	// --------CREATE USER FOLDER------------//
	path := "./storage/" + fmt.Sprint(user.UserID) + "/posts"
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return user, err
	}

	// ------------INSERT USER--------------//
	bio := ""
	_, err = db.c.Exec(query_ADDUSER, user.UserID, user.Username, bio)
	if err != nil {
		return user, err
	}

	return user, nil
}
