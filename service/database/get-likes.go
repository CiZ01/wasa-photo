package database

var query_GETLIKES = `SELECT userID FROM Like WHERE postID = ? AND ownerID = ? LIMIT ?,?`

func (db *appdbimpl) GetLikes(ownerID uint32, postID uint32, offset uint32, limit uint32) ([]User, error) {
	var likes []User
	rows, err := db.c.Query(query_GETLIKES, ownerID, postID, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var userID uint32
		err = rows.Scan(&userID)
		if err != nil {
			return nil, err
		}
		user, err := db.GetUserByID(userID)
		if err != nil {
			return nil, err
		}
		likes = append(likes, user)
	}
	return likes, nil
}
