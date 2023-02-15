package database

import (
	"database/sql"
)

var query_GETLASTCOMMENTID = `SELECT MAX(commentID) FROM Comment WHERE ownerID=? AND postID=?;`

func (db *appdbimpl) GetLastCommentID(ownerID uint32, postID uint32) (uint32, error) {
	var _lastCommentID sql.NullInt32 = sql.NullInt32{Int32: 0, Valid: false}
	var lastCommentID uint32 = 0
	rows, err := db.c.Query(query_GETLASTCOMMENTID, ownerID, postID)
	if err != nil {
		return 0, err
	}
	defer func() { err = rows.Close() }()

	for rows.Next() {
		if err := rows.Err(); err != nil {
			return 0, err
		}

		err = rows.Scan(&_lastCommentID)
		if err != nil && err != sql.ErrNoRows {
			return 0, err
		}

		if !_lastCommentID.Valid {
			lastCommentID = 0
		} else {
			lastCommentID = uint32(_lastCommentID.Int32)
		}
	}

	return lastCommentID, err
}
