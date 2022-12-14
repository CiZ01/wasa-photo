package database

var query_DELETECOMMENT = `DELETE FROM Comment WHERE 
commentID=? AND ownerID=? AND postID=?`

func (db *appdbimpl) DeleteComment(commentID uint32, ownerID uint32, postID uint32) error {
	_, err := db.c.Exec(query_DELETECOMMENT, commentID, ownerID, postID)
	if err != nil {
		return err
	}

	return nil
}
