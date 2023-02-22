package database

var query_GETFOLLOWINGS = `SELECT userID, username, userPropicURL FROM User WHERE userID IN (SELECT followedID FROM Follow WHERE followerID=? LIMIT ?, ?)`

func (db *appdbimpl) GetFollowings(followerID int, offset int, limit int) ([]User, error) {
	// Get the followings from the database
	rows, err := db.c.Query(query_GETFOLLOWINGS, followerID, offset, limit)
	if err != nil {
		return nil, err
	}
	defer func() { err = rows.Close() }()

	var followings []User

	for rows.Next() {
		if rows.Err() != nil {
			return nil, err
		}
		var following User

		// Get following data
		err = rows.Scan(&following.UserID, &following.Username, &following.UserPropicURL)
		if err != nil {
			return nil, err
		}
		followings = append(followings, following)
	}

	return followings, err
}
