package database

import "fmt"

var query_GETFOLLOWERS = `SELECT userID, username, userPropicURL FROM User WHERE userID IN (SELECT followerID FROM Follow WHERE followedID=? LIMIT ?, ?)`

func (db *appdbimpl) GetFollowers(followedID uint32, offset uint32, limit uint32) ([]User, error) {
	// Get the followers from the database
	rows, err := db.c.Query(query_GETFOLLOWERS, followedID, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followers []User

	for rows.Next() {
		var follower User

		// Get follower data
		err = rows.Scan(&follower.UserID, &follower.Username, &follower.UserPropicURL)
		if err != nil {
			return nil, err
		}
		fmt.Printf("%+v", follower)
		followers = append(followers, follower)
	}

	return followers, nil
}
