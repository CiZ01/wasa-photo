package database

import (
	"database/sql"
	"fmt"
	"os"
)

var query_ADDUSER = `INSERT INTO User (userID, username, bio, userPropicURL)
					 VALUES (?, ?, ?, ?);`
var query_MAXID = `SELECT MAX(userID) FROM User`

func (db *appdbimpl) CreateUser(u User) (User, error) {
	var user User
	user.Username = u.Username

	// ------SET PROPIC URL---------//
	user.UserPropicURL = "./default.png"

	// ------FIND USERID---------//
	var _maxID sql.NullInt32 = sql.NullInt32{Int32: 0, Valid: false}
	row, err := db.c.Query(query_MAXID)
	if err != nil {
		return user, err
	}

	var maxID uint32
	for row.Next() {
		err = row.Scan(&_maxID)
		if err != nil && err != sql.ErrNoRows {
			return user, err
		}

		if !_maxID.Valid {
			maxID = 0
		} else {
			maxID = uint32(_maxID.Int32)
		}
	}

	// ------------INSERT USER--------------//
	bio := ""
	_, err = db.c.Exec(query_ADDUSER, maxID+1, user.Username, bio, user.UserPropicURL)
	if err != nil {
		return user, err
	}

	// --------SET USERID------------//
	user.UserID = maxID + 1

	// --------CREATE USER FOLDER------------//
	path := "./storage/" + fmt.Sprint(user.UserID) + "/posts"
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		return user, err
	}

	return user, nil
}
