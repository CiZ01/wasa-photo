package database

import (
	"database/sql"
	"errors"
	"time"
)

var query_CREATECOMMENT = `INSERT INTO Comment (commentID, userID, ownerID, postID, commentText) VALUES (?, ?, ?, ?, ?);`

func (db *appdbimpl) CreateComment(userID int, ownerID int, postID int, commentText string) (Comment, error) {
	var comment Comment

	// Get the last commentID
	var lastCommentID int
	lastCommentID, err := db.GetLastCommentID(ownerID, postID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
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
		PostID:    postID,
		OwnerID:   ownerID,
		Text:      commentText,
		Timestamp: time.Time(time.Now()),
	}

	return comment, err
}
