package database

var query_DELETELIKE = `DELETE FROM Like WHERE userID = ? AND postID = ? AND ownerID = ?`

func (db *appdbimpl) DeleteLike(ownerID uint32, postID uint32, userID uint32) error {
	_, err := db.c.Exec(query_DELETELIKE, ownerID, postID, userID)
	if err != nil {
		return err
	}

	return nil
}
