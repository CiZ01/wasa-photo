package database

var query_GETALLCOMMENTS = `SELECT commentID, userID, ownerID, postID, commentText, timestamp FROM Comment WHERE ownerID = ? AND postID = ? AND hidden="0" LIMIT ?,?`

func (db *appdbimpl) GetComments(ownerID int, postID int, offset int, limit int) ([]Comment, error) {
	var comments []Comment
	rows, err := db.c.Query(query_GETALLCOMMENTS, ownerID, postID, offset, limit)
	if err != nil {
		return nil, err
	}

	defer func() { err = rows.Close() }()

	for rows.Next() {
		if rows.Err() != nil {
			return nil, err
		}
		var comment Comment
		var userID int
		err = rows.Scan(&comment.CommentID, &userID, &comment.OwnerID, &comment.PostID, &comment.Text, &comment.Timestamp)
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

	return comments, err
}
