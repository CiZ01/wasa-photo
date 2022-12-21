package database

var query_GETALLCOMMENTS = `SELECT commentID, userID, commentText, timestamp FROM Comment WHERE ownerID = ? AND postID = ? AND hidden="0" LIMIT ?,?`

func (db *appdbimpl) GetComments(ownerID uint32, postID uint32, offset uint32, limit uint32) ([]Comment, error) {
	var comments []Comment
	rows, err := db.c.Query(query_GETALLCOMMENTS, ownerID, postID, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	
	for rows.Next() {
		var comment Comment
		var userID uint32
		err = rows.Scan(&comment.CommentID, &userID, &comment.Text, &comment.Timestamp)
		if err != nil {
			return nil, err
		}
		user, err := db.GetUserByID(userID)
		if err != nil {
			return nil, err
		}
		comment.User = user

		comments = append(comments, comment)
	}
	return comments, nil
}
