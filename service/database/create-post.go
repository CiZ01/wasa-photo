package database

import (
	"database/sql"
	"fmt"
)

var query_CREATEPOST = `INSERT INTO Post (postID, userID, postImageURL, caption) VALUES (?, ?, ?, ?)`

func (db *appdbimpl) CreatePost(p Post) (Post, error) {
	// Check last postID in db
	_postID, err := db.GetLastPostID(p.User.UserID)
	if err != nil {
		return p, err
	}

	tx, err := db.c.BeginTx(db.ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return p, err
	}

	defer func() {
		if err != nil {
			err = tx.Rollback()
		}
		err = tx.Commit()
	}()

	postID := fmt.Sprint(_postID + 1)
	userID := fmt.Sprint(p.User.UserID)

	p.PostID = postID

	path := "./storage/" + userID + "/posts/" + postID + ".jpg"

	_, err = db.c.Exec(query_CREATEPOST, postID, userID, path, p.Caption)
	if err != nil {
		return p, err
	}

	newPost, err := db.GetPosts(p.User.UserID, p.User.UserID, 0, 1)
	if err != nil {
		return p, err
	}

	return newPost[0], nil
}
