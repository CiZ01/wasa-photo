package database

var query_ISFOLLOWING = `SELECT * FROM Follow WHERE followerID = ? AND followedID = ?`

func (db *appdbimpl) IsFollowing(followerID uint32, followedID uint32) (bool, error) {
	res, err := db.c.Query(query_ISFOLLOWING, followerID, followedID)
	if err != nil {
		return false, err
	}

	var isFollowing string
	res.Scan(&isFollowing)
	return isFollowing != "", nil
}
