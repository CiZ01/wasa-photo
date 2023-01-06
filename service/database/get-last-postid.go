package database

var query_GETLASPOSTID = `SELECT MAX(postID) FROM Post WHERE userID= ?;`

func (db *appdbimpl) GetLastPostID(userID uint32) (uint32, error) {
	var postID *uint32 = new(uint32)
	res, err := db.c.Query(query_GETLASPOSTID, userID)
	if err != nil {
		return 0, err
	}
	defer func() { err = res.Close() }()

	for res.Next() {
		if err := res.Err(); err != nil {
			return 0, err
		}
		
		err = res.Scan(&postID)
		if err != nil {
			return 0, err
		}

		if postID == nil {
			postID = new(uint32)
			*postID = 1
		}
	}

	return *postID, err
}
