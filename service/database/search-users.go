package database

import (
	_ "github.com/mattn/go-sqlite3"
)

var query_GETUSERS = `SELECT userID, username FROM User WHERE username regexp ? ORDER BY username LIMIT ?, ? `

func (db *appdbimpl) SearchUsers(userID int, search string, offset int, limit int) ([]User, error) {
	var users []User

	rows, err := db.c.Query(query_GETUSERS, "^"+search, offset, limit)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if rows.Err() != nil {
			return nil, err
		}

		var u User
		if err := rows.Scan(&u.UserID, &u.Username); err != nil {
			return nil, err
		}

		isBanned, err := db.IsBanned(userID, u.UserID)
		if err != nil {
			return nil, err
		}
		if !isBanned {
			users = append(users, u)

		}
	}
	defer func() { err = rows.Close() }()

	return users, err
}
