package database

var query_GETFOLLOWINGS = `SELECT userID, username, profileImageURL FROM User WHERE userID IN (SELECT followingID FROM Follow WHERE followerID=?) LIMIT ?, ?`

func (db *appdbimpl) GetFollowings(userID uint32, limit uint32, offset uint32) ([]User, error) {
	// Get the followings from the database
	rows, err := db.c.Query(query_GETFOLLOWINGS, userID, offset, limit)
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
