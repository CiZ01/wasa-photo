package database

var query_ISFOLLOWING = `SELECT followerID FROM Follow WHERE followerID = ? AND followedID = ?`

func (db *appdbimpl) IsFollowing(followerID int, followedID int) (bool, error) {
	var isFollowing string
	rows, err := db.c.Query(query_ISFOLLOWING, followerID, followedID)
	if err != nil {
		return false, err
	}

	defer func() { err = rows.Close() }()

	for rows.Next() {
		if rows.Err() != nil {
			return false, err
		}
		err := rows.Scan(&isFollowing)
		if err != nil {
			return false, err
		}
	}
	return isFollowing != "", err
}
