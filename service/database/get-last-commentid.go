package database

import "database/sql"

var query_GETLASTCOMMENTID = `SELECT MAX(commentID) FROM Comment WHERE ownerID=? AND postID=?;`

func (db *appdbimpl) GetLastCommentID(ownerID uint32, postID uint32) (uint32, error) {
	_lastCommentID := sql.NullInt32{Int32: 0, Valid: false}
	var lastCommentID uint32 = 0
	err := db.c.QueryRow(query_GETLASTCOMMENTID, ownerID, postID).Scan(&lastCommentID)
	if err != nil {
		return 0, err
	}

	if _lastCommentID.Valid == false {
		lastCommentID = 1
	} else {
		lastCommentID = uint32(_lastCommentID.Int32)
	}

	return lastCommentID, nil
}
