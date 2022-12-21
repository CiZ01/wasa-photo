package database

import (
	"fmt"
	"os"
)

var query_ADDUSER = `INSERT INTO User (userID, username, bio, userPropicURL)
					 VALUES (?, ?, ?);`
var query_MAXID = `SELECT MAX(userID) FROM User`

func (db *appdbimpl) CreateUser(u User) (User, error) {
	var user User
	user.Username = u.Username

	// ------SET PROPIC URL---------//
	user.UserPropicURL = "./default.png"

	// ------FIND USERID---------//
	var maxID uint32
	err := db.c.QueryRow(query_MAXID).Scan(&maxID)
	if err != nil {
		return user, err
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
