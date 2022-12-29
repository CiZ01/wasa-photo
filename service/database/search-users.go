package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var query_GETUSERS_FOLLOW = `SELECT userID, username, userPropicURL 
							FROM User WHERE userID IN (SELECT followedID FROM Follow WHERE followerID = ?) 
							AND username REGEXP ? ORDER BY username ASC LIMIT ?, ?`
var query_GETUSERS = `SELECT userID, username, userPropicURL FROM User WHERE username REGEXP ? ORDER BY username ASC LIMIT ?, ? `

func (db *appdbimpl) SearchUsers(userID uint32, search string, from_follow bool, offset uint32, limit int32) ([]User, error) {
	var users []User

	var rows *sql.Rows
	var err error

	if from_follow {
		rows, err = db.c.Query(query_GETUSERS_FOLLOW, userID, search+"*", offset, limit)
		if err != nil {
			return nil, err
		}
	} else {
		rows, err = db.c.Query(query_GETUSERS, userID, search+"*", offset, limit)
		if err != nil {
			return nil, err
		}
	}

	for rows.Next() {
		if rows.Err() != nil {
			return nil, err
		}
		var u User
		if err := rows.Scan(&u.UserID, &u.Username, &u.UserPropicURL); err != nil {
			return nil, err
		}
		fmt.Printf("User: %v", u)
		users = append(users, u)
	}
	defer func() { err = rows.Close() }()

	return users, err
}
