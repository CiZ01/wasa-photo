package database

import "time"

var query_CREATECOMMENT = `INSERT INTO Comment (commentID, userID, ownerID, postID, commentText) VALUES (?, ?, ?, ?, ?);`

func (db *appdbimpl) CreateComment(userID uint32, ownerID uint32, postID uint32, commentText string) (Comment, error) {
	var comment Comment

	// Get the last commentID
	var lastCommentID uint32
	lastCommentID, err := db.GetLastCommentID(ownerID, postID)
	if err != nil {
		return comment, err
	}

	// Create the comment
	_, err = db.c.Exec(query_CREATECOMMENT, lastCommentID+1, userID, ownerID, postID, commentText)
	if err != nil {
		return comment, err
	}

	user, err := db.GetUserByID(userID)
	if err != nil {
		return comment, err
	}

	var timestamp time.Time
	err = db.c.QueryRow(
		`SELECT timestamp FROM Comment WHERE commentID=? AND ownerID=? AND postID=?;`,
		lastCommentID+1, ownerID, postID).Scan(&timestamp)
	if err != nil {
		return comment, err
	}

	comment = Comment{
		CommentID: lastCommentID + 1,
		User:      user,
		Text:      commentText,
		Timestamp: timestamp,
	}

	return comment, nil
}
