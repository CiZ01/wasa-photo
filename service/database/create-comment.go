package database

import (
	"database/sql"
	"time"
)

var query_CREATECOMMENT = `INSERT INTO Comment (commentID, userID, ownerID, postID, commentText) VALUES (?, ?, ?, ?, ?);`

func (db *appdbimpl) CreateComment(userID uint32, ownerID uint32, postID uint32, commentText string) (Comment, error) {
	var comment Comment

	// Get the last commentID
	var lastCommentID uint32
	lastCommentID, err := db.GetLastCommentID(ownerID, postID)
	if err != nil && err != sql.ErrNoRows {
		return comment, err
	}

	tx, err := db.c.BeginTx(db.ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return comment, err
	}

	defer func() {
		if err != nil {
			err = tx.Rollback()
		}
		err = tx.Commit()
	}()

	// Create the comment
	_, err = tx.Exec(query_CREATECOMMENT, lastCommentID+1, userID, ownerID, postID, commentText)
	if err != nil {
		return comment, err
	}

	user, err := db.GetUserByID(userID)
	if err != nil {
		return comment, err
	}

	comment = Comment{
		CommentID: lastCommentID + 1,
		User:      user,
		Text:      commentText,
		Timestamp: time.Time(time.Now()),
	}

	return comment, nil
}
