package database

import "database/sql"

var query_GETLASPOSTID = `SELECT MAX(postID) FROM Post WHERE userID= ?;`

func (db *appdbimpl) GetLastPostID(userID uint32) (uint32, error) {
	var _postID sql.NullInt32 = sql.NullInt32{Int32: 0, Valid: false}
	var postID uint32 = 0
	res, err := db.c.Query(query_GETLASPOSTID, userID)
	if err != nil {
		return 0, err
	}
	defer func() { err = res.Close() }()

	for res.Next() {
		if err := res.Err(); err != nil {
			return 0, err
		}

		err = res.Scan(&_postID)
		if err != nil && err != sql.ErrNoRows {
			return 0, err
		}

		if !_postID.Valid {
			postID = 0
		} else {
			postID = uint32(_postID.Int32)
		}
	}

	return postID, err
}
