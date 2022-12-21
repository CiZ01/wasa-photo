package database

var query_GETFOLLOWINGS = `SELECT userID, username, userPropicURL FROM User WHERE userID IN (SELECT followedID FROM Follow WHERE followerID=? LIMIT ?, ?)`

func (db *appdbimpl) GetFollowings(followerID uint32, offset uint32, limit uint32) ([]User, error) {
	// Get the followings from the database
	rows, err := db.c.Query(query_GETFOLLOWINGS, followerID, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followings []User

	for rows.Next() {
		var following User

		// Get following data
		err = rows.Scan(&following.UserID, &following.Username, &following.UserPropicURL)
		if err != nil {
			return nil, err
		}
		followings = append(followings, following)
	}

	return followings, nil
}
