package database

var query_GETLASTCOMMENTID = `SELECT MAX(commentID) FROM Comment WHERE ownerID=? AND postID=?;`

func (db *appdbimpl) GetLastCommentID(ownerID uint32, postID uint32) (uint32, error) {
	var lastCommentID = new(uint32)
	err := db.c.QueryRow(query_GETLASTCOMMENTID, ownerID, postID).Scan(&lastCommentID)
	if err != nil {
		return 0, err
	}

	if lastCommentID == nil {
		lastCommentID = new(uint32)
		*lastCommentID = uint32(1)
	}

	return *lastCommentID, nil
}
