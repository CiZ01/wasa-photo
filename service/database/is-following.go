package database

import "fmt"

var query_ISFOLLOWING = `SELECT followerID FROM Follow WHERE followerID = ? AND followedID = ?`

func (db *appdbimpl) IsFollowing(followerID uint32, followedID uint32) (bool, error) {
	var isFollowing string
	rows, err := db.c.Query(query_ISFOLLOWING, followerID, followedID)
	if err != nil {
		return false, err
	}

	defer rows.Close()
	for rows.Next() {
		rows.Scan(&isFollowing)
	}
	fmt.Printf("isFollowing: %s", isFollowing)
	return isFollowing != "", nil
}
