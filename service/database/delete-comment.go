package database

var query_DELETECOMMENT = `DELETE FROM Comment WHERE 
commentID=? AND ownerID=? AND postID=?`

/*
DeleteComment deletes the comment with the given commentID, from the post with the given postID, and the owner with the given ownerID.
Return an error if the comment does not exist.
*/
func (db *appdbimpl) DeleteComment(commentID uint32, ownerID uint32, postID uint32) error {
	_, err := db.c.Exec(query_DELETECOMMENT, commentID, ownerID, postID)
	if err != nil {
		return err
	}

	return nil
}
