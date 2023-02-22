package database

import "database/sql"

var query_GETLASPOSTID = `SELECT MAX(postID) FROM Post WHERE userID= ?;`

func (db *appdbimpl) GetLastPostID(userID int) (int, error) {
	var _postID = sql.NullInt64{Int64: 0, Valid: false}
	var postID int = 0
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
			postID = int(_postID.Int64)
		}
	}

	return postID, err
}
