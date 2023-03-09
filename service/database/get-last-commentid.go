package database

import (
	"database/sql"
	"errors"
)

var query_GETLASTCOMMENTID = `SELECT MAX(commentID) FROM Comment WHERE ownerID=? AND postID=?;`

func (db *appdbimpl) GetLastCommentID(ownerID int, postID int) (int, error) {
	_lastCommentID := sql.NullInt64{Int64: 0, Valid: false}
	var lastCommentID int = 0
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
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return 0, err
		}

		if !_lastCommentID.Valid {
			lastCommentID = 0
		} else {
			lastCommentID = int(_lastCommentID.Int64)
		}
	}

	return lastCommentID, err
}
